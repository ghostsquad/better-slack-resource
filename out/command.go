package out

import (
	"github.com/ghostsquad/slack-off/resourcemodels"
	"github.com/ghostsquad/slack-off/out/stepmodels"
	"io"
	"github.com/ghostsquad/slack-off"
	"gopkg.in/go-playground/validator.v9"
	"fmt"
)

type Command struct {
	writer 			io.Writer
	httpPoster	slackoff.HttpPoster
}

func NewCommand(writer io.Writer, httpPoster slackoff.HttpPoster) *Command {
	return &Command{
		writer: writer,
		httpPoster: httpPoster,
	}
}

func (c *Command) Run(request stepmodels.Request) (*stepmodels.Response, error) {
	v := slackoff.InitValidator()

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

	return &stepmodels.Response{
		Version:  resourcemodels.Version{},
		Metadata: []resourcemodels.MetadataPair{},
	}, nil
}
