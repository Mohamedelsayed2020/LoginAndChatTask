package model

import (
	"LoginAndChatTask/server"
	"github.com/google/uuid"
)

func CreateSession(userId string) (error, string) {
	var session Session
	session.UserId = userId
	session.SessionId = uuid.New().String()
	if queryResult := server.Conn().Create(&session); queryResult.Error != nil {
		return queryResult.Error, ""
	}
	return nil, session.SessionId
}

func (self *User) IsSessionExist(id string) bool {
	session := Session{}
	queryResult := server.Conn().Where(&Session{SessionId: id}).Find(&session)
	if queryResult.Error != nil {
		return false
	}
	return true
}
