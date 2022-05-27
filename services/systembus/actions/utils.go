package actions

import (
	"bytes"
	"text/template"
)

func applyTemplate(templateText string, data interface{}) (string, error) {
	tmpl, err := template.New("template").Parse(templateText)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)
	return buf.String(), nil
}
