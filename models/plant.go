package models

import (
	"time"
)

type Plant struct {
	ID              string      `json:"id"`
	Species         string      `json:"species"`
	PlantingDate    time.Time   `json:"planting_date"`
	LastWateredDate time.Time   `json:"last_watered_Date"`
	Health          int         `json:"health"`
	PlantStage      GrowthStage `json:"growth"`
}
