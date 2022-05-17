package pkg

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/SourceFellows/go-fidl-dbus-generator/pkg/templates"
	"go/format"
	"io"
	"strings"
	"text/template"
	"unicode"
)

type WriterType struct {
	template string
}

var (
	ServerWriter = WriterType{templates.ServerTemplate}
	ClientWriter = WriterType{templates.ClientTemplate}
)

func Write(fidl *Fidl, writerType WriterType, writer io.Writer) error {

	funcs := template.FuncMap{
		"nameify":               toGoIdentifierName,
		"extractLastPartOfName": extractLastPartOfName,
		"exportNameOf":          exportNameOf,
		"goType":                mapFidlTypeToGoType,
		"derefStr":              deref,
	}

	tmpl, err := template.New("type").
		Funcs(funcs).
		Parse(writerType.template)

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

	_, err = writer.Write(src)
	return err
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
	mappings["Boolean"] = "bool"
	mappings["UInt8"] = "uint8"
	mappings["UInt16"] = "uint16"

	if v, ok := mappings[fidlString]; ok {
		return v
	}
	return fidlString
}

func deref(val *string) any {
	return *val
}

func extractLastPartOfName(val string) string {
	idx := strings.LastIndex(val, ".")
	if idx == -1 {
		return val
	}
	return val[idx+1:]
}
