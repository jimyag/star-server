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
	if data.StudentId == "" || data.StudentName == "" {
		utils.ResponseOk(context, errmsg.ParameterError)
		return
	}
	// 验证学生是否存在
	_, code := model.GetStudent(data.StudentId)
	if code == errmsg.SUCCESS {
		utils.ResponseOk(context, errmsg.StudentExist)
		return
	}
	code = model.CreateStudent(&data)
	utils.ResponseDataOk(context, code, data)
}

func GetStudent(context *gin.Context) {
	studentId := context.Param("student_id")
	data, code := model.GetStudent(studentId)
	if code == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.StudentNotExist)
		return
	}
	utils.ResponseDataOk(context, code, data)
}

func UpdateStudent(context *gin.Context) {
	var student model.Student
	_ = context.ShouldBindJSON(&student)

	if student.StudentId == "" {
		utils.ResponseOk(context, errmsg.ParameterError)
		return
	}

	newStudent, err := model.UpdateStudent(student)
	if err == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.ERROR)
		return
	}
	utils.ResponseDataOk(context, errmsg.SUCCESS, newStudent)
}
