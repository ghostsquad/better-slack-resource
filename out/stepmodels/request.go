package stepmodels

import (
	"github.com/ghostsquad/slack-off/resourcemodels"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"encoding/json"
	"github.com/ghostsquad/slack-off"
)

type Request struct {
	Source resourcemodels.Source `json:"source" validate:"required"`
	Params Params                `json:"params" validate:"required"`
}

func (req *Request) Populate(re io.Reader) error {
	return json.NewDecoder(re).Decode(req)
}

func (req *Request) RegisterValidations(val *validator.Validate) {
	req.Params.RegisterValidations(val)
}

func (req *Request) GetAllChannels(reader slackoff.FileReader) (channels []string, err error) {
	if len(req.Params.Channel) > 0 {
		channels = append(channels, req.Params.Channel)
	} else if len(req.Source.Channel) > 0 {
		channels = append(channels, req.Source.Channel)
	}

	extraChannels, getErr := req.Params.GetExtraChannels(reader)
	if getErr != nil {
		err = getErr
		return
	}

	channels = append(channels, extraChannels...)

	return
}

// Interface assertions
var _ slackoff.Validatable = (*Request)(nil)
