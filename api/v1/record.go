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
	uid, _ := strconv.Atoi(context.Param("uid"))
	if realUid := context.Keys["uid"]; realUid != uid {
		utils.ResponseOk(context, errmsg.ERROR)
		return
	}
	stuSect, err := model.FindStuSectorUseUid(model.StuSector{Uid: uid})
	if err == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	var data model.Record
	_ = context.ShouldBindJSON(&data)
	if stuSect.StudentId != data.StudentId {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	var stuSe = model.StuSector{
		SectorName: data.SectorName,
		StudentId:  data.StudentId,
	}
	_, code := model.FindStuSectorUseSidSeName(stuSe)
	if code == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	data.BeginTime = int(time.Now().Unix())
	code = model.CreateForm(&data)
	utils.ResponseDataOk(context, code, data)
}

func UpdateForm(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Param("uid"))
	if realUid := context.Keys["uid"]; realUid != uid {
		utils.ResponseOk(context, errmsg.ERROR)
		return
	}
	stu_sector, code := model.FindStuSectorUseUid(model.StuSector{Uid: uid})
	if code == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	var data model.Record
	_ = context.ShouldBindJSON(&data)
	if stu_sector.StudentId != data.StudentId {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	if stu_sector.SectorName != data.SectorName {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	rid, _ := strconv.Atoi(context.Param("rid"))
	record, code := model.GetFormUseId(uint(rid))
	if code == errmsg.ERROR {
		utils.ResponseMsgOk(context, errmsg.ERROR, "这条工作记录错误")
		return
	}
	if record.EndTime != 0 {
		utils.ResponseMsgOk(context, errmsg.ERROR, "已经更新此条工作记录")
		return
	}
	data.EndTime = int(time.Now().Unix())
	data.ID = uint(rid)
	utils.ResponseOk(context, model.UpdateForm(&data))

}

func GetStuForm(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Param("uid"))
	if realUid := context.Keys["uid"]; realUid != uid {
		utils.ResponseOk(context, errmsg.ERROR)
		return
	}
	stu_sector, code := model.FindStuSectorUseUid(model.StuSector{Uid: uid})
	if code == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	forms, err := model.GetForms(stu_sector.StudentId)
	utils.ResponseDataOk(context, err, forms)
}
