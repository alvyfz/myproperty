package database

import (
	"myproperty-api/config"
	model "myproperty-api/models"
)

func GetChats() []model.Chat {
	var chats []model.Chat
	config.DB.Find(&chats)
	return chats
}

func GetChatByID(id string) model.Chat {
	var chat model.Chat
	config.DB.Where("id = ?", id).Find(&chat)
	return chat
}

func CreateChat(chat model.Chat) model.Chat {
	config.DB.Create(&chat)
	return chat
}

func DeleteChatByID(id string) {
	var chat model.Chat
	config.DB.Where("id = ?", id).Delete(&chat)
}

func UpdateChatByID(id string, chat model.Chat) {
	config.DB.Where("id = ?", id).Updates(&chat)
}
