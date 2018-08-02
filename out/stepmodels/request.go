package stepmodels

import (
	"github.com/ghostsquad/slack-off/resourcemodels"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"encoding/json"
	"github.com/ghostsquad/slack-off"
)

type Request struct {
	Source resourcemodels.Source	`json:"source"  validate:"required"`
	Params Params 								`json:"params"  validate:"required"`
}

func (req *Request) Load(re io.Reader) error {
	return json.NewDecoder(re).Decode(req)
}

func (req *Request) RegisterValidations(val *validator.Validate) {
	req.Params.RegisterValidations(val)
}

// Interface assertions
var _ slackoff.Validatable = (*Request)(nil)
