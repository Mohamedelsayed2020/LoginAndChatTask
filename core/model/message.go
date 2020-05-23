package model

import (
	"github.com/google/uuid"
	"time"
)

// Define our message object
type Message struct {
	Id        uuid.UUID
	Message   string    `json:"message"`
	CreatedAt time.Time `gorm:"column:createdAt"json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"json:"updatedAt"`
}

func NewMessage() *Message {
	var message Message
	return &message
}
