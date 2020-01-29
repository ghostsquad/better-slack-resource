package stepmodels

import (
	"github.com/ghostsquad/slack-off/resourcemodels"
	"gopkg.in/go-playground/validator.v9"
	"github.com/ghostsquad/slack-off"
)

type Request struct {
	Source resourcemodels.Source `json:"source" validate:"required"`
	Params params                `json:"params" validate:"required"`
}

func NewRequest(srcDir string, fileReader slackoff.FileReader) *Request {
	return &Request{
		Params: params{
			srcDir: srcDir,
			fileReader: fileReader,
		},
	}
}

func (req *Request) RegisterValidations(val *validator.Validate) {
	req.Params.RegisterValidations(val)
}

// Interface assertions
var _ slackoff.ValidatorRegisterable = (*Request)(nil)
