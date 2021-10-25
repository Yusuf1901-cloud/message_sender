package models

type Message struct {
	ChatID   int64  `json:"chat_id"`
	Priority string `json:"priority"`
	Text     string `json:"text"`
}
