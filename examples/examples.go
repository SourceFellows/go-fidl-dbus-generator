package examples

import _ "embed"

//go:embed Notifications.fidl
var NotificationFidl []byte

//go:embed SystemdManager.fidl
var SystemManagerFidl []byte

//go:embed FireAndForget.fidl
var FireAndForgetsFidl []byte
