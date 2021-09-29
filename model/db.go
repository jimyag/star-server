package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"star-server/utils"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName))
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
	}

	// 禁用默认表的复数形式
	db.SingularTable(true)
	// 数据库的自动迁移
	db.AutoMigrate(&User{},
		&Notice{},
		&Article{},
		&Schedule{},
		&Sector{},
		&Student{},
		&WorkForm{},
		&Authentication{},
		&SectorKey{},
		&StuSector{})

	// 设置连接池中最大的闲置连接数
	db.DB().SetMaxIdleConns(10)

	// 设置数据库的最大连接数
	db.DB().SetMaxOpenConns(100)

	// 设置连接的最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

}
