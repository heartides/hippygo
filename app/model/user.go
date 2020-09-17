package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string
	PassWord string
	NickName string
	Status   uint8
	Avatar   string
}
