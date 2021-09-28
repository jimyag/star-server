package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type Notice struct {
	gorm.Model
	Content string `gorm:"type:varchar(100);" json:"content"`
	Remark  string `gorm:"type:varchar(100)" json:"remark"`
}

func GetNotice() (int, []Notice) {
	var notices []Notice
	err := db.Order("created_at desc").Limit(3).Find(&notices).Error
	if err != nil {
		return errmsg.ERROR, nil
	}
	return errmsg.SUCCESS, notices
}

func AddNotice(data *Notice) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}
