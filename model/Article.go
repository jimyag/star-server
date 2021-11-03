package model

import (
	"star-server/utils/errmsg"
)

type Article struct {
	Model
	Title         string `gorm:"type:varchar(100)" json:"title"`
	ReleaseTime   string `gorm:"type:varchar(100)" json:"release_time"`
	ReleaseSector string `gorm:"type:varchar(100)" json:"release_sector"`
	Content       string `gorm:"type:text" json:"content"`
	ReleaseName   string `gorm:"type:varchar(40)" json:"release_name"`
	Phone         string `gorm:"type:char(11)" json:"phone"`
	Attachment    string `gorm:"type:text" json:"attachment"`
}

func CreatePaper(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetPaper(pageSize int, pageIndex int) ([]Article, int) {
	var paperList []Article
	err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&paperList).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return paperList, errmsg.SUCCESS

}
