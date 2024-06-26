package controllers

import (
	"encoding/json"
	"fmt"
	"twitter-back-end/initializers"
	"twitter-back-end/models"
	"twitter-back-end/services"
	"twitter-back-end/structs"

	"github.com/gin-gonic/gin"
)

func SentMessage(ctx *gin.Context) {
	var body models.Message
	ctx.Bind(&body)

	isSentResponse, err := services.SendMessage(body)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"sent": isSentResponse,
	})
}

func GetMessages(ctx *gin.Context) {
	var params structs.MessagePage
	ctx.Bind(&params)

	messageResponse, err := services.GetMessages(params)
	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"messages": messageResponse,
	})
}

func GetConversations(ctx *gin.Context) {
	var params models.User
	ctx.Bind(&params)

	conversationResponse, err := services.GetConversations(params.Id)

	if err != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"conversations": conversationResponse,
	})
}

func HandleSocketMessaging(ctx *gin.Context) {
	messageChannel := make(chan structs.Message)
	var params struct {
		Id int
	}
	ctx.Bind(&params)

	//Connection
	conn, err := initializers.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade the connection: %v\n", err)
		return
	}
	initializers.Clients[params.Id] = conn

	//Reading and writing

	go func() {

		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		var receivedMsg structs.Message
		err = json.Unmarshal(msg, &receivedMsg)

		messageChannel <- receivedMsg

	}()

	go func() {
		receivedMsg := <-messageChannel

		recipientConn, ok := initializers.Clients[receivedMsg.RecieverId]
		if ok {
			if err = recipientConn.WriteJSON(receivedMsg); err != nil {
				fmt.Printf("Error sending message to %d: %v\n", receivedMsg.RecieverId, err)
			}
		}

	}()

}
