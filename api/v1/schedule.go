package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/model"
	"star-server/utils/errmsg"
)

func GetSchedule(context *gin.Context) {
	sectorName := context.Query("sectorName")
	data, code := model.GetSectorSchedule(sectorName)
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

func AddOneRecord(context *gin.Context) {
	var schedule model.Schedule
	_ = context.ShouldBindJSON(schedule)
	code := model.CreateSchedule(&schedule)
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
		"data": schedule,
	})
}
