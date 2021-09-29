package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		utils.ResponseOk(context, errmsg.ERROR)
		return
	}
	avatarUrl := body["avatarUrl"]
	nickName := body["nickName"]
	gender, _ := strconv.Atoi(body["gender"])
	language := body["language"]
	city := body["city"]
	country := body["country"]
	province := body["province"]
	findToken, err := model.UseTokenGetAuth(openid)
	if err != errmsg.SUCCESS {
		utils.ResponseOk(context, err)
		return
	}

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

	//使用uid和token查找该用户
	//var code int

	// a GET request to /user/john
	// id := c.Param("id") id == "john"
	id, _ := strconv.Atoi(context.Param("uid"))
	if id < 1 {
		utils.ResponseOk(context, errmsg.ERROR)
		return
	}

	uidInterface := context.Keys["uid"]
	uid := uidInterface.(int)
	uToken, err := model.UseUidGetAuth(uint(uid))
	if err != errmsg.SUCCESS {
		utils.ResponseOk(context, err)
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

// EditUser 编辑用户
func EditUser(context *gin.Context) {

	var user model.User
	_ = context.ShouldBindJSON(&user)
	var id, _ = strconv.Atoi(context.Param("uid"))
	user.ID = uint(id)
	if user.ID == context.Keys["uid"].(uint) {
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
	var id, _ = strconv.Atoi(context.Param("uid"))
	user.ID = uint(id)
	//fmt.Println(user.ID)
	if user.ID == context.Keys["uid"].(uint) {
		if model.UpdateUserAuth(&user) == errmsg.ERROR {
			utils.ResponseOk(context, errmsg.ERROR)
			return
		}

		utils.ResponseOk(context, errmsg.SUCCESS)

	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": errmsg.UserNotExist,
			"msg":  errmsg.GetErrMsg(errmsg.UserNotExist),
		})
		utils.ResponseOk(context, errmsg.UserNotExist)
	}
}
