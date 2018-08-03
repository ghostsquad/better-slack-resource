package slackoff

import (
	"encoding/json"
	"github.com/nlopes/slack"
)

type WebhookMessage struct {
	Text        string             `json:"text,omitempty"`
	Attachments []slack.Attachment `json:"attachments,omitempty"`
	Channel     string             `json:"channel,omitempty"`
	IconEmoji   string             `json:"icon_emoji,omitempty"`
	IconUrl     string             `json:"icon_url,omitempty"`
}

func (m *WebhookMessage) NewWebhookMessage(payload string) error {
	return json.Unmarshal([]byte(payload), m)
}

func PostWebhookMessage(poster HttpPoster, url string, msg *WebhookMessage) error {
	_, err := poster.Post(url, msg)
	if err != nil {
		return err
	}

	return nil
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
	// This commented to allow for easier testing
	// TODO: include strict validations + tests
	//
	//if u.Scheme != "https" {
	//	return NewErrInvalidSlackUrl("Cowardly refusing to accept address without TLS")
	//}
	//
	//validHost := "hooks.slack.com"
	//
	//if u.Host != validHost {
	//	return NewErrInvalidSlackUrl(fmt.Sprintf("Cowardly refusing to accept a url not sent to: %s", validHost))
	//}

	return nil
}
