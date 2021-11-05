package v1

import (
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
)

func GetSchedule(context *gin.Context) {
	sectorName := context.Param("sector_name")
	// 验证这个同学是否在这个部门
	uid := context.Keys["uid"]
	stuSector, code := model.FindStuSectorUseUid(model.StuSector{Uid: uid.(int)})
	if code == errmsg.ERROR {
		utils.ResponseMsgOk(context, errmsg.ERROR, "该学生没有添加部门")
		return
	}
	if stuSector.SectorName != sectorName {
		utils.ResponseMsgOk(context, errmsg.ERROR, "该学生没有添加该部门")
		return
	}
	data, code := model.GetSectorSchedule(sectorName)
	utils.ResponseDataOk(context, code, data)
}

func AddOneRecord(context *gin.Context) {
	sectorName := context.Param("sector_name")
	// 验证是否有这个部门
	_, code := model.UseNameGetSector(sectorName)
	if code == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.SectorNotExist)
		return
	}
	var schedule model.Schedule
	_ = context.ShouldBindJSON(&schedule)
	// 验证学号和姓名是否正确
	if !model.MatchStuNameAndId(schedule.StudentName, schedule.StudentId) {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	// 这个同学是否加这个部门了
	_, code = model.FindStuSectorUseSidSeName(model.StuSector{StudentId: schedule.StudentId, SectorName: sectorName})
	if code == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	schedule.SectorName = sectorName
	if schedule.DayOfWeek > 5 || schedule.DayOfWeek < 1 || schedule.CourseIndex < 1 || schedule.CourseIndex > 4 {
		utils.ResponseMsgOk(context, errmsg.ERROR, "time error")
		return
	}
	// 创建一条记录
	if !model.ScheduleRecordEquals(schedule) {
		utils.ResponseMsgOk(context, errmsg.ERROR, "已有一条该地点时间的值班记录了")
		return
	}
	code = model.CreateSchedule(&schedule)
	if code == errmsg.ERROR {
		utils.ResponseOk(context, code)
		return
	}
	utils.ResponseDataOk(context, code, schedule)
}
