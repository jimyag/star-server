package model

import (
	"star-server/utils/errmsg"
)

type Notice struct {
	Model
	Content string `gorm:"type:varchar(100);" json:"content"`
	Remark  string `gorm:"type:varchar(100)" json:"remark"`
}

func GetNotice() (int, []Notice) {
	var notices []Notice
	if result := db.Order("created_at desc").Limit(3).Find(&notices); result.RowsAffected == 0 {
		return errmsg.ERROR, nil
	}
	return errmsg.SUCCESS, notices
}

func AddNotice(data *Notice) int {
	if result := db.Create(&data); result.RowsAffected == 0 {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}
