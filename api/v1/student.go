package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/model"
	"star-server/utils/errmsg"
)

func CreateStudent(context *gin.Context) {
	var data model.Student
	_ = context.ShouldBindJSON(&data)
	code := model.CreateStudent(&data)
	if code == errmsg.ERROR {
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
			"data": nil,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

func GetStudent(context *gin.Context) {
	studentId := context.Query("studentId")
	data, code := model.GetStudent(studentId)
	if code == errmsg.ERROR {
		code = errmsg.StudentNotExist
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

}
