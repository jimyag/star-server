package v1

import (
	"fmt"
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
	_ = context.ShouldBindJSON(&schedule)
	var code int
	// 验证学号和姓名是否正确
	if !model.MatchStuNameAndId(schedule.StudentName, schedule.StudentId) {
		code = errmsg.StudentNotExist
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		context.Abort()
		return
	}
	// 验证是否有这个部门
	sector, code := model.UseNameGetSector(schedule.SectorName)
	if code == errmsg.ERROR {
		code = errmsg.SectorNotExist
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		context.Abort()
		return
	}
	// 验证地点是否正确
	if sector.Address != schedule.Address {
		code = errmsg.SectorNotExist
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  "地址错误",
		})
		context.Abort()
		return
	}
	fmt.Println(schedule.DayOfWeek)
	if schedule.DayOfWeek > 5 || schedule.DayOfWeek < 1 || schedule.CourseIndex < 1 || schedule.CourseIndex > 4 {
		context.JSON(http.StatusOK, gin.H{
			"code": errmsg.ERROR,
			"msg":  "time error",
			"data": nil,
		})
		context.Abort()
		return
	}
	// 创建一条记录
	code = model.CreateSchedule(&schedule)
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
		"data": schedule,
	})
}
