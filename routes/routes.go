package routes

import (
	"github.com/Rainor13/VirtualPlantAPI/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	plantRoutes := router.Group("/api")
	{
		plantRoutes.POST("/plant", controllers.CreatePlant)
		plantRoutes.GET("/", controllers.GetPlants)
		plantRoutes.POST("/plant/{id}/water", controllers.GetPlants)
		plantRoutes.GET("/plant/{id}", controllers.GetPlants)
		//plantRoutes.GET("/:id", controllers.GetPlantByID)
		//plantRoutes.PUT("/:id", controllers.UpdatePlant)
		//plantRoutes.DELETE("/:id", controllers.DeletePlant)
	}
}
