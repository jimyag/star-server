package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/model"
	"star-server/utils/errmsg"
	"strconv"
)

func getKey(sectorName string) string {
	return sectorName
}

func CreateStuSect(context *gin.Context) {
	var stuSect model.StuSector
	var maps = make(map[string]string)
	_ = context.ShouldBindJSON(&maps)
	stuSect.Uid, _ = strconv.Atoi(maps["id"])
	stuSect.StudentId = maps["student_id"]
	stuSect.SectorName = maps["sector_name"]

	if getKey(stuSect.SectorName) == maps["sector_key"] {

	}
	code := model.CreateStuSect(&stuSect)
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
