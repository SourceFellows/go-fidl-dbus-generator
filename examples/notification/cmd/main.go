package main

import (
	"context"
	"fmt"
	"log"
	"notification"
)

func main() {

	notificationsClient, err := notification.NewNotificationsClient("org.freedesktop.Notifications",
		"/org/freedesktop/Notifications")
	if err != nil {
		log.Fatal(err)
	}

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
