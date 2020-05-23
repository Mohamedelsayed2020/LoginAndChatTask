package services

import (
	"LoginAndChatTask/api/server"
	"LoginAndChatTask/core/common"
	"LoginAndChatTask/core/model"
	"strings"
)

type User model.User

type UserLogin struct {
	Email     string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string `json:"password"`
	SessionId string `json:"sessionId"`
	IsActive  bool   `json:"isActive"`
}
func (self *UserLogin) Format() *UserLogin {
	self.Email = strings.ToLower(self.Email)
	return self
}

func (self *UserLogin) ValidateLogin() (string, *User) {
	user := &User{}

	if self.Email != "" && self.Password != "" {
		_, user = user.FindByEmail(self.Email)
		if user == nil || user.Email == "" {
			return "UserNotFound", nil
		} else if self.Email != user.Email {
			return "LoginFailed", nil
		}
		password := common.CheckPasswordHash(self.Password, user.Password)
		if password == false {
			return "LoginFailed", nil
		}
		if user.IsActive == false {
			return "EmptyLoginFields", nil
		}
	} else {
		return "UserNotFound", nil
	}
	user.Password = ""
	return "", user
}

func (self *User) FindByEmail(mail string) (error, *User) {
	newUser := &User{}
	queryResult := server.Conn().Where(&User{Email: mail}).First(newUser)
	if queryResult.Error != nil {
		return queryResult.Error, nil
	} else {
		return nil, newUser
	}
}

func (self *User) UpdateUser(email string) interface{} {
	if queryResult := server.Conn().Model(&self).Where("email = ?", email).Updates(map[string]interface{}{
		"sessionId":        &self.SessionId,
	}); queryResult.Error != nil {
		return queryResult.Error.Error()
	}
	return nil
}
