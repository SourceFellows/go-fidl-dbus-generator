package pkg

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"strings"
	"text/template"
)

//go:embed Client-template.gotmpl
var clientTemplate string

func Write(fidl *Fidl, writer io.Writer) error {

	funcs := template.FuncMap{
		"nameify":      nameify,
		"exportNameOf": exportNameOf,
		"goType":       convertToGoType,
		"derefStr":     deref,
	}

	tmpl, err := template.New("type").
		Funcs(funcs).
		Parse(clientTemplate)

	if err != nil {
		return err
	}

	var bites bytes.Buffer

	err = tmpl.Execute(&bites, fidl)
	if err != nil {
		return err
	}

	src, err := format.Source(bites.Bytes())
	if err != nil {
		return err
	}

	fmt.Println(string(src))
	return nil
}

func nameify(typeName string) string {
	first := strings.ToLower(string(typeName[0]))
	return fmt.Sprintf("%s%sImpl", first, typeName[1:])
}

func exportNameOf(name string) string {
	first := strings.ToUpper(string(name[0]))
	return fmt.Sprintf("%s%s", first, name[1:])
}

func convertToGoType(fidlString string) string {

	if fidlString == "String" {
		return "string"
	}
	if fidlString == "Boolean" {
		return "bool"
	}

	return fidlString
}

func deref(val *string) any {
	return *val
}
