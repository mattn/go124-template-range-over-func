package main

import (
	"os"
	"strings"
	"text/template"
)

const contents = `
{{- range $l := .}}
{{- $l | trim }}
{{ end -}}
`

const lines = `
xxxx
yyyy
zzzz
`

func main() {
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"trim": strings.TrimSpace,
	}).Parse(contents))

	tmpl.Execute(os.Stdout, strings.Lines(strings.TrimSpace(lines)))
}
