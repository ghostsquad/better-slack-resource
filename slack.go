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
