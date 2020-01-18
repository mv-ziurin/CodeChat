package template

import (
	"fmt"
	"html/template"
	"github.com/inwady/easyconfig"
	"bytes"
	"encoding/json"
	"easysender/template/data"
)

type TemplateExecuter interface {
	Execute(json []byte) (string, error)
}

type TemplateData struct {
	template *template.Template
	generator func() interface{}
}

var (
	templates = map[string]*TemplateData{
		"greeting": {
			template: generateTemplate(easyconfig.GetString("template.file.greeting", "")),
			generator: func() interface{} { return new(data.SimpleData) },
		},

		"passrestore": {
			template: generateTemplate(easyconfig.GetString("template.file.passrestore", "")),
			generator: func() interface{} { return new(data.SimpleData) },
		},

		"subscribe": {
			template: generateTemplate(easyconfig.GetString("template.file.subscribe", "")),
			generator: func() interface{} { return new(interface{}) },
		},
	}
)

func ExecuteTemplate(t string, json []byte) (string, error) {
	templateData, ok := templates[t]
	if !ok {
		return "", fmt.Errorf("not found template")
	}

	return templateData.execute(json)
}

func (gt *TemplateData) execute(data []byte) (string, error) {
	var tpl bytes.Buffer

	gtd := gt.generator()
	err := json.Unmarshal(data, gtd)
	if err != nil {
		return "", err
	}

	err = gt.template.Execute(&tpl, gtd)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func generateTemplate(fileName string) *template.Template {
	return template.Must(template.ParseFiles(fileName))
}
