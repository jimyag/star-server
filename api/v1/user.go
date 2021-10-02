package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"star-server/middleware"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
	"strconv"
)

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
		utils.ResponseMsgOk(context, errmsg.ERROR, errMsg)
		return
	}
	avatarUrl := body["avatarUrl"]
	nickName := body["nickName"]
	gender, _ := strconv.Atoi(body["gender"])
	language := body["language"]
	city := body["city"]
	country := body["country"]
	province := body["province"]
	findToken, _ := model.UseTokenGetAuth(openid)

	// 该用户已经存在
	if findToken.ID != 0 {
		user, _ := model.GetUser(int(findToken.Uid))
		data["token"] = findToken.Token
		data["data"] = user
		utils.ResponseDataOk(context, errmsg.UserAlreadyExist, data)
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
	token, c := middleware.SetToken(int(user.ID))
	tokens := model.Authentication{
		Uid:    user.ID,
		Openid: openid,
		Token:  token,
	}

	// 设置token失败
	if c == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.TokenCreateError)
		return
	}

	isCreate := model.CreateTokens(&tokens)
	//创建token失败
	if isCreate == errmsg.InsertError {
		utils.ResponseOk(context, isCreate)
		return
	}

	// 最后成功
	data["token"] = token
	data["data"] = user
	utils.ResponseDataOk(context, code, data)
}

// GetUser 查询单个用户
func GetUser(context *gin.Context) {

	id, _ := strconv.Atoi(context.Param("uid"))
	if id < 1 {
		utils.ResponseOk(context, errmsg.ERROR)
		return
	}
	uid := context.Keys["uid"]
	uToken, err := model.UseUidGetAuth(uint(uid.(int)))
	if err != errmsg.SUCCESS {
		utils.ResponseOk(context, errmsg.UserNotExist)
		return
	}

	if (uint(id)) == uToken.Uid {
		data, code := model.GetUser(id)
		if code == errmsg.ERROR {
			utils.ResponseOk(context, errmsg.UserNotExist)
			return
		}
		utils.ResponseDataOk(context, code, data)
		return
	} else {
		utils.ResponseOk(context, errmsg.UserNotExist)
	}
}

// UpdateUser 编辑用户
func UpdateUser(context *gin.Context) {

	var user model.User
	_ = context.ShouldBindJSON(&user)
	var id, _ = strconv.Atoi(context.Param("uid"))
	user.ID = uint(id)

	uid := context.Keys["uid"]
	fmt.Println(uid, user.ID)
	if int(user.ID) == uid {
		// 编辑用户资料
		if model.EditUser(&user) == errmsg.ERROR {
			utils.ResponseOk(context, errmsg.UpdateError)
			return
		}
		newUser, _ := model.GetUser(int(user.ID))
		utils.ResponseDataOk(context, errmsg.SUCCESS, newUser)
	} else {
		//	id 和 openid不匹配
		utils.ResponseOk(context, errmsg.UserNotExist)
	}
}

func UpdateUserAuth(context *gin.Context) {
	var user model.User
	_ = context.ShouldBindJSON(&user)
	var id, _ = strconv.Atoi(context.Param("uid"))
	fmt.Println(user.ID)
	user.ID = uint(id)

	if id == context.Keys["uid"] {
		if model.UpdateUserAuth(&user) == errmsg.ERROR {
			utils.ResponseOk(context, errmsg.ERROR)
			return
		}
		newUser, _ := model.GetUser(id)
		utils.ResponseDataOk(context, errmsg.SUCCESS, newUser)
	} else {
		utils.ResponseOk(context, errmsg.UserNotExist)
	}
}
