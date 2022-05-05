package examples

import _ "embed"

//go:embed Notifications.fidl
var NotificationFidl []byte

//go:embed SystemdManager.fidl
var SystemManagerFidl []byte
