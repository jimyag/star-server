package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	AvatarUrl string `gorm:"type:text;not null" json:"avatar_url"`
	NickName  string `gorm:"type:varchar(255) ;not null" json:"nick_name"`
	Gender    int    `gorm:"type:int;not null" json:"gender"`
	Province  string `gorm:"type:varchar(50);" json:"province"`
	City      string `gorm:"type:varchar(50);" json:"city"`
	Language  string `gorm:"type:varchar(20);" json:"language"`
}
