package model

import "time"

// Define our message object
type Message struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `gorm:"column:createdAt"json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"json:"updatedAt"`
}

func NewMessage() *Message {
	var meessag Message
	return &meessag
}
