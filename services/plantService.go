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

func FindPositionPlant(id string) (models.Plant, int, bool) {
	for i, plant := range plants {
		if plant.ID == id {
			return plant, i, true
		}
	}
	return models.Plant{}, 0, false
}

func WaterPlant(id string) models.Plant {
	var plant *models.Plant
	var found bool
	plant, found = FindPlant(id)

	if !found {
		return models.Plant{}
	} else {
		plant.LastWateredDate = time.Now()
		if plant.Health < 100 && plant.Health >= 90 {
			plant.Health = 100
		} else if plant.Health < 90 {
			plant.Health += 10
		}
		return *plant
	}
}

func CheckPlantStatus(id string) (models.Plant, bool) {
	plant, found := FindPlant(id)
	if !found {
		return models.Plant{}, found
	}
	return *plant, found
}

func HarvestPlant(id string) bool {
	plant, pos, found := FindPositionPlant(id)

	if found && plant.PlantStage == models.Mature {
		if pos == 0 {
			plants = plants[1:]
		} else if pos > 0 && pos < len(plants)-1 {
			plants = append(plants[:pos], plants[pos+1:]...)
		} else if pos == len(plants)-1 {
			plants = plants[:len(plants)-1]
		}
		return true
	}

	return false
}

func UpdateWaterPlantsStatus() {
	for i := range plants {
		minuteSinceWatered := time.Since(plants[i].LastWateredDate).Minutes()
		if minuteSinceWatered > 1 { // por ejemplo, más de 2 días sin regar
			plants[i].Health -= 10
			if plants[i].Health < 0 {
				plants[i].Health = 0
			}
		}
	}
}

//cada 4 minutos y si estan al 70% de regado entonces pasan a el siguiente estado

func UpdateStagePlantsStatus() {
	for i := range plants {
		minuteSincePlanted := time.Since(plants[i].PlantingDate).Minutes()
		if minuteSincePlanted > 4 && plants[i].Health >= 70 {
			plants[i].PlantStage.Next()
		}
	}
}

func StartBackgroundWorker() {
	ticker := time.NewTicker(1 * time.Minute) // cada minuto (ajusta el tiempo a lo que necesites)
	go func() {
		for range ticker.C {
			// Aquí llamas a tu función que actualiza la salud de todas las plantas
			UpdateWaterPlantsStatus()
			UpdateStagePlantsStatus()
		}
	}()
}
