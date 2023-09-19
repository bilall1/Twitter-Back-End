// service.go
package services

import (
	"twitter-back-end/models"
	"twitter-back-end/repository"
	"twitter-back-end/structs"
)

func SendMessage(msg models.Message) (bool, error) {
	isSent, err := repository.SendMessage(msg.SenderId, msg.RecieverId, msg.MessageType, msg.Status, msg.Content)
	return isSent, err
}

func GetMessages(msg structs.MessagePage) ([]models.Message, error) {

	itemsPerPage := 10
	startIndex := (msg.Page - 1) * itemsPerPage

	messageResponse, err := repository.GetMessages(msg.SenderId, msg.RecieverId, itemsPerPage, startIndex)
	return messageResponse, err
}

func GetConversations(userId int) ([]structs.ConversationData, error) {
	conversationResponse, err := repository.GetConversations(userId)
	return conversationResponse, err
}
