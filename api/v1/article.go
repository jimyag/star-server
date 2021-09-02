package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/model"
	"star-server/utils/errmsg"
	"strconv"
)

func GetArticle(context *gin.Context) {
	pageSize, _ := strconv.Atoi(context.Query("pageSize"))
	pageIndex, _ := strconv.Atoi(context.Query("pageIndex"))
	var code int
	if pageIndex == 0 {
		pageIndex = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	data, code := model.GetPaper(pageSize, pageIndex)

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

func CreateArticle(context *gin.Context) {
	var paper model.Article
	_ = context.ShouldBindJSON(&paper)
	code := model.CreatePaper(&paper)
	if code == errmsg.ERROR {
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
		"data": paper,
	})
}
