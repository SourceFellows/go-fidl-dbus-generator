# Go FIDL DBus Generator

Go FIDL DBus Generator is a tool which parses FIDL files and generates the
corresponding sender/receiver implementation of DBus services.

## How to run

`go run pkg/cmd/main.go -sender -in "path/to/fidl/file"`