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
	//使用自己的log
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	authV1 := r.Group("api/v1")
	authV1.Use(middleware.JwtToken())
	{
		//用户模块
		authV1.PUT("user/:uid/update", v1.UpdateUser)
		authV1.GET("user/:uid", v1.GetUser)

		//部门
		authV1.POST("sector/:uid/join", v1.CreateStuSect)
		authV1.GET("sector/:uid/find", v1.FindStuSector)
		//学生
		authV1.GET("student/:student_id", v1.GetStudent)
		//值班表
		authV1.GET("schedule/:sector_name", v1.GetSchedule)

		// 工作记录
		authV1.POST("workform/add", v1.CreateForm)
		authV1.PUT("workform/update/:uid", v1.UpdateForm)
		authV1.GET("workform/get/:student_id", v1.GetStuForm)
		// todo 获得用户的工作情况

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

		routerV1.GET("test", v1.GetTest)
		routerV1.POST("test", v1.CreateTest)

	}
	adminV1 := r.Group("api/v1")
	adminV1.Use(middleware.AdminToken())
	{
		// 紧急通知
		adminV1.POST("notice/add", v1.CreateNotice)
		// 文章
		adminV1.POST("article/add", v1.CreateArticle)
		//部门
		adminV1.POST("sector/add", v1.CreateSector)
		// 学生
		adminV1.POST("student/add", v1.CreateStudent)
		// 一条值班记录
		adminV1.POST("schedule/add", v1.AddOneRecord)
		//部门邀请码
		adminV1.POST("sectorkey/add", v1.CreateSectorKey)
		// 更新用户权限
		adminV1.PUT("user/:uid/updateauth", v1.UpdateUserAuth)

	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}

}
