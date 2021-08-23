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

	}
	routerV1 := r.Group("api/v1")
	{
		routerV1.POST("registration", v1.AddUser)
		routerV1.GET("user/:id", v1.GetUser)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}

}
