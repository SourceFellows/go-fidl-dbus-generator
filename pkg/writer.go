package pkg

import (
	"io"
	"text/template"
)

var typeTemplate = `type {{.InterfaceName}} strutc {
	{{range .Entry.TypeRef}}
		{{if .Method}}
		func {{.Method.Name}}() {}
		{{end}}
	{{end}}
}`

func Write(fidl *Fidl, writer io.Writer) error {

	tmpl, err := template.New("type").
		Parse(typeTemplate)

	if err != nil {
		return err
	}

	return tmpl.Execute(writer, fidl)

}
