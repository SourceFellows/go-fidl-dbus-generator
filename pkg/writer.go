package pkg

import (
	"io"
	"text/template"
)

var typeTemplate = `
{{ $ImplementationName := nameify .InterfaceName }}

type {{$ImplementationName}} struct {
}

{{range .Entry.TypeRef}}
	{{if .Method}}
func (impl *{{$ImplementationName}}) {{.Method.Name}}() error {}
	{{end}}
{{end}}
`

func Write(fidl *Fidl, writer io.Writer) error {

	funcs := template.FuncMap{"nameify": nameify}

	tmpl, err := template.New("type").
		Funcs(funcs).
		Parse(typeTemplate)

	if err != nil {
		return err
	}

	return tmpl.Execute(writer, fidl)

}

func nameify(val string) string {
	return "val"
}
