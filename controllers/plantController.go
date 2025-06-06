package controllers

import (
	"github.com/Rainor13/VirtualPlantAPI/models"
	"github.com/Rainor13/VirtualPlantAPI/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PlantController struct {
	Service *services.PlantService
}

// Constructor
func NewPlantController(service *services.PlantService) *PlantController {
	return &PlantController{Service: service}
}

func (pc *PlantController) CreatePlant() {
	var newPlant models.Plant

	//err contiene el resultado de asignar los valores del post a newPlant si no es
	//posible entra en el if y nos muestra el error, es decir err contiene el erro en concreto
	if err := c.ShouldBindJSON(&newPlant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Se ha ejecutado correctamente")

	createdPlant := pc.Service.CreatePlant(newPlant)

	c.JSON(http.StatusCreated, createdPlant)
}

func GetPlants(c *gin.Context) {
	// aquí devuelves todas las plantas
	c.JSON(http.StatusOK, gin.H{"plants": []models.Plant{}})
}

// Crea las otras funciones (CreatePlant, GetPlantByID, etc.)
