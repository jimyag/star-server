package model

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
	"star-server/utils/errmsg"
)

type User struct {
	Model
	AvatarUrl string `gorm:"type:text;not null" json:"avatarUrl"`
	NickName  string `gorm:"type:varchar(20) ;not null" json:"nickName"`
	Gender    int    `gorm:"type:int;not null" json:"gender"`
	Province  string `gorm:"type:varchar(50);" json:"province"`
	City      string `gorm:"type:varchar(50);" json:"city"`
	Language  string `gorm:"type:varchar(30);" json:"language"`
	Country   string `gorm:"type:varchar(50);" json:"country"`
	Authority int    `gorm:"type:int;default 0" json:"authority"`
}

// CheckUser 用户是否存在
func CheckUser(id uint) (code int) {
	var user User
	db.Select("id").Where("id=?", id).First(&user)
	if user.ID > 0 {
		return errmsg.UserAlreadyExist //101
	}
	return errmsg.UserNotExist
}

func CreateUser(data *User) (code int) {
	if result := db.Create(&data); result.RowsAffected == 0 {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

func GetUser(id int) (User, int) {
	var user User
	if result := db.Limit(1).Where("id= ?", id).Find(&user); result.RowsAffected == 0 {
		return User{}, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 获得用户列表
func GetUsers(pageSize int, pageIndex int) []User {
	var users []User
	if result := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&users); result.RowsAffected == 0 {
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

// EditUser 只能更新处认证之外的其他信息
func EditUser(user *User) (code int) {
	var data User
	db.Model(&user).Select("authority").Find(&data)
	user.Authority = data.Authority
	if result := db.Model(&user).Updates(&user); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// UpdateUserAuth 只允许更新认证信息
func UpdateUserAuth(data *User) (code int) {
	var auth = make(map[string]int)
	auth["authority"] = data.Authority
	if result := db.Model(&data).Updates(auth); result.RowsAffected == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
