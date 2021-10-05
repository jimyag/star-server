package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"star-server/middleware"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
	"strconv"
)

type addUser struct {
	model.User
	Code string `json:"code"`
}

//AddUser 添加用户
func AddUser(context *gin.Context) {

	// 绑定前端发过来的json
	var body addUser
	_ = context.ShouldBindJSON(&body)

	//在后端验证openid
	openid, errMsg := utils.GetOpenid(body.Code)

	//openid生成错误
	if errMsg != "" {
		utils.ResponseMsgOk(context, errmsg.ERROR, errMsg)
		return
	}

	var data = make(map[string]interface{})
	//判断用户是否已经存在了
	authUser, err := model.UseOpenidGetAuth(openid)
	if err == errmsg.SUCCESS {
		user, _ := model.GetUser(int(authUser.Uid))
		data["data"] = user
		data["token"] = authUser.Token
		utils.ResponseDataOk(context, errmsg.UserAlreadyExist, data)
		return
	}
	var user model.User
	copyErr := copier.Copy(&user, &body)
	if copyErr != nil {
		utils.ResponseOk(context, errmsg.ERROR)
		return
	}
	//用户不存在进行注册
	err = model.CreateUser(&user)

	// 用户创建失败了
	if err == errmsg.ERROR {
		utils.ResponseOk(context, err)
		return
	}

	token, err := middleware.SetToken(int(user.ID))
	// token 设置失败
	if err == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.TokenCreateError)
		return
	}

	// 插入生成的token信息
	newToken := model.Authentication{
		Uid:    user.ID,
		Openid: openid,
		Token:  token,
	}
	err = model.CreateTokens(&newToken)

	//插入token认证失败
	if err == errmsg.InsertError {
		utils.ResponseOk(context, err)
		return
	}

	// 最后成功
	data["token"] = token
	data["data"] = user
	utils.ResponseDataOk(context, err, data)
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
