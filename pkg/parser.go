package main

import (
	"github.com/alecthomas/participle/v2"
)

type (
	Fidl struct {
		PackageName   string  `"package" @(Ident ("." Ident)*)`
		PackageImport string  `("import" @(Ident ("." Ident)* ("." "*") ("from" String)?))*`
		InterfaceName *string `("interface" @Ident`
		TypeName      *string `| "typeCollection" @Ident)?`
		Entry         *Entry  `"{" @@ "}"`
	}

	Entry struct {
		Version *Version   `"version" "{" @@ "}"`
		TypeRef []*TypeRef `@@*`
	}

	Version struct {
		MajorVersion int `"major" @Int`
		MinorVersion int `"minor" @Int`
	}

	TypeRef struct {
		Description *Description `("<" "*" "*" @@ "*" "*" ">")?`
		Attribute   *Attribute   `("attribute" @@`
		Method      *Method      `| "method" @@`
		Enumeration *Enumeration `| "enumeration" @@`
		Broadcast   *Broadcast   `| "broadcast" @@)*`
	}

	Attribute struct {
		Type       string `@Ident`
		Name       string `@Ident`
		Permission string `@Ident`
	}

	Method struct {
		Name   string  `@Ident`
		Params *Params `"{" @@ "}"`
	}

	Enumeration struct {
		Name   string  `@Ident`
		Values []*Enum `"{" @@* "}"`
	}

	Broadcast struct {
		Name      string   `@Ident`
		Selective *string  `"selective"?`
		OutParams []*Param `"{" "out" "{" @@* "}" "}"`
	}

	Enum struct {
		Description *Description `("<" "*" "*" @@ "*" "*" ">")*`
		Name        string       `(@Ident "=")*`
		Index       int          `@Int*`
	}

	Params struct {
		InParams  []*Param `("in" "{" @@* "}")*`
		OutParams []*Param `("out" "{" @@* "}")*`
	}

	Param struct {
		Description *Description `("<" "*" "*" @@ "*" "*" ">")?`
		Type        string       `@Ident`
		IsArray     *string      `("[" "]")?`
		Name        string       `@Ident`
	}

	Description struct {
		Tag     string `@("@" "description" ":")?`
		Content string `@(Ident* ("("* Ident* ")"* Ident* Int* String* "."* ","* ":"* "?"* "-"* "$"*)*)`
	}
)

var fidlParser = participle.MustBuild(&Fidl{}, participle.Unquote("String"))

func parseFidl(input []byte) (*Fidl, error) {
	fidl := &Fidl{}
	err := fidlParser.ParseBytes("", input, fidl)
	if err != nil {
		return nil, err
	}

	return fidl, nil
}
