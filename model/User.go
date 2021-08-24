package model

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
	"star-server/utils/errmsg"
)

type User struct {
	gorm.Model
	AvatarUrl string `gorm:"type:text;not null" json:"avatar_url"`
	NickName  string `gorm:"type:varchar(255) ;not null" json:"nick_name"`
	Gender    int    `gorm:"type:int;not null" json:"gender"`
	Province  string `gorm:"type:varchar(50);" json:"province"`
	City      string `gorm:"type:varchar(50);" json:"city"`
	Language  string `gorm:"type:varchar(20);" json:"language"`
	Country   string `gorm:"type:varchar(50);" json:"country"`
}

func CheckUser(id uint) (code int) {
	var user User
	db.Select("id").Where("id=?", id).First(&user)
	if user.ID > 0 {
		return errmsg.UserAlreadyExist //101
	}
	return errmsg.UserNotExist
}

func CreateUser(data *User) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

func GetUser(id int) (User, int) {
	var user User
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 获得用户列表
func GetUsers(pageSize int, pageIndex int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

// ScryptPw 密码加密
func ScryptPw(passwd string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 25, 26, 125, 22, 11, 55, 99}
	HashPw, err := scrypt.Key([]byte(passwd), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatalln(err)
	}
	Fpw := base64.StdEncoding.EncodeToString(HashPw)
	return Fpw
}
