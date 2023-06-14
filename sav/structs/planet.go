package structs

import (
	"fmt"
	"github.com/mansam/imptools/reader"
	"github.com/mansam/imptools/sav/labels"
	"os"
)

type Planet struct {
	Name        string
	Owner       int
	X           int
	Y           int
	Map         int
	Fighter1    int
	Fighter2    int
	Fighter3    int
	Fighter4    int
	Fighter5    int
	Fighter6    int
	Tank1       int
	Tank2       int
	Tank3       int
	Tank4       int
	Unused      int // likely unused vehicle slots
	Car1        int
	Car2        int
	Car3        int
	Race        int
	Type        int
	Unknown     int // still no clue what this is. After turning on cheats all planets show 208
	Visibility  int
	Population  int
	RankListed  int
	RankVisible int
	BuildIndex  int
	Morale      int
	HasSat      bool
	HasSpySat   bool
	HasSpySat2  bool
	HasHubble2  bool
	Orbit1      int
	Orbit2      int
	Orbit3      int
	TaxLevel    int
	Virus       bool
}

func (r *Planet) String() string {
	return fmt.Sprintf("%-12s \t%-25s \t%-10s \t%-6v \t%-3v \t%-4v \t%-4d \t%-4v", r.Name, labels.Owner(r.Owner), labels.Race(r.Race), r.Morale, r.Unknown, r.Car1, r.Type, r.Map)
}

func ReadPlanet(f *os.File) (p Planet) {
	nameLen := reader.Btoi(reader.ReadN(f, 1))
	rawName := reader.ReadN(f, PlanetNameLen)
	p.Name = string(rawName[:nameLen])
	p.X = reader.Btoi(reader.ReadN(f, 2))
	p.Y = reader.Btoi(reader.ReadN(f, 2))
	p.Map = reader.Btoi(reader.ReadN(f, 1))
	p.Fighter1 = reader.Btoi(reader.ReadN(f, 2))
	p.Fighter2 = reader.Btoi(reader.ReadN(f, 2))
	p.Fighter3 = reader.Btoi(reader.ReadN(f, 2))
	p.Fighter4 = reader.Btoi(reader.ReadN(f, 2))
	p.Fighter5 = reader.Btoi(reader.ReadN(f, 2))
	p.Fighter6 = reader.Btoi(reader.ReadN(f, 2))
	p.Tank1 = reader.Btoi(reader.ReadN(f, 2))
	p.Tank2 = reader.Btoi(reader.ReadN(f, 2))
	p.Tank3 = reader.Btoi(reader.ReadN(f, 2))
	p.Tank4 = reader.Btoi(reader.ReadN(f, 2))
	p.Unused = reader.Btoi(reader.ReadN(f, 4))
	p.Car1 = reader.Btoi(reader.ReadN(f, 2))
	p.Car2 = reader.Btoi(reader.ReadN(f, 2))
	p.Car3 = reader.Btoi(reader.ReadN(f, 2))
	p.Race = reader.Btoi(reader.ReadN(f, 1))
	p.Owner = reader.Btoi(reader.ReadN(f, 1))
	p.Type = reader.Btoi(reader.ReadN(f, 1))
	p.Unknown = reader.Btoi(reader.ReadN(f, 1))
	p.Visibility = reader.Btoi(reader.ReadN(f, 1))
	p.Population = reader.Btoi(reader.ReadN(f, 4))
	p.RankListed = reader.Btoi(reader.ReadN(f, 1))
	p.RankVisible = reader.Btoi(reader.ReadN(f, 1))
	p.BuildIndex = reader.Btoi(reader.ReadN(f, 1))
	p.Morale = reader.Btoi(reader.ReadN(f, 2))
	p.HasSat = reader.Bool(reader.ReadN(f, 1)[0])
	p.HasSpySat = reader.Bool(reader.ReadN(f, 1)[0])
	p.HasSpySat2 = reader.Bool(reader.ReadN(f, 1)[0])
	p.HasHubble2 = reader.Bool(reader.ReadN(f, 1)[0])
	p.Orbit1 = reader.Btoi(reader.ReadN(f, 1))
	p.Orbit2 = reader.Btoi(reader.ReadN(f, 1))
	p.Orbit3 = reader.Btoi(reader.ReadN(f, 1))
	p.TaxLevel = reader.Btoi(reader.ReadN(f, 1))
	p.Virus = reader.Bool(reader.ReadN(f, 1)[0])
	return
}
