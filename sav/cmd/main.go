package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/mansam/imptools/reader"
	"github.com/mansam/imptools/sav/labels"
	"github.com/mansam/imptools/sav/structs"
)

// Flags
var (
	SaveFile        string
	OutputHeader    bool
	OutputPlanets   bool
	OutputBuildings bool
	OutputTech      bool
	OutputFleets    bool
	OutputShips     bool
	OutputMessages  bool
)

// Read IG save files.
func main() {
	if len(os.Args) != 3 {
		panic(fmt.Sprintf("usage: %s savefile flags", os.Args[0]))
	}
	SaveFile = os.Args[1]
	OutputHeader = strings.Contains(os.Args[2], "h")
	OutputPlanets = strings.Contains(os.Args[2], "p")
	OutputBuildings = strings.Contains(os.Args[2], "b")
	OutputTech = strings.Contains(os.Args[2], "t")
	OutputFleets = strings.Contains(os.Args[2], "f")
	OutputShips = strings.Contains(os.Args[2], "s")
	OutputMessages = strings.Contains(os.Args[2], "m")

	f, err := os.Open(SaveFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if OutputHeader {
		fmt.Println(structs.ReadHeader(f))
	}

	if OutputFleets {
		numFleets := reader.Btoi(reader.ReadNAt(f, 2, structs.NumFleetsOffset))

		_, err = f.Seek(structs.FleetsOffset, 0)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%-3s %-12s %-16s %-4s %-4s %-4s %-4s %-4s %-4s %-4s %-4s %-4s %-4s %-4s %-4s %-4s %-8s %-8s %-8s\n",
			"#",
			"Filename",
			"Coords",
			"F1",
			"F2",
			"F3",
			"F4",
			"F5",
			"F6",
			"T1",
			"T2",
			"T3",
			"T4",
			"C1",
			"C2",
			"C3",
			"Fighters",
			"Vehicles",
			"Unknown")
		for i := 0; i < numFleets; i++ {
			fleet := structs.ReadFleet(f)
			fmt.Printf("%-3d %s\n", i+1, fleet)
		}
	}

	if OutputShips {
		numShips := reader.Btoi(reader.ReadNAt(f, 2, structs.NumShipsOffset))

		_, err = f.Seek(structs.ShipsOffset, 0)
		if err != nil {
			panic(err)
		}
		for i := 0; i < numShips; i++ {
			ship := structs.ReadShip(f)
			fmt.Printf("%-3d %s\n", i+1, ship)
		}
	}

	_, err = f.Seek(structs.PlanetDefOffset, 0)
	if err != nil {
		panic(err)
	}

	if OutputPlanets {
		planets := []structs.Planet{}
		for i := 0; i < structs.NumPlanets; i++ {
			planets = append(planets, structs.ReadPlanet(f))
		}
		sort.Slice(planets, func(i, j int) bool {
			return planets[i].Name() < planets[j].Name()
		})

		fmt.Printf("%-12s \t%-24s \t%-10s \t%-16s \t%-6s \t%-8s \t%-4v \t%-4s \t%-4s\n", "Filename", "Owner", "Race", "Coords", "Morale", "Flag", "Vis", "Type", "Map")
		for _, v := range planets {
			fmt.Println(v)
		}
	}

	if OutputBuildings {
		f.Seek(structs.BuildingNumberOffset, 0)
		numberOfBuildings := reader.Btoi(reader.ReadN(f, 2))
		fmt.Printf("\nNumber of Buildings: %d\n", numberOfBuildings)
		buildings := []structs.Building{}
		for i := 0; i < numberOfBuildings; i++ {
			buildings = append(buildings, structs.ReadBuilding(f))
			// one byte gap between buildings
			f.Seek(1, 1)
		}
		fmt.Printf("%-12s %-25s %-20s %-10s %-9s %-9s %-9s %-9s %-10s %-12s %-9s %-9s %-9s\n",
			"Planet",
			"Owner",
			"Building",
			"Coords",
			"Powered",
			"Repairing",
			"Damage",
			"Remaining",
			"Efficiency",
			"Operational",
			"Unused",
			"KW/h",
			"Workers")
		for _, v := range buildings {
			fmt.Println(v.String())
		}
	}

	if OutputTech {
		f.Seek(structs.TechnologyOffset, 0)
		technologies := []structs.Technology{}
		for i := 0; i < structs.NumTechnologies; i++ {
			technologies = append(technologies, structs.ReadTechnology(f))
		}
		fmt.Println("Technologies:")
		fmt.Printf(
			"%-18s \t%-7s \t%-7s\n",
			"Filename",
			"Unknown1",
			"Unknown4",
		)
		for i, v := range technologies {
			fmt.Printf("%-18s \t%s\n", labels.TechnologyName(uint8(i+1)), v.String())
		}
	}

	if OutputMessages {
		f.Seek(structs.SpecialMessageOffset, 0)
		special := structs.ReadSpecialMessage(f)
		fmt.Printf("%d: %s\n", special.Length, special)
		f.Seek(structs.ScrollingMessageOffset, 0)
		for i := 0; i < structs.NumMessages; i++ {
			message := structs.ReadScrollingMessage(f)
			fmt.Printf("%d: %s$\n", message.Length, message)
		}
	}
}
