package controllers

import (
	"github.com/Rainor13/VirtualPlantAPI/models"
	"github.com/Rainor13/VirtualPlantAPI/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreatePlant(c *gin.Context) {
	var newPlant models.Plant

	//err contiene el resultado de asignar los valores del post a newPlant si no es
	//posible entra en el if y nos muestra el error, es decir err contiene el erro en concreto
	if err := c.ShouldBindJSON(&newPlant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPlant := services.CreatePlant(newPlant)

	log.Println("Se ha ejecutado correctamente")

	c.JSON(http.StatusCreated, createdPlant)
}

func WaterPlant(c *gin.Context) {
	var updatedPlant models.Plant
	var id string

	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPlant = services.WaterPlant(id)

	log.Println("Se ha ejecutado correctamente")
	log.Println("Plant watered successfully")

	c.JSON(http.StatusOK, updatedPlant)

}

func CheckPlantStatus(c *gin.Context) {
	var id string
	var plant models.Plant
	var found bool

	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plant, found = services.CheckPlantStatus(id)
	log.Println("Se ha ejecutado correctamente")

	if !found {
		log.Println("No se ha encontrado la planta")
	} else {
		log.Println("Se ha encontrado la planta y se ha actualizado: Estado ", plant.PlantStage, " Salud ", plant.Health)
	}
}

func HarvestPlant(c *gin.Context) {
	var id string

	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success := services.HarvestPlant(id)

	if success {
		log.Println("Se ha ejecutado correctamente")
		log.Println("Plant harvested successfully")
		c.JSON(http.StatusOK, success)
	} else {
		log.Println("No se encontró la planta")
		c.JSON(http.StatusNotFound, success)
	}
}

func GetPlants(c *gin.Context) {
	// aquí devuelves todas las plantas
	c.JSON(http.StatusOK, gin.H{"plants": []models.Plant{}})
}

// Crea las otras funciones (CreatePlant, GetPlantByID, etc.)
