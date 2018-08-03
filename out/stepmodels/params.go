package stepmodels

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/ghostsquad/slack-off"
	"strings"
	"github.com/hashicorp/go-multierror"
)

type Params struct {
	Template        string            `json:"template"`
	TemplateFile    string            `json:"template_file"`
	FileVars        map[string]string `json:"file_vars"      validate:"dive,keys,required,endkeys,required"`
	Vars            map[string]string `json:"vars"           validate:"dive,keys,required,endkeys,required"`
	Channel         string            `json:"channel"`
	ChannelAppend   string            `json:"channel_append"`
	ChannelFile     string            `json:"channel_file"`
	IconUrl         string            `json:"icon_url"`
	IconEmoji       string            `json:"icon_emoji"`
}

func (p *Params) RegisterValidations(val *validator.Validate) {
	val.RegisterStructValidation(paramsStructLevelValidation, Params{})
}

func paramsStructLevelValidation(sl validator.StructLevel) {
	params := sl.Current().Interface().(Params)

	templateOptions := 0

	if len(params.Template) > 0 {
		templateOptions++
	}

	if len(params.TemplateFile) > 0 {
		templateOptions++
	}

	if templateOptions != 1 {
		sl.ReportError(params.Template, "Template", "template", "templateortemplate_file", "")
		sl.ReportError(params.Template, "TemplateFile", "template_file", "templateortemplate_file", "")
	}
}

// this relies on validation to assert that exactly 1 of Template or TemplateFile are provided
func (p *Params) GetTemplate(reader slackoff.FileReader) (template string, err error) {
	if len(p.Template) > 0 {
		template = p.Template
	} else if len(p.TemplateFile) > 0 {
		template, err = reader.ReadFile(p.TemplateFile)
	}

	return
}

func (p *Params) GetExtraChannels(reader slackoff.FileReader) (channels []string, err error) {
	if len(p.ChannelAppend) > 0 {
		channels = append(channels, strings.Fields(p.ChannelAppend)...)
	}

	if len(p.ChannelFile) > 0 {
		channelFileContent, readErr := reader.ReadFile(p.ChannelFile)
		if readErr != nil {
			err = readErr
			return
		}

		channels = append(channels, strings.Fields(channelFileContent)...)
	}

	return
}

func (p *Params) GetFileVars(reader slackoff.FileReader) (map[string]string, error) {
	var errs *multierror.Error
	fileVars := make(map[string]string)

	for k, v := range p.FileVars {
		content, readErr := reader.ReadFile(v)
		if readErr != nil {
			errs = multierror.Append(errs, readErr)
		}
		fileVars[k] = content
	}

	err := errs.ErrorOrNil()

	return fileVars, err
}

// Interface assertions
var _ slackoff.Validatable = (*Params)(nil)
