package models

type GrowthStage string

const (
	Seed    GrowthStage = "seed"
	Seeding GrowthStage = "seeding"
	Young   GrowthStage = "young"
	Mature  GrowthStage = "mature"
)
