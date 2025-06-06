package main

import (
	"fmt"
	"github.com/Rainor13/VirtualPlantAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	s := "gopher"
	fmt.Println("Esto es el main de las plantas, %s!", s)

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")

}
