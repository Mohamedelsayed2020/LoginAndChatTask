package model

import "github.com/jinzhu/gorm"

type Session struct {
	gorm.Model
	UserId    string
	SessionId string
	Ip        string
}
