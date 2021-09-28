package v1

import (
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
)

func GetSchedule(context *gin.Context) {
	sectorName := context.Param("sector_name")
	data, code := model.GetSectorSchedule(sectorName)
	utils.RequestDataOk(context, code, data)
}

func AddOneRecord(context *gin.Context) {
	var schedule model.Schedule
	_ = context.ShouldBindJSON(&schedule)
	var code int
	// 验证学号和姓名是否正确
	if !model.MatchStuNameAndId(schedule.StudentName, schedule.StudentId) {
		utils.RequestOk(context, errmsg.StudentNotExist)
		return
	}
	// 验证是否有这个部门
	sector, code := model.UseNameGetSector(schedule.SectorName)
	if code == errmsg.ERROR {
		utils.RequestOk(context, errmsg.SectorNotExist)
		return
	}
	// 验证地点是否正确
	if sector.Address != schedule.Address {
		utils.RequestOk(context, errmsg.SectorAddressNotExist)
		return
	}
	if schedule.DayOfWeek > 5 || schedule.DayOfWeek < 1 || schedule.CourseIndex < 1 || schedule.CourseIndex > 4 {
		utils.RequestMsgOk(context, errmsg.ERROR, "time error")
		return
	}
	// 创建一条记录
	code = model.CreateSchedule(&schedule)
	if code == errmsg.ERROR {
		utils.RequestOk(context, code)
		return
	}
	utils.RequestDataOk(context, code, schedule)
}
