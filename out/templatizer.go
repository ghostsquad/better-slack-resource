package out

import (
	"bytes"
	"text/template"
	"github.com/ghostsquad/slack-off/out/stepmodels"
	"github.com/ghostsquad/slack-off"
)

type Templatizer interface {
	Render() (string, error)
}

type templatizer struct {
	tmpl string
	tmplData stepmodels.TemplateData
}

func NewTemplatizer(srcDir string, params stepmodels.Params, fileReader slackoff.FileReader) (*templatizer, error) {

	fileVars, err := params.GetFileVars(srcDir, fileReader)
	if err != nil {
		return nil, err
	}

	templateText, err := params.GetTemplate(srcDir, fileReader)
	if err != nil {
		return nil, err
	}

	return &templatizer{
		tmplData: stepmodels.TemplateData{
			FileVars: fileVars,
			Vars: params.Vars,
		},
		tmpl: templateText,
	}, nil
}

func (t *templatizer) Render() (string, error) {
	tm := template.New("slack")
	tm.Parse(t.tmpl)

	var tpl bytes.Buffer
	if err := tm.Execute(&tpl, t.tmplData); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

// Interface assertions
var _ Templatizer = (*templatizer)(nil)
