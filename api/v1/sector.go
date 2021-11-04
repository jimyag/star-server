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
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 5
	}
	data, code := model.GetSector(pageSize, pageIndex)
	utils.ResponseDataOk(context, code, data)
}

func CreateSector(context *gin.Context) {
	var data model.Sector
	_ = context.ShouldBindJSON(&data)
	if data.SectorName == "" || data.SectorInformation == "" || data.Address == "" || data.Phone == "" || data.Supervisor == "" {
		utils.ResponseOk(context, errmsg.ParameterError)
		return
	}
	code := model.CreateSector(&data)
	if code == errmsg.ERROR {
		utils.ResponseOk(context, code)
		return
	}
	utils.ResponseDataOk(context, code, data)
}
