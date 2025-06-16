package routes

import (
	"github.com/Rainor13/VirtualPlantAPI/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	plantRoutes := router.Group("/api")
	{
		plantRoutes.POST("/plant", controllers.CreatePlant)
		plantRoutes.POST("/plant/:id/water", controllers.WaterPlant)
		plantRoutes.GET("/plant/:id", controllers.CheckPlantStatus)
		plantRoutes.POST("/api/plant/:id/harvest", controllers.HarvestPlant)
		plantRoutes.GET("/api/plant/", controllers.GetPlants)
		plantRoutes.POST("/api/plant/:id/fertilize", controllers.GetPlants)
