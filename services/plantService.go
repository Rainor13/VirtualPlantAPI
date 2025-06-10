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
	p.Health = 100

	plants = append(plants, p)
	return p
}

//The GET /api/plant/{id} endpoint should return the current status of the plant, including its health and growth stage.

// for _, plant := range plants... Cuando usas este bucle estas accediendo a la copia de los elementos de la lista por tanto si haces &plant
// estas devolviendo la copia de plant que se elimina cuando acaba la funcion, elemento temporal del slice.
// Por otro lado si lo que hacemos es usar el indice con i &plants[i] devuelves elelemento real no de una copia.

//models.Plant{}, false para cuando el return esta dentro de un if y si llega a fuera del if como tiene que devolver un objeto se puede poner como el tipo, que entiendo que es como un objeto vacio

func FindPlant(id string) (*models.Plant, bool) {
	for i := range plants {
		if plants[i].ID == id {
			return &plants[i], true
		}
	}
	return nil, false
}

func WaterPlant(id string) {
	var plant *models.Plant
	var found bool
	plant, found = FindPlant(id)

	if found {
		plant.LastWateredDate = time.Now()
		if plant.Health < 100 && plant.Health >= 90 {
			plant.Health = 100
		} else if plant.Health < 90 {
			plant.Health += 10
		}
	}
}

func CheckPlantStatus(id string) (models.GrowthStage, int) {
	plant, found := FindPlant(id)
	if !found {
		return "", 0
	}
	return plant.PlantStage, plant.Health
}
