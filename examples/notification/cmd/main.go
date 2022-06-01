package main

import (
	"context"
	"fmt"
	"github.com/SourceFellows/go-fidl-dbus-generator/example/notification"
	"log"
)

//go:generate go-fidl -sender -package notification -in ../../Notifications.fidl -out ../NotificationSender.go

func main() {

	notificationsClient, err := notification.NewNotificationsSender("org.freedesktop.Notifications",
		"/org/freedesktop/Notifications")
	if err != nil {
		log.Fatal(err)
	}
	defer notificationsClient.Close()

	notify, err := notificationsClient.Notify(context.Background(),
		"app",
		1,
		"",
		"summary",
		"body",
		[]string{"default"},
		nil,
		100000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(notify)

}
