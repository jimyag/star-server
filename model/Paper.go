package model

import (
	"github.com/jinzhu/gorm"
	"star-server/utils/errmsg"
)

type Paper struct {
	gorm.Model
	Title         string `gorm:"type:varchar(255)" json:"title"`
	ReleaseTime   string `gorm:"type:varchar(50)" json:"release_time"`
	ReleaseSector string `gorm:"type:varchar(50)" json:"release_sector"`
	Content       string `gorm:"type:text" json:"content"`
	ReleaseName   string `gorm:"type:varchar(40)" json:"release_name"`
	Phone         string `gorm:"type:varchar(20)" json:"phone"`
	Attachment    string `gorm:"type:text" json:"attachment"`
}

func CreatePaper(data *Paper) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetPaper(pageSize int, pageIndex int) ([]Paper, int) {
	var paperList []Paper
	err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&paperList).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return paperList, errmsg.SUCCESS

}
