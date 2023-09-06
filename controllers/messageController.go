package controllers

import (
	"twitter-back-end/models"
	"twitter-back-end/services"

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
	var params models.Message
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
