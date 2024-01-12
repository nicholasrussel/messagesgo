package models

import "time"

type Message struct {
	ID           uint      `json:"id"`
	JapaneseText string    `json:"japanese_text"`
	EnglishText  string    `json:"english_text"`
	UserID       string    `json:"user_id"`
	CompanyID    string    `json:"company_id"`
	ChatRoomID   string    `json:"chat_room_id"`
	CreatedAt    time.Time `json:"created_at"`
}

