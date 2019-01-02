package out

import (
	"github.com/ghostsquad/slack-off/resourcemodels"
	"github.com/ghostsquad/slack-off/out/stepmodels"
	"io"
	"github.com/ghostsquad/slack-off"
	"time"
)

type command struct {
	srcDir      string
	writer      io.Writer
	httpPoster  slackoff.HttpPoster
}

func NewCommand(tmplzr Templatizer, writer io.Writer, httpPoster slackoff.HttpPoster) *command {
	return &command{
		writer: writer,
		httpPoster: httpPoster,
	}
}

func (c *command) Run(request *stepmodels.Request) (*stepmodels.Response, error) {
	v := slackoff.InitValidator()

	errs := slackoff.Validate(request, v)
	if errs != nil {
		slackoff.WriteValidationErrors(errs, c.writer)
		return nil, errs
	}

	webhookMsg, err := slackoff.NewWebHookMessage(message)
	if err != nil {
		return nil, err
	}

	// TODO look at the response, and do something with it
	_, err = c.httpPoster.Post(request.Source.Url, webhookMsg)
	if err != nil {
		return nil, err
	}

	return &stepmodels.Response{
		Version:  resourcemodels.Version{
			Timestamp: time.Now().Unix(),
		},
	}, nil
}
