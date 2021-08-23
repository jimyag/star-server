package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/middleware"
	"star-server/model"
	"star-server/utils/errmsg"
	"strconv"
)

//var code int

// UserExist 用户是否存在
func UserExist(context *gin.Context) {

}

// AddUser 添加用户
func AddUser(context *gin.Context) {
	//fmt.Println(context.GetHeader("token"))
	var code int
	data := make(map[string]interface{}) // 响应data
	body := make(map[string]string)      // json 参数

	_ = context.ShouldBindJSON(&body)
	openid := body["openid"]
	avatarUrl := body["avatarUrl"]
	nickName := body["nickName"]
	gender, _ := strconv.Atoi(body["gender"])
	language := body["language"]
	city := body["city"]
	country := body["country"]
	province := body["province"]
	// 必须要有openid
	if openid == "" {
		code = errmsg.ParameterConstraintError
		data["param"] = "openid"
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": data,
		})
		return
	}

	user := model.User{
		AvatarUrl: avatarUrl,
		NickName:  nickName,
		Gender:    gender,
		Language:  language,
		City:      city,
		Country:   country,
		Province:  province,
	}
	model.CreateUser(&user)
	token, c := middleware.SetToken(openid)
	fmt.Println(token)
	if c == errmsg.ERROR {
		code = errmsg.TokenCreateError
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": nil,
		})
		return
	}

	tokens := model.Tokens{
		Openid: openid,
		Token:  token,
	}
	isCreate := model.CreateTokens(&tokens)

	if isCreate == errmsg.ERROR {
		code = errmsg.TokenCreateError
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": nil,
		})
		return
	}

	data["token"] = token
	data["data"] = user
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

// GetUser 查询单个用户
func GetUser(context *gin.Context) {

}

// GetUsers 查询用户列表
func GetUsers(context *gin.Context) {

}

// EditUser 编辑用户
func EditUser(context *gin.Context) {

}