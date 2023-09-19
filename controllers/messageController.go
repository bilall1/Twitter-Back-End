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

func SentMessage(c *gin.Context) {
	var body models.Message
	c.Bind(&body)

	isSentResponse, err := services.SendMessage(body)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"sent": isSentResponse,
	})
}

func GetMessages(c *gin.Context) {
	var params structs.MessagePage
	c.Bind(&params)

	messageResponse, err := services.GetMessages(params)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"messages": messageResponse,
	})
}

func GetConversations(c *gin.Context) {
	var params models.User
	c.Bind(&params)

	conversationResponse, err := services.GetConversations(params.Id)

	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"conversations": conversationResponse,
	})
}

func HandleSocketMessaging(c *gin.Context) {
	messageChannel := make(chan structs.Message)
	var params struct {
		Id int
	}
	c.Bind(&params)

	//Connection
	conn, err := initializers.Upgrader.Upgrade(c.Writer, c.Request, nil)
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
