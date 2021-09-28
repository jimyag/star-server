package v1

import (
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
	"strconv"
)

func GetSector(context *gin.Context) {
	pageSize, _ := strconv.Atoi(context.Query("pageSize"))
	pageIndex, _ := strconv.Atoi(context.Query("pageIndex"))
	if pageIndex == 0 {
		pageIndex = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	data, code := model.GetSector(pageSize, pageIndex)
	utils.RequestDataOk(context, code, data)
}

func CreateSector(context *gin.Context) {
	var data model.Sector
	_ = context.ShouldBindJSON(&data)
	code := model.CreateSector(&data)
	if code == errmsg.ERROR {
		utils.RequestOk(context, code)
		return
	}
	utils.RequestDataOk(context, code, data)
}
