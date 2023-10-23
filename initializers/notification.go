package initializers

import (
	"context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

var opt = option.WithCredentialsFile("twitter-clone-695cb-firebase-adminsdk-vxd8z-52c0580c7f.json")
var app, _ = firebase.NewApp(context.Background(), nil, opt)

// Get the messaging client.
var Client, _ = app.Messaging(context.Background())
