package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/model"
	"star-server/utils/errmsg"
)

func AddNotice(context *gin.Context) {
	var notice model.Notice
	var code int
	_ = context.ShouldBindJSON(&notice)
	err := model.AddNotice(&notice)
	if err == errmsg.ERROR {
		code = errmsg.ERROR
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": nil,
		})
	}
	code = errmsg.SUCCESS
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": notice,
	})

}

func GetNotice(context *gin.Context) {
	notices := model.GetNotice()
	code := errmsg.SUCCESS
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": notices,
	})
}
