package model

import "github.com/jinzhu/gorm"

type Tokens struct {
	gorm.Model
	Openid string `gorm:"type:varchar(100);not null" json:"openid"`
	Token  string `gorm:"type:varchar(100);not null" json:"token"`
}
