package main

import (
	"fmt"
	"github.com/mansam/imptools/reader"
	"github.com/mansam/imptools/sav/labels"
	"github.com/mansam/imptools/sav/structs"
	"os"
	"sort"
	"strings"
)

func ReadHeader(f *os.File) {
	f.Seek(structs.HeaderOffset, 0)

}

// Read IG save files.
func main() {
	if len(os.Args) != 3 {
		panic(fmt.Sprintf("usage: %s savefile flags", os.Args[0]))
	}
	filepath := os.Args[1]
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//var outputHeader bool
	var outputPlanets bool
	var outputBuildings bool
	var outputTech bool
	//var outputFleets bool
	outputPlanets = strings.Contains(os.Args[2], "p")
	outputBuildings = strings.Contains(os.Args[2], "b")
	outputTech = strings.Contains(os.Args[2], "t")
	//outputFleets = strings.Contains(os.Args[2], "f")

	_, err = f.Seek(structs.PlanetDefOffset, 0)
	if err != nil {
		panic(err)
	}

	if outputPlanets {
		planets := []structs.Planet{}
		for i := 0; i < structs.NumPlanets; i++ {
			planets = append(planets, structs.ReadPlanet(f))
		}
		sort.Slice(planets, func(i, j int) bool {
			return planets[i].Name < planets[j].Name
		})

		fmt.Printf("%-12s \t%-25s \t%-10s \t%-6s \t%-3s \t%-4v \t%-4s \t%-4s\n", "Name", "Owner", "Race", "Morale", "Flag", "Vis", "Type", "Map")
		for _, v := range planets {
			fmt.Println(v.String())
		}
	}

	if outputBuildings {
		f.Seek(structs.BuildingNumberOffset, 0)
		numberOfBuildings := reader.Btoi(reader.ReadN(f, 2))
		fmt.Printf("\nNumber of Buildings: %d\n", numberOfBuildings)
		buildings := []structs.Building{}
		for i := 0; i < numberOfBuildings; i++ {
			buildings = append(buildings, structs.ReadBuilding(f))
			f.Seek(1, 1)
		}
		fmt.Printf("%-12s %-25s %-20s %-9s %-9s %-9s %-9s %-9s %-9s %-9s %-9s %-9s\n",
			"Planet",
			"Owner",
			"Building",
			"Powered",
			"Repairing",
			"Damage",
			"%Complete",
			"Efficiency",
			"Flag1",
			"Short1",
			"KW/h",
			"Workers")
		for _, v := range buildings {
			fmt.Println(v.String())
		}
	}

	if outputTech {
		f.Seek(structs.TechnologyOffset, 0)
		technologies := []structs.Technology{}
		for i := 0; i < structs.NumTechnologies; i++ {
			technologies = append(technologies, structs.ReadTechnology(f))
		}
		fmt.Println("Technologies:")
		fmt.Printf(
			"%-18s \t%-7s \t%-7s\n",
			"Name",
			"Unknown1",
			"Unknown2",
		)
		for i, v := range technologies {
			fmt.Printf("%-18s \t%s\n", labels.TechnologyName(i+1), v.String())
		}
	}
}
