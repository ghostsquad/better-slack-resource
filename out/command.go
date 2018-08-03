package out

import (
	"github.com/ghostsquad/slack-off/resourcemodels"
	"github.com/ghostsquad/slack-off/out/stepmodels"
	"io"
	"github.com/ghostsquad/slack-off"
	"gopkg.in/go-playground/validator.v9"
	"fmt"
	"text/template"
	"bytes"
	"encoding/json"
)

type command struct {
	fileReader			slackoff.FileReader
	writer 			io.Writer
	httpPoster	slackoff.HttpPoster
}

func NewCommand(fileReader slackoff.FileReader, writer io.Writer, httpPoster slackoff.HttpPoster) *command {
	return &command{
		fileReader: fileReader,
		writer: writer,
		httpPoster: httpPoster,
	}
}

func (c *command) Run(request stepmodels.Request) (*stepmodels.Response, error) {
	v := slackoff.InitValidator()

	request.RegisterValidations(v)

	errs := v.Struct(request)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			fmt.Fprintf(c.writer, "%s", err.Namespace())
			fmt.Fprintf(c.writer, "%s", err.Field())
			// can differ when a custom TagNameFunc is registered or
			fmt.Fprintf(c.writer, "%s", err.StructNamespace())
			// by passing alt name to ReportError like below
			fmt.Fprintf(c.writer, "%s", err.StructField())
			fmt.Fprintf(c.writer, "%s", err.Tag())
			fmt.Fprintf(c.writer, "%s", err.ActualTag())
			fmt.Fprintf(c.writer, "%s", err.Kind())
			fmt.Fprintf(c.writer, "%s", err.Type())
			fmt.Fprintf(c.writer, "%s", err.Value())
			fmt.Fprintf(c.writer, "%s", err.Param())
			fmt.Fprintf(c.writer, "")
		}

		return nil, errs
	}

	channels, err := request.GetAllChannels(c.fileReader)
	if err != nil {
		return nil, err
	}

	fileVars, err := request.Params.GetFileVars(c.fileReader)
	if err != nil {
		return nil, err
	}

	templateText, err := request.Params.GetTemplate(c.fileReader)
	if err != nil {
		return nil, err
	}

	data := templateData{
		Vars: request.Params.Vars,
		FileVars: fileVars,
	}

	message, err := renderTemplate(templateText, data)
	if err != nil {
		return nil, err
	}

	webhookMsg := &slackoff.WebhookMessage{}
	err = json.Unmarshal([]byte(message), *webhookMsg)
	if err != nil {
		return nil, err
	}

	webhookMsg.IconUrl = request.Params.IconUrl
	webhookMsg.IconEmoji = request.Params.IconEmoji

	for _, channel := range channels {
		webhookMsg.Channel = channel

		err = slackoff.PostWebhookMessage(c.httpPoster, request.Source.Url, webhookMsg)
		if err != nil {
			return nil, err
		}
	}

	return &stepmodels.Response{
		Version:  resourcemodels.Version{},
		Metadata: []resourcemodels.MetadataPair{},
	}, nil
}

type templateData struct {
	Vars     map[string]string
	FileVars map[string]string
}

func renderTemplate(tmpl string, data templateData) (string, error) {
	t := template.New("slack")
	t.Parse(tmpl)

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
