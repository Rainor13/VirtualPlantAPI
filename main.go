package main

import (
	"fmt"
	"github.com/Rainor13/VirtualPlantAPI/controllers"
	"github.com/Rainor13/VirtualPlantAPI/routes"
	"github.com/Rainor13/VirtualPlantAPI/services"
	"github.com/gin-gonic/gin"
)

func main() {

	s := "gopher"
	fmt.Println("Esto es el main de las plantas, %s!", s)

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")

	fmt.Println("== Prueba manual de VirtualPlantAPI ==")

	// Instanciar service
	plantService := services.NewPlantService()

	// Instanciar controller pasándole el service
	plantController := controllers.NewPlantController(plantService)

	// Usar método "manual" de prueba
	plantController.CreatePlant()
}
