package structs

import (
	"fmt"
	"github.com/mansam/imptools/reader"
	"github.com/mansam/imptools/sav/labels"
	"os"
)

// Building: 18 bytes
type Building struct {
	Type       int  // 1
	Owner      int  // 1
	Powered    bool // 1
	Repairing  bool // 1
	Damage     int  // 2
	Remaining  int  // 1
	Index      int  // 1
	X          int  // 1
	Y          int  // 1
	Efficiency int  // 1
	Flag1      int  // 1
	Short1     int  // 2
	Kwh        int  // 2
	Workers    int  // 2
}

func (r Building) Complete() int {
	return 100 - (r.Remaining / 2)
}

func (r Building) String() string {
	return fmt.Sprintf("%-12s %-25s %-20s %-9v %-9v %-9d %-9d %-9d %-9d %-9d %-9d %-9d",
		labels.PlanetName(r.Index),
		labels.Owner(r.Owner),
		labels.BuildingName(r.Type),
		r.Powered,
		r.Repairing,
		r.Damage,
		r.Complete(),
		r.Efficiency,
		r.Flag1,
		r.Short1,
		r.Kwh,
		r.Workers)
}

func ReadBuilding(f *os.File) (b Building) {
	b.Type = reader.Btoi(reader.ReadN(f, 1))
	b.Owner = reader.Btoi(reader.ReadN(f, 1))
	b.Powered = reader.Bool(reader.ReadN(f, 1)[0])
	b.Repairing = reader.Bool(reader.ReadN(f, 1)[0])
	b.Damage = reader.Btoi(reader.ReadN(f, 2))
	b.Remaining = reader.Btoi(reader.ReadN(f, 1))
	b.Index = reader.Btoi(reader.ReadN(f, 1))
	b.X = reader.Btoi(reader.ReadN(f, 1))
	b.Y = reader.Btoi(reader.ReadN(f, 1))
	b.Efficiency = reader.Btoi(reader.ReadN(f, 1))
	b.Flag1 = reader.Btoi(reader.ReadN(f, 1))
	b.Short1 = reader.Btoi(reader.ReadN(f, 2))
	b.Kwh = reader.Btoi(reader.ReadN(f, 2))
	b.Workers = reader.Btoi(reader.ReadN(f, 2))
	return
}
