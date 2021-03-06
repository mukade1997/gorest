package constructors

import (
	"io"
	"text/template"

	"github.com/kepkin/gorest/internal/generator/constructors/fields"
	"github.com/kepkin/gorest/internal/generator/translator"
)

// MakePathParamsConstructor receive a path params struct definition and generate corresponding constructor
func MakePathParamsConstructor(wr io.Writer, def translator.TypeDef) error {
	return pathParamsConstructorTemplate.Execute(wr, def)
}

var pathParamsConstructorTemplate = template.Must(template.New("pathParamsConstructor").Funcs(fields.BaseConstructor).Parse(`
func Make{{ .Name }}(c *gin.Context) (result {{ .Name }}, errors []FieldError) {
	{{- if .HasNoStringFields }}
	var err error
	{{ end }}

	{{- range $, $field := .Fields }}
	{{- with $field }}
		{{- if .CheckDefault}}
			{{ .StrVarName }}, ok := c.Params.Get("{{ .Parameter }}")
			if !ok {
			   {{ .StrVarName }} = "{{ .Schema.Default }}"
			}
		{{- else }}
			{{ .StrVarName }}, _ := c.Params.Get("{{ .Parameter }}")
		{{- end }}

		{{- BaseValueFieldConstructor . "InPath" }}

	{{- end -}}
	{{ end -}}
	return
}
`))
