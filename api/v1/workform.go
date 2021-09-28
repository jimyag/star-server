package v1

import (
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
	"strconv"
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
		utils.RequestOk(context, errmsg.StudentNotExist)
		return
	}
	utils.RequestDataOk(context, model.CreateForm(&data), data)
}

func UpdateForm(ctx *gin.Context) {
	var data model.WorkForm
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)
	utils.RequestOk(ctx, model.UpdateForm(id, &data))

}

func GetStuForm(context *gin.Context) {
	var studentId = context.Param("student_id")
	forms, err := model.GetForm(studentId)
	utils.RequestDataOk(context, err, forms)
}
