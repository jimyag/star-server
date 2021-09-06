package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/model"
	"star-server/utils/errmsg"
	"strconv"
)

func CreateStuSect(context *gin.Context) {
	var stuSect model.StuSector
	var maps = make(map[string]string)
	_ = context.ShouldBindJSON(&maps)
	stuSect.Uid, _ = strconv.Atoi(context.Param("uid"))
	_, code := model.FindStuSector(&stuSect)
	if code == errmsg.SUCCESS {
		code = errmsg.StudentExist
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		context.Abort()
		return
	}
	stuSect.StudentId = maps["student_id"]
	stuSect.SectorName = maps["sector_name"]

	var stuName = maps["student_name"]

	fmt.Println(stuSect.Uid)
	// 学生和学号不匹配
	if !model.MatchStuNameAndId(stuName, stuSect.StudentId) {
		code = errmsg.StudentNotExist
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		context.Abort()
		return
	}
	// 密钥不正确
	var k, e = model.FindSectorKey(model.SectorKey{SectorName: stuSect.SectorName})
	if e == errmsg.ERROR && k.Key != maps["sector_key"] {
		code = errmsg.SectorKeyNotExist
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		context.Abort()
		return
	}
	code = model.CreateStuSect(&stuSect)
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
