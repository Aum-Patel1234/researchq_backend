package models

import "gorm.io/gorm"

type ChatModel struct {
	gorm.Model

	chat   string `json:"chat"`
	userId uint64 `json:"user_id"`
}

type ConversationMode struct {
	gorm.Model

	userId uint64
	title  string
	chats  []ChatModel
}
