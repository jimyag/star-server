package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type Tokens struct {
	gorm.Model
	Openid string `gorm:"type:text;not null" json:"openid"`
	Token  string `gorm:"type:text;not null" json:"token"`
}

func CreateTokens(data *Tokens) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}
