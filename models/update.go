package models

type Message struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}
