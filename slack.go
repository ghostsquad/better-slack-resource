package slackoff

import (
	"github.com/nlopes/slack"
	"net/url"
	"fmt"
)

type WebhookMessage struct {
	Text    		string 							`json:"text,omitempty"`
	Attachments []slack.Attachment	`json:"attachments,omitempty"`
}

type ErrInvalidSlackUrl struct {
	message string
}
func NewErrInvalidSlackUrl(message string) *ErrInvalidSlackUrl {
	return &ErrInvalidSlackUrl{
		message: message,
	}
}
func (e *ErrInvalidSlackUrl) Error() string {
	return e.message
}

func AssertSlackUrl(slackUrl string) error {
	u, err := url.Parse(slackUrl)

	if err != nil {
		return err
	}

	if u.Scheme != "https" {
		return NewErrInvalidSlackUrl("Cowardly refusing to accept address without TLS")
	}

	validHost := "hooks.slack.com"

	if u.Host != validHost {
		return NewErrInvalidSlackUrl(fmt.Sprintf("Cowardly refusing to accept a url not sent to: %s", validHost))
	}

	return nil
}
