package main

import (
	"twitter-back-end/initializers"

	"twitter-back-end/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {

	//initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {

	// opt := option.WithCredentialsFile("twitter-clone-695cb-firebase-adminsdk-vxd8z-52c0580c7f.json")
	// app, err := firebase.NewApp(context.Background(), nil, opt)
	// if err != nil {
	// 	log.Fatalf("Error initializing Firebase app: %v", err)
	// }

	// // Get the messaging client.
	// client, err := app.Messaging(context.Background())
	// if err != nil {
	// 	log.Fatalf("Error getting Messaging client: %v", err)
	// }

	// // Construct a message.
	// message := &messaging.Message{
	// 	Notification: &messaging.Notification{
	// 		Title: "Hello, User",
	// 		Body:  "This is a sample notification",
	// 	},
	// 	Token: "e-8Nhrdas1AwocOp2_ftXh:APA91bFIM7Tur7XKk_Zpdt50n-IsJrn8K5U_XvudzlUoDamgaZcemSVCCmCdMKa5GIXTKhU7xKzK2w9confHI0i8UVs6-1wGyO151QnVPjoQyEg9cjtIrQSqCY6zpdhMF54A1161TVPT", // replace with the target device token
	// }

	// // Send a message to the desired device.
	// response, err := client.Send(context.Background(), message)
	// if err != nil {
	// 	log.Fatalf("Error sending message: %v", err)
	// }

	// fmt.Printf("Successfully sent message: %s", response)

	r := gin.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowHeaders("ThirdParty")

	r.Use(cors.New(corsConfig))

	api.HandleApi(r)

	r.Run()
}
