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

func GetMessages(msg models.Message) ([]models.Message, error) {
	messageResponse, err := repository.GetMessages(msg.SenderId, msg.RecieverId)
	return messageResponse, err
}

func GetConversations(userId int) ([]structs.ConversationData, error) {
	conversationResponse, err := repository.GetConversations(userId)
	return conversationResponse, err
}
