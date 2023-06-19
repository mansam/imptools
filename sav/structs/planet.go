package structs

import (
	"encoding/binary"
	"fmt"
	"github.com/mansam/imptools/sav/labels"
	"os"
)

type Planet struct {
	NameLength  uint8
	Name_       [12]byte
	X           int16
	Y           int16
	Map         uint8
	Fighter1    uint16
	Fighter2    uint16
	Fighter3    uint16
	Fighter4    uint16
	Fighter5    uint16
	Fighter6    uint16
	Tank1       uint16
	Tank2       uint16
	Tank3       uint16
	Tank4       uint16
	Unused      uint32 // likely unused vehicle slots
	Car1        uint16
	Car2        uint16
	Car3        uint16
	Race        uint8
	Owner       uint8
	Type        uint8
	Flags       uint8 // Bit flags. Humans start at 208. Adding a spy sat 2 to a planet changes it to 208.
	Visibility  uint8
	Population  uint32
	RankListed  uint8
	RankVisible uint8
	BuildIndex  uint8
	Morale      uint16
	HasSat      bool
	HasSpySat   bool
	HasSpySat2  bool
	HasHubble2  bool
	Orbit1      uint8
	Orbit2      uint8
	Orbit3      uint8
	TaxLevel    uint8
	Virus       bool
}

func (r Planet) Name() string {
	return string(r.Name_[:r.NameLength])
}

func (r Planet) String() string {
	return fmt.Sprintf("%-12s \t%-24s \t%-10s \t%-6v \t%08b \t%-4v \t%-4d \t%-4v", r.Name(), labels.Owner(r.Owner), labels.Race(r.Race), r.Morale, r.Flags, r.Visibility, r.Type, r.Map)
}

func WritePlanet(p Planet, f *os.File) (err error) {
	err = binary.Write(f, binary.LittleEndian, p)
	return
}

func ReadPlanet(f *os.File) (p Planet) {
	_ = binary.Read(f, binary.LittleEndian, &p)
	return
}
