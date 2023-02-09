# Go FIDL DBus Generator

Go FIDL DBus Generator is a tool which parses FIDL files and generates the
corresponding sender/receiver implementation of DBus services.

## Install

```
go install github.com/SourceFellows/go-fidl-dbus-generator/pkg/cmd/go-fidl@latest
```

## How to run

Parameters

| Name     | Description                                                                         |
|----------|-------------------------------------------------------------------------------------|
| in       | FIDL file which should be parsed                                                    |
| out      | optional output file. if nothing is specified the result is printed to the terminal |
| package  | target package                                                                      |
| receiver | indicates that receiver code should be generated                                    |
| sender   | indicates that serevr code should be generated                                      |
| debug    | show debug information                                                              |


Sample:

`go-fidl -sender -in "path/to/fidl/file"`

## Generate the examples

```
go run cmd/go-fidl/main.go -in ../examples/Notifications.fidl -package notification -sender -out ../examples/notification/NotificationSender.go
```