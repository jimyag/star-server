package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
	"strconv"
)

func CreateStuSect(context *gin.Context) {
	var stuSect model.StuSector
	var maps = make(map[string]string)
	_ = context.ShouldBindJSON(&maps)
	stuSect.Uid, _ = strconv.Atoi(context.Param("uid"))

	uid := context.Keys["uid"]
	if uid != stuSect.Uid {
		utils.ResponseOk(context, errmsg.ERROR)
		return
	}
	sectorName := context.Param("sector_name")
	_, code := model.UseNameGetSector(sectorName)
	if code == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.SectorNotExist)
		return
	}
	stuSect.StudentId = maps["student_id"]
	stuSect.SectorName = sectorName
	real_key, code := model.UseSectorNameFindSectorKey(stuSect.SectorName)
	if code == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.SectorKeyNotExist)
		return
	}
	if real_key != maps["sector_key"] {
		utils.ResponseOk(context, errmsg.SectorKeyNotExist)
		return
	}
	_, code = model.FindStuSectorUseSidSeName(stuSect)
	if code == errmsg.SUCCESS {
		utils.ResponseOk(context, errmsg.StudentExist)
		return
	}

	var stuName = maps["student_name"]

	//fmt.Println(stuSect.Uid)
	// 学生和学号不匹配
	if !model.MatchStuNameAndId(stuName, stuSect.StudentId) {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	// 密钥不正确
	var k, e = model.FindSectorKey(model.SectorKey{SectorName: stuSect.SectorName})
	//fmt.Println(k.Key)
	if e == errmsg.ERROR && k.Key != maps["sector_key"] {
		utils.ResponseOk(context, errmsg.SectorKeyNotExist)
		return
	}
	fmt.Println("111")
	utils.ResponseOk(context, model.CreateStuSect(&stuSect))
}

func FindStuSector(context *gin.Context) {
	var uid, _ = strconv.Atoi(context.Param("uid"))
	if realUid := context.Keys["uid"]; realUid != uid {
		utils.ResponseOk(context, errmsg.ERROR)
		return
	}
	var stuS, err = model.FindStuSectorUseUid(model.StuSector{Uid: uid})
	if err == errmsg.ERROR {
		utils.ResponseMsgOk(context, err, "该同学没有加入部门")
		return
	}
	var stu, _ = model.GetStudent(stuS.StudentId)
	var data = make(map[string]interface{})
	data["student_id"] = stu.StudentId
	data["student_name"] = stu.StudentName
	data["sector_name"] = stuS.SectorName
	utils.ResponseDataOk(context, errmsg.SUCCESS, data)

}
