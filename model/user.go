package model

import "time"

type User struct {
	Id        string    `json:"id"`
	FirstName string    `gorm:"type:varchar(255);column:firstName"json:"firstName"`
	LastName  string    `gorm:"type:varchar(255);column:lastName"json:"lastName"`
	Email     string    `gorm:"type:varchar(255);column:email"json:"email"`
	Password  string    `gorm:"type:varchar(255);column:password"json:"password"`
	IsActive  bool      `gorm:"column:isActive"json:"isActive"`
	SessionId string    `json:"sessionId"`
	CreatedAt time.Time `gorm:"column:createdAt"json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"json:"updatedAt"`
}

func NewUser() *User {
	var user User
	user.IsActive = true
	return &user
}
