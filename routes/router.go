package routes

import (
	"github.com/gin-gonic/gin"
	v1 "star-server/api/v1"
	"star-server/middleware"
	"star-server/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	authV1 := r.Group("api/v1")
	authV1.Use(middleware.JwtToken())
	{
		//用户模块
		authV1.PUT("user/update/:id", v1.EditUser)
		authV1.GET("user/:id", v1.GetUser)
		authV1.POST("notice/add", v1.AddNotice)
		authV1.POST("paper/add", v1.CreatePaper)
		authV1.POST("sector/add", v1.CreateSector)
		authV1.POST("student/add", v1.CreateStudent)
		authV1.GET("schedule", v1.GetSchedule)
		authV1.POST("schedule/add", v1.AddOneRecord)
	}
	routerV1 := r.Group("api/v1")
	{
		routerV1.POST("registration", v1.AddUser)
		routerV1.GET("notice", v1.GetNotice)
		routerV1.GET("paper", v1.GetPaper)
		routerV1.GET("sector", v1.GetSector)
		routerV1.GET("student", v1.GetStudent)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}

}
