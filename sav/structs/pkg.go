package structs

const HeaderOffset = 0x0
const PlanetDefOffset = 0x21f0e
const BuildingOffset = 0x122b6
const BuildingNumberOffset = 0x122b4
const PlanetDefLength = 71
const PlanetNameLen = 12
const NumPlanets = 105
const BuildingLength = 18
const TechnologyLength = 42
const TechnologyOffset = 0x10f04
const NumTechnologies = 120
const ShipsOffset = 0x2261
const FleetsOffset = 0xea53

// const NumShipsOffset = 0x0f58
const NumShipsOffset = 0x2259
const NumFleetsOffset = 0x225b

// technology name indices
const (
	TechFighters         = 0
	TechDestroyers       = 6
	TechFlagships        = 12
	TechSatellites       = 18
	TechSpaceBases       = 24
	TechHyperdrives      = 30
	TechModules          = 36
	TechRadars           = 42
	TechShields          = 48
	TechLasers           = 60
	TechGuns             = 66
	TechMissiles         = 72
	TechTanks            = 78
	TechCars             = 84
	TechBuildings        = 90
	TechFortresses       = 96
	TechPlanetaryShields = 102
	TechPlanetaryRadars  = 106
	TechPlanetaryGuns    = 114
)
