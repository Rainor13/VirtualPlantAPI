package services

import (
	"github.com/Rainor13/VirtualPlantAPI/models"
	"github.com/google/uuid"
	"time"
)

var plants = []models.Plant{}

func CreatePlant(p models.Plant) models.Plant {
	p.ID = uuid.New().String() // Asignamos un ID único
	p.PlantingDate = time.Now()
	p.LastWateredDate = time.Now()
	p.PlantStage = models.Seed

	plants = append(plants, p)
	return p
}
