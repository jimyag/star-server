package main

import (
	"star-server/model"
	"star-server/routes"
)

func main() {
	model.InitDb()

	routes.InitRouter()
}
