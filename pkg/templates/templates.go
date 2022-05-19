package templates

import _ "embed"

//go:embed Sender-template.gotmpl
var SenderTemplate string

//go:embed Receiver-template.gotmpl
var ReceiverTemplate string
