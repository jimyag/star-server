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
	if pageIndex < 0 || pageSize < 0 {
		utils.ResponseOk(context, errmsg.ParameterError)
		return
	}
	if pageIndex == 0 {
		pageIndex = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	data, code := model.GetPaper(pageSize, pageIndex)
	utils.ResponseDataOk(context, code, data)
}

func CreateArticle(context *gin.Context) {
	var paper model.Article
	_ = context.ShouldBindJSON(&paper)
	if paper.Title == "" || paper.Content == "" {
		utils.ResponseOk(context, errmsg.ParameterError)
		return
	}

	code := model.CreatePaper(&paper)
	if code == errmsg.ERROR {
		utils.ResponseOk(context, code)
		return
	}
	utils.ResponseDataOk(context, code, paper)
}
