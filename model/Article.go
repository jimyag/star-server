package model

import (
	"github.com/lib/pq"
	"star-server/utils/errmsg"
)

type Article struct {
	Model
	Title         string         `gorm:"type:varchar(100)" json:"title"`
	ReleaseTime   string         `gorm:"type:varchar(100)" json:"release_time"`
	ReleaseSector string         `gorm:"type:varchar(100)" json:"release_sector"`
	Content       string         `gorm:"type:text" json:"content"`
	Author        string         `gorm:"type:varchar(40)" json:"author"`
	Attachment    pq.StringArray `gorm:"type:text[]" json:"attachment"`
}

func CreatePaper(data *Article) int {
	if result := db.Create(&data); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetPaper(pageSize int, pageIndex int) ([]Article, int) {
	var paperList []Article
	if result := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&paperList); result.RowsAffected == 0 {
		return nil, errmsg.ERROR
	}
	return paperList, errmsg.SUCCESS

}
