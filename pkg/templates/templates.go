package templates

import _ "embed"

//go:embed Client-template.gotmpl
var ClientTemplate string

//go:embed Server-template.gotmpl
var ServerTemplate string
