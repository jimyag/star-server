package v1

import (
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
	"strconv"
	"time"
)

func CreateForm(context *gin.Context) {
	var data model.WorkForm
	_ = context.ShouldBindJSON(&data)
	var stuSe = model.StuSector{
		SectorName: data.SectorName,
		StudentId:  data.StudentId,
	}
	_, code := model.FindStuSectorUseSidSeName(&stuSe)
	if code == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	data.BeginTime = int(time.Now().Unix())
	code = model.CreateForm(&data)
	utils.ResponseDataOk(context, code, data)
}

func UpdateForm(ctx *gin.Context) {
	var data model.WorkForm
	uid, _ := strconv.Atoi(ctx.Param("uid"))
	_ = ctx.ShouldBindJSON(&data)

	stu_sector, code := model.FindStuSectorUseUid(model.StuSector{Uid: uid})
	if code == errmsg.ERROR {
		utils.ResponseOk(ctx, errmsg.StudentNotExist)
		return
	}
	if stu_sector.SectorName != data.SectorName {
		utils.ResponseOk(ctx, errmsg.StudentNotExist)
		return
	}
	_, code = model.GetFormUseId(data.ID)
	if code == errmsg.ERROR {
		utils.ResponseMsgOk(ctx, errmsg.ERROR, "这条工作记录错误")
		return
	}
	data.EndTime = int(time.Now().Unix())
	utils.ResponseOk(ctx, model.UpdateForm(&data))

}

func GetStuForm(context *gin.Context) {
	var studentId = context.Param("student_id")
	forms, err := model.GetForms(studentId)
	utils.ResponseDataOk(context, err, forms)
}
