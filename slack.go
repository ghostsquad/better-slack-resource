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

func NewWebHookMessage(payload string) (*WebhookMessage, error) {
	m := &WebhookMessage{}
	err := m.FromJson(payload)
	return m, err
}

func (m *WebhookMessage) FromJson(payload string) error {
	return json.Unmarshal([]byte(payload), m)
}

func (m *WebhookMessage) ToJson() (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
