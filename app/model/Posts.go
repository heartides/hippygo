package model

import "github.com/jinzhu/gorm"

type Posts struct {
	gorm.Model
	Title string
}
