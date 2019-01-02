package stepmodels

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/ghostsquad/slack-off"
	"github.com/hashicorp/go-multierror"
	"path"
)

type params struct {
	Template        string            `json:"template"`
	TemplateFile    string            `json:"template_file"`
	FileVars        map[string]string `json:"file_vars"      validate:"dive,keys,required,endkeys,required"`
	Vars            map[string]string `json:"vars"           validate:"dive,keys,required,endkeys,required"`
	srcDir		 string
	fileReader slackoff.FileReader
}

func (p *params) RegisterValidations(val *validator.Validate) {
	val.RegisterStructValidation(paramsStructLevelValidation, params{})
}

func paramsStructLevelValidation(sl validator.StructLevel) {
	p := sl.Current().Interface().(params)

	templateOptions := 0

	if len(p.Template) > 0 {
		templateOptions++
	}

	if len(p.TemplateFile) > 0 {
		templateOptions++
	}

	if templateOptions != 1 {
		sl.ReportError(p.Template, "Template", "template", "templateortemplate_file", "")
		sl.ReportError(p.Template, "TemplateFile", "template_file", "templateortemplate_file", "")
	}
}

// this relies on validation to assert that exactly 1 of Template or TemplateFile are provided
func (p *params) GetTemplate() (template string, err error) {
	if len(p.Template) > 0 {
		template = p.Template
	} else if len(p.TemplateFile) > 0 {
		template, err = p.fileReader.ReadFile(path.Join(p.srcDir, p.TemplateFile))
	}

	return
}

func (p *params) GetFileVars() (map[string]string, error) {
	var errs *multierror.Error
	fileVars := make(map[string]string)

	for k, v := range p.FileVars {
		content, readErr := p.fileReader.ReadFile(path.Join(p.srcDir, v))
		if readErr != nil {
			errs = multierror.Append(errs, readErr)
		}
		fileVars[k] = content
	}

	err := errs.ErrorOrNil()

	return fileVars, err
}

// Interface assertions
var _ slackoff.ValidatorRegisterable = (*params)(nil)
