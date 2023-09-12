package routes

import (
	"github.com/starlingvibes/HNGX-BE-Task_2/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/api", controllers.CreateUser())
	router.GET("/api/:userId", controllers.GetAUser())
	router.PUT("/api/:userId", controllers.EditAUser())
	router.DELETE("/api/:userId", controllers.DeleteAUser())
	router.GET("/api", controllers.GetAllUsers())
}
