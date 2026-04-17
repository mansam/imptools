package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mansam/imptools/sav/labels"
	"github.com/mansam/imptools/sav/structs"
	"github.com/mansam/imptools/sav/tables"
)

var (
	MainExe              string
	OutputBuildings      bool
	OutputPlanets        bool
	OutputUnknownStructs bool
)

const (
	BuildingDefinitionOffset = 0x0AEB14
	NumBuildingDefinitions   = 43
	PlanetOffset             = 0xA9D44
	NumPlanets               = 105
)

func main() {
	if len(os.Args) != 3 {
		panic(fmt.Sprintf("usage: %s main_exe flags", os.Args[0]))
	}
	MainExe = os.Args[1]
	OutputBuildings = strings.Contains(os.Args[2], "b")
	OutputPlanets = strings.Contains(os.Args[2], "p")
	OutputUnknownStructs = strings.Contains(os.Args[2], "u")

	f, err := os.Open(MainExe)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Seek(BuildingDefinitionOffset, 0)
	if err != nil {
		panic(err)
	}
	var buildingDefinitions []any
	for i := 0; i < NumBuildingDefinitions; i++ {
		buildingDefinitions = append(buildingDefinitions, structs.ReadBuildingDefinition(f))
	}

	_, err = f.Seek(PlanetOffset, 0)
	if err != nil {
		panic(err)
	}
	var planets []any
	for i := 0; i < NumPlanets; i++ {
		planets = append(planets, structs.ReadPlanet(f))
	}

	_, err = f.Seek(0xa6a9b, 0)
	if err != nil {
		panic(err)
	}
	var unknown22 []any
	for i := 0; i < 228; i++ {
		unknown22 = append(unknown22, structs.ReadUnknownStructEntry22(f))
	}
	var unknown14 []any
	for i := 0; i < 11; i++ {
		unknown14 = append(unknown14, structs.ReadUnknownStructEntry14(f))
	}

	_, err = f.Seek(structs.UnknownStruct156Offset, 0)
	if err != nil {
		panic(err)
	}
	var unknown156 []any
	for i := 0; i < 14; i++ {
		unknown156 = append(unknown156, structs.ReadUnknownStruct156(f))
	}

	if OutputBuildings {
		tables.Tableize(buildingDefinitions, labels.BuildingName)
	}
	if OutputPlanets {
		tables.Tableize(planets)
	}
	if OutputUnknownStructs {
		tables.Tableize(unknown22)
		tables.Tableize(unknown14)
		tables.Tableize(unknown156)
	}
}
