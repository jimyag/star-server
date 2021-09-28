package v1

import (
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
)

func CreateStudent(context *gin.Context) {
	var data model.Student
	_ = context.ShouldBindJSON(&data)
	_, code := model.GetStudent(data.StudentId)
	if code == errmsg.SUCCESS {
		utils.RequestOk(context, code)
		return
	}
	code = model.CreateStudent(&data)
	utils.RequestDataOk(context, code, data)
}

func GetStudent(context *gin.Context) {
	studentId := context.Param("student_id")
	data, code := model.GetStudent(studentId)
	utils.RequestDataOk(context, code, data)

}
