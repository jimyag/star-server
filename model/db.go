package model

import (
	"fmt"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"star-server/utils"
)

var db *gorm.DB
var err error

func InitDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s", utils.DbHost, utils.DbUser, utils.DbPassWord, utils.DbName, utils.DbPort, utils.DbTimeZone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
		return
	}
	// 数据库的自动迁移
	err = db.AutoMigrate(&User{},
		&Notice{},
		&Article{},
		&Schedule{},
		&Sector{},
		&Student{},
		&WorkForm{},
		&Authentication{},
		&SectorKey{},
		&StuSector{},
		&Mytest{})
	if err != nil {
		fmt.Println("数据库创建失败", err)
		return
	}

	// 设置连接池中最大的闲置连接数
	//a, err = db.DB()
	//a.SetMaxIdleConns(10)
	//db.DB().SetMaxIdleConns(10)

	// 设置数据库的最大连接数
	//a.SetMaxOpenConns(100)
	//db.DB().SetMaxOpenConns(100)

	// 设置连接的最大可复用时间
	//a.SetConnMaxLifetime(10 * time.Second)
	//db.DB().SetConnMaxLifetime(10 * time.Second)

}
