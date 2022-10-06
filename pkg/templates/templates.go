package templates

import _ "embed"

//go:embed Sender-template.gotmpl
var SenderTemplate string

//go:embed Receiver-template.gotmpl
var ReceiverTemplate string

//go:embed Struct.gotmpl
var StructTemplate string

//go:embed DBusInterface.gotmpl
var DBusInterfaceTemplate string
