package main

import (
	"github.com/starlingvibes/HNGX-BE-Task_2/configs"
	"github.com/starlingvibes/HNGX-BE-Task_2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)

	router.Run("localhost:6000")
}
