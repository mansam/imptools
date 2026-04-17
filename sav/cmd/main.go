package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mansam/imptools/reader"
	"github.com/mansam/imptools/sav/labels"
	"github.com/mansam/imptools/sav/structs"
	"github.com/mansam/imptools/sav/tables"
	"gopkg.in/yaml.v3"
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
	OutputRaw       bool
	OutputYaml      bool
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
	OutputRaw = strings.Contains(os.Args[2], "r")
	OutputYaml = strings.Contains(os.Args[2], "y")

	f, err := os.Open(SaveFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Seek(structs.FleetsOffset, 0)
	if err != nil {
		panic(err)
	}

	numFleets := reader.Btoi(reader.ReadNAt(f, 2, structs.NumFleetsOffset))
	var fleets []any
	for i := 0; i < numFleets; i++ {
		fleets = append(fleets, structs.ReadFleet(f))
	}
	numShips := reader.Btoi(reader.ReadNAt(f, 2, structs.NumShipsOffset))
	var ships []any
	_, err = f.Seek(structs.ShipsOffset, 0)
	if err != nil {
		panic(err)
	}
	for i := 0; i < numShips; i++ {
		ships = append(ships, structs.ReadShip(f))
	}
	_, err = f.Seek(structs.PlanetDefOffset, 0)
	if err != nil {
		panic(err)
	}
	var planets []any
	for i := 0; i < structs.NumPlanets; i++ {
		planets = append(planets, structs.ReadPlanet(f))
	}
	//sort.Slice(planets, func(i, j int) bool {
	//	return planets[i].(structs.Planet).BuildIndex < planets[j].(structs.Planet).BuildIndex
	//})
	f.Seek(structs.BuildingNumberOffset, 0)
	numberOfBuildings := reader.Btoi(reader.ReadN(f, 2))
	var buildings []any
	for i := 0; i < numberOfBuildings; i++ {
		buildings = append(buildings, structs.ReadBuilding(f))
		//one byte gap between buildings
		f.Seek(1, 1)
	}
	f.Seek(structs.TechnologyOffset, 0)
	var technologies []any
	for i := 0; i < structs.NumTechnologies; i++ {
		technologies = append(technologies, structs.ReadTechnology(f))
	}

	if OutputHeader {
		fmt.Println(structs.ReadHeader(f))
	}

	if OutputFleets {
		if OutputRaw {
			tables.Tableize(fleets)
		} else if OutputYaml {
			var out []byte
			out, err = yaml.Marshal(fleets)
			fmt.Printf("%s\n", out)
		} else {
			t := table.NewWriter()
			t.SetAutoIndex(true)
			t.SetStyle(table.StyleBold)
			h := []any{"Name", "Owner", "X", "Y", "Visible", "Orders"}
			t.AppendHeader(h)
			for _, raw := range fleets {
				fleet := raw.(structs.Fleet)
				r := []any{fleet.Name(), labels.Owner(fleet.Owner), fleet.X, fleet.Y, fleet.Visible, fleet.Orders(planets, fleets)}
				t.AppendRow(r)
			}
			fmt.Println(t.Render())
		}
	}
	if OutputShips {
		if OutputRaw {
			tables.Tableize(ships)
		} else {
			for i, s := range ships {
				fmt.Printf("%-3d %s\n", i+1, s)
			}
		}
	}

	if OutputPlanets {
		//sort.Slice(planets, func(i, j int) bool {
		//	return planets[i].Name() < planets[j].Name()
		//})
		if OutputRaw {
			tables.Tableize(planets)
		} else {
			t := table.NewWriter()
			t.SetAutoIndex(true)
			t.SetStyle(table.StyleBold)
			h := []any{"Name", "Owner", "Race", "Coords", "Morale", "Flag", "Vis", "Type", "Map", "Tax"}
			t.AppendHeader(h)
			for _, p := range planets {
				t.AppendRow(p.(structs.Planet).Fields())
			}
			fmt.Println(t.Render())
		}
	}

	if OutputBuildings {
		if OutputRaw {
			tables.Tableize(buildings)
		} else {
			t := table.NewWriter()
			t.SetAutoIndex(true)
			t.SetStyle(table.StyleBold)
			h := []any{
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
				"Workers",
			}
			t.AppendHeader(h)
			for _, b := range buildings {
				t.AppendRow(b.(structs.Building).Fields())
			}
			fmt.Println(t.Render())
		}
	}

	if OutputTech {
		if OutputRaw {
			tables.Tableize(technologies)
		} else {
			fmt.Println("Technologies:")
			fmt.Printf(
				"%-18s \t%-7s \t%-7s\n",
				"Filename",
				"Unknown1",
				"Unknown4",
			)
			for i, v := range technologies {
				fmt.Printf("%-18s \t%s\n", labels.TechnologyName(uint8(i+1)), v.(structs.Technology).String())
			}
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
