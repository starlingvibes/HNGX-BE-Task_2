package main

import (
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":"+port)
}
