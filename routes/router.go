package routes

import (
	"github.com/gin-gonic/gin"
	v1 "star-server/api/v1"
	"star-server/middleware"
	"star-server/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	authV1 := r.Group("api/v1")
	authV1.Use(middleware.JwtToken())
	{
		//用户模块
		authV1.PUT("user/:uid/update", v1.EditUser)
		authV1.PUT("user/:uid/updateauth", v1.UpdateUserAuth)
		authV1.GET("user/:uid", v1.GetUser)

		// 紧急通知
		authV1.POST("notice/add", v1.CreateNotice)

		// 文章
		authV1.POST("article/add", v1.CreateArticle)

		//部门
		authV1.POST("sector/add", v1.CreateSector)
		authV1.POST("sector/:uid/join", v1.CreateStuSect)
		authV1.GET("sector/:uid/find", v1.FindStuSector)
		//学生
		authV1.POST("student/add", v1.CreateStudent)
		authV1.GET("student/:student_id", v1.GetStudent)
		//值班表
		authV1.GET("schedule/:sector_name", v1.GetSchedule)
		authV1.POST("schedule/add", v1.AddOneRecord)

		// 工作记录
		authV1.POST("workform/add", v1.CreateForm)
		authV1.PUT("workform/update/:uid", v1.UpdateForm)
		authV1.GET("workform/get/:student_id", v1.GetStuForm)
		// todo 获得用户的工作情况

		//部门邀请码
		authV1.POST("sectorkey/add", v1.CreateSectorKey)

	}
	routerV1 := r.Group("api/v1")
	{
		// 注册
		routerV1.POST("registration", v1.AddUser)
		// 紧急通知
		routerV1.GET("notice", v1.GetNotice)
		//文章
		routerV1.GET("article", v1.GetArticle)
		// 部门
		routerV1.GET("sector", v1.GetSector)

	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}

}
