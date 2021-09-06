package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/middleware"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
	"star-server/utils/verify"
	"strconv"
)

// UserExist 用户是否存在
func UserExist(context *gin.Context) {

}

//AddUser 添加用户
func AddUser(context *gin.Context) {
	var code int
	data := make(map[string]interface{}) // 响应data
	body := make(map[string]string)      // json 参数

	_ = context.ShouldBindJSON(&body)
	// 在后端验证openid
	openid, errMsg := utils.GetOpenid(body["code"])
	// openid生成错误
	if errMsg != "" {
		context.JSON(http.StatusOK, gin.H{
			"code": errmsg.ERROR,
			"msg":  errMsg,
			"data": nil,
		})
		context.Abort()
		return
	}
	avatarUrl := body["avatarUrl"]
	nickName := body["nickName"]
	gender, _ := strconv.Atoi(body["gender"])
	language := body["language"]
	city := body["city"]
	country := body["country"]
	province := body["province"]
	findToken := model.UseOpenidGetUid(openid)

	// 该用户已经存在
	if findToken.ID != 0 {
		user, _ := model.GetUser(int(findToken.Uid))
		data["token"] = findToken.Token
		data["data"] = user
		code = errmsg.UserAlreadyExist
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": data,
		})
		context.Abort()
		return
	}

	token, c := middleware.SetToken(openid)
	// 设置token失败
	if c == errmsg.ERROR {
		code = errmsg.TokenCreateError
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": nil,
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
	tokens := model.Authentication{
		Uid:    user.ID,
		Openid: openid,
		Token:  token,
	}
	isCreate := model.CreateTokens(&tokens)
	//创建token失败
	if isCreate == errmsg.ERROR {
		code = errmsg.TokenCreateError
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": nil,
		})
		return
	}
	// 最后成功
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
	//使用uid和token查找该用户
	//var code int

	id, _ := strconv.Atoi(context.Param("id"))
	if id < 1 {
		context.JSON(http.StatusOK, gin.H{
			"code": errmsg.ERROR,
			"msg":  errmsg.GetErrMsg(errmsg.ERROR),
			"data": nil,
		})
		context.Abort()
		return
	}
	opid := context.Keys["openid"]
	var openid = opid.(string)
	uToken := model.UseOpenidGetUid(openid)
	if (uint(id)) == uToken.Uid {
		data, code := model.GetUser(id)
		if code == errmsg.ERROR {
			code = errmsg.UserNotExist
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
				"data": nil,
			})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": data,
		})
		context.Abort()
		return
	} else {
		code := errmsg.UserNotExist
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": nil,
		})
	}
}

// GetUsers 查询用户列表
func GetUsers(context *gin.Context) {

}

// EditUser 编辑用户
func EditUser(context *gin.Context) {
	var user model.User
	_ = context.ShouldBindJSON(&user)
	var id, _ = strconv.Atoi(context.Param("id"))
	user.ID = uint(id)
	if verify.MatchIdToken(user.ID, context.Keys["openid"].(string)) {
		// 编辑用户资料
		if model.EditUser(&user) == errmsg.ERROR {
			context.JSON(http.StatusOK, gin.H{
				"code": errmsg.ERROR,
				"msg":  errmsg.GetErrMsg(errmsg.ERROR),
			})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code": errmsg.SUCCESS,
			"msg":  errmsg.GetErrMsg(errmsg.SUCCESS),
		})
		//	id 和 openid不匹配
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": errmsg.UserNotExist,
			"msg":  errmsg.GetErrMsg(errmsg.UserNotExist),
		})
	}
}

func UpdateUserAuth(context *gin.Context) {
	var user model.User
	_ = context.ShouldBindJSON(&user)
	var id, _ = strconv.Atoi(context.Param("id"))
	user.ID = uint(id)
	fmt.Println(user.ID)
	if verify.MatchIdToken(user.ID, context.Keys["openid"].(string)) {
		if model.UpdateUserAuth(&user) == errmsg.ERROR {
			context.JSON(http.StatusOK, gin.H{
				"code": errmsg.ERROR,
				"msg":  errmsg.GetErrMsg(errmsg.ERROR),
			})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code": errmsg.SUCCESS,
			"msg":  errmsg.GetErrMsg(errmsg.SUCCESS),
		})

	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": errmsg.UserNotExist,
			"msg":  errmsg.GetErrMsg(errmsg.UserNotExist),
		})
	}
}
