package services

import (
	"fmt"
	"github.com/Rainor13/VirtualPlantAPI/models"
	"github.com/google/uuid"
	"time"
)

type PlantService struct {
	plants []models.Plant
}

// Constructor
func NewPlantService() *PlantService {
	return &PlantService{plants: []models.Plant{}}
}

// El receptor de esta funcion es PlantService que basicamente modifica el contenido original de ese struct
// despues tenemos la variable p por valor del tip Plant
// El tipo de devuelve la funcion es models.Plant
func (ps *PlantService) CreatePlant(p models.Plant) models.Plant {
	p.ID = uuid.New().String() // Asignamos un ID único
	p.PlantingDate = time.Now()
	p.LastWateredDate = time.Now()
	p.PlantStage = models.Seed

	ps.plants = append(ps.plants, p)

	fmt.Println("ID:", p.ID)
	fmt.Println("Species:", p.Species)
	fmt.Println("Planting Date:", p.PlantingDate)
	fmt.Println("Last Watered Date:", p.LastWateredDate)
	fmt.Println("Health:", p.Health)
	fmt.Println("Plant Stage:", p.PlantStage)

	return p
}
