package main

import (
	"fmt"

	"github.com/gg-restapi-my/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")

	router := SetupRouter()
	router.Run(":8077")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/flower", controllers.Create)
		v1.GET("/flower/:id", controllers.GetFlower)
		v1.GET("/flowers", controllers.GetAllFlower)
		v1.DELETE("/flower", controllers.DeleteFlower)
		v1.GET("/check", controllers.HealthCheck)
	}
	return router
}
