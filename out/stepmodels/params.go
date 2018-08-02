package stepmodels

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/ghostsquad/slack-off"
	"os"
)

type Params struct {
	Template       	string 						`json:"template"`
	TemplateFile   	string 						`json:"template_file"`
	FileVars  		 	map[string]string `json:"file_vars"      validate:"dive,keys,required,endkeys,required"`
	Vars       		 	map[string]string `json:"vars"           validate:"dive,keys,required,endkeys,required"`
	Channel				 	string 						`json:"channel"`
	ChannelAppend  	string 						`json:"channel_append"`
	ChannelFile			string 						`json:"channel_file"`
	IconUrl					string 						`json:"icon_url"`
	IconEmoji				string 						`json:"icon_emoji"`
}

func (p *Params) RegisterValidations(val *validator.Validate) {
	val.RegisterStructValidation(ParamsStructLevelValidation, Params{})
}

func ParamsStructLevelValidation(sl validator.StructLevel) {
	params := sl.Current().Interface().(Params)

	if (len(params.Template) == 0 && len(params.TemplateFile) == 0) ||
		(len(params.Template) > 0 && len(params.TemplateFile) > 0) {
		sl.ReportError(params.Template, "Template", "template", "templateortemplate_file", "")
		sl.ReportError(params.Template, "TemplateFile", "template_file", "templateortemplate_file", "")
	}

	if len(params.TemplateFile) > 0 {
		if _, ok := fileExists(params.TemplateFile); !ok {
			sl.ReportError(params.TemplateFile, "TemplateFile", "template_file", "templatefileread", "")
		}
	}

	if len(params.ChannelFile) > 0 {
		if _, ok := fileExists(params.ChannelFile); !ok {
			sl.ReportError(params.ChannelFile, "ChannelFile", "channel_file", "channelfileread", "")
		}
	}
}

func fileExists(path string) (err error, ok bool) {
	stat, err := os.Stat(path)
	return err, err == nil && !stat.IsDir()
}

// Interface assertions
var _ slackoff.Validatable = (*Params)(nil)
