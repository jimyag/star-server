package model

import "github.com/jinzhu/gorm"

type Notice struct {
	gorm.Model
	Content string `gorm:"type:text"json:"content"`
	Remark  string `gorm:"type:text"json:"remark"`
}
