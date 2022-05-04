package pkg

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"strings"
	"text/template"
	"unicode"
)

//go:embed Client-template.gotmpl
var clientTemplate string

func Write(fidl *Fidl, writer io.Writer) error {

	funcs := template.FuncMap{
		"nameify":      toGoIdentifierName,
		"exportNameOf": exportNameOf,
		"goType":       mapFidlTypeToGoType,
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

	//fmt.Println(string(bites.Bytes()))

	src, err := format.Source(bites.Bytes())
	if err != nil {
		return err
	}

	fmt.Println(string(src))
	return nil
}

func toGoIdentifierName(typeName string) string {

	internalName := []rune(typeName)
	makeLower := true
	for i, r := range internalName {
		if !unicode.IsLower(r) && makeLower {
			internalName[i] = unicode.ToLower(r)
		} else {
			makeLower = false
		}
	}

	return string(internalName)
}

func exportNameOf(name string) string {
	first := strings.ToUpper(string(name[0]))
	return fmt.Sprintf("%s%s", first, name[1:])
}

func mapFidlTypeToGoType(fidlString string) string {

	mappings := map[string]string{}
	mappings["String"] = "string"
	mappings["UInt8"] = "unit8"
	mappings["UInt16"] = "uint16"

	if v, ok := mappings[fidlString]; ok {
		return v
	}
	return fidlString
}

func deref(val *string) any {
	return *val
}
