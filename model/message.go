package model

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func NewMessage() *Message {
	var meessag Message
	return &meessag
}
