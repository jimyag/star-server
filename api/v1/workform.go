package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/model"
	"star-server/utils/errmsg"
	"strconv"
)

func CreateForm(ctx *gin.Context) {
	var data model.WorkForm
	_ = ctx.ShouldBindJSON(&data)
	var stuSe = model.StuSector{
		SectorName: data.SectorName,
		StudentId:  data.StudentId,
	}
	_, code := model.FindStuSectorUseSidSeName(&stuSe)
	if code == errmsg.ERROR {
		code = errmsg.StudentNotExist
		ctx.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		ctx.Abort()
		return
	}
	code = model.CreateForm(&data)
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

func UpdateForm(ctx *gin.Context) {
	var data model.WorkForm
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)
	code := model.UpdateForm(id, &data)
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
