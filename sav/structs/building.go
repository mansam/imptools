package structs

import (
	"encoding/binary"
	"fmt"
	"github.com/mansam/imptools/sav/labels"
	"os"
)

// Building - 18 bytes
// Instance of a building on a planet surface.
// Base stats like total HP and required power/workers
// found in executable's building definitions.
// Unused sometimes contains apparent garbage.
// Saving -> Reloading -> Saving again causes it to be zeroed out again
// except for a few planets that seem to be generated with some garbage.
type Building struct {
	Type        uint8
	Owner       uint8
	Powered     bool
	Repairing   bool
	Damage      uint16
	Remaining   uint8
	Index       uint8
	X           int8
	Y           int8
	Efficiency  uint8
	Operational bool
	Unused      uint16
	Kwh         uint16
	Workers     uint16
}

// this is not accurate.
// build time must be part of building definitions in executable.
//func (r Building) Complete() uint8 {
//	return (r.Remaining / 255) * 100
//}

func (r Building) String() string {
	return fmt.Sprintf("%-12s %-25s %-20s %-10v %-9v %-9v %-9d %-9d %-10d %-12v %-9d %-9d %-9d",
		labels.PlanetName(r.Index),
		labels.Owner(r.Owner),
		labels.BuildingName(r.Type),
		fmt.Sprintf("(%d, %d)", r.X, r.Y),
		r.Powered,
		r.Repairing,
		r.Damage,
		r.Remaining,
		r.Efficiency,
		r.Operational,
		r.Unused,
		r.Kwh,
		r.Workers)
}

func WriteBuilding(b Building, f *os.File) (err error) {
	err = binary.Write(f, binary.LittleEndian, b)
	return
}

func ReadBuilding(f *os.File) (b Building) {
	_ = binary.Read(f, binary.LittleEndian, &b)
	return
}
