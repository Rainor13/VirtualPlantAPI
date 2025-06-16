package models

type GrowthStage string

const (
	Seed    GrowthStage = "seed"
	Seeding GrowthStage = "seeding"
	Young   GrowthStage = "young"
	Mature  GrowthStage = "mature"
)

// Mapa de transición ordenada de etapas
var nextStage = map[GrowthStage]GrowthStage{
	Seed:    Seeding,
	Seeding: Young,
	Young:   Mature,
	Mature:  Mature, // ya no avanza más
}

// Metodo que devuelve la siguiente etapa
func (g GrowthStage) Next() GrowthStage {
	return nextStage[g]
}
