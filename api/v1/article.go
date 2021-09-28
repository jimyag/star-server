package v1

import (
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
	"strconv"
)

func GetArticle(context *gin.Context) {
	pageSize, _ := strconv.Atoi(context.Query("pageSize"))
	pageIndex, _ := strconv.Atoi(context.Query("pageIndex"))
	if pageIndex == 0 {
		pageIndex = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	data, code := model.GetPaper(pageSize, pageIndex)
	utils.RequestDataOk(context, code, data)
}

func CreateArticle(context *gin.Context) {
	var paper model.Article
	_ = context.ShouldBindJSON(&paper)
	code := model.CreatePaper(&paper)
	if code == errmsg.ERROR {
		utils.RequestOk(context, code)
		return
	}
	utils.RequestDataOk(context, code, paper)
}
