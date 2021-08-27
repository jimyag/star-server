package main

import (
	"star-server/model"
	"star-server/routes"
)

func main() {
	model.InitDb()

	routes.InitRouter()
	//res,errmsg := utils.GetOpenid("003QFEGa1uocFB05GYIa1XzjVk1QFEGX")
	//fmt.Println(res)
	//fmt.Println(errmsg)
}
