package services

import (
	"context"
	"fmt"
	"twitter-back-end/initializers"

	"firebase.google.com/go/messaging"
)

func SendMessageNotification(FirstName string, LastName string, Content string, Token string) {

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: FirstName + " " + LastName,
			Body:  Content,
		},
		Token: Token, // replace with the target device token
	}

	// Send a message to the desired device.
	response, err := initializers.Client.Send(context.Background(), message)
	if err != nil {
		fmt.Printf("Error sending message: %v", err)
	}
	fmt.Printf("Successfully sent message: %s", response)

}
