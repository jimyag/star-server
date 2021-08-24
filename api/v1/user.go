package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/middleware"
	"star-server/model"
	"star-server/utils/errmsg"
	"strconv"
)

// UserExist 用户是否存在
func UserExist(context *gin.Context) {

}

// AddUser 添加用户
func AddUser(context *gin.Context) {
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

	tokens := model.Tokens{
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

	}
	opid := context.Keys["openid"]
	var openid = opid.(string)
	uToken := model.UseOpenidGetUid(openid)
	if (uint(id)) == uToken.Uid {

	}
	//data, code := model.GetUser(id)
	//if code == errmsg.ERROR {
	//	code = errmsg.UserNotExist
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": code,
	//		"msg":  errmsg.GetErrMsg(code),
	//		"data": nil,
	//	})
	//	return
	//}
	//context.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  errmsg.GetErrMsg(code),
	//	"data": data,
	//	"openid":context.Keys,
	//})
	//id := context.Param("id")
	//tokens := model.UseOpenidGetUid(id)
	//if tokens.ID == 0 {
	//	fmt.Println("不存在")
	//} else {
	//	fmt.Println(tokens.ID)
	//}
}

// GetUsers 查询用户列表
func GetUsers(context *gin.Context) {

}

// EditUser 编辑用户
func EditUser(context *gin.Context) {

}
