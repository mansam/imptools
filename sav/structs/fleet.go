package structs

import (
	"encoding/binary"
	"fmt"
	"github.com/mansam/imptools/sav/labels"
	"os"
)

// Offset: 0xEA53 (93 bytes)
type Fleet struct {
	NameLength    uint8
	Name_         [12]byte
	Owner         uint8
	Color         uint8 // affects appearance of empire ships (yellow w/ blue, yellow /w red, etc)
	Unknown       [2]byte
	Controllable  uint8
	Visible       uint8
	Alert         uint8
	Unknown0      [4]byte
	X             uint16
	Unknown1      uint16
	Y             uint16
	Unknown2      [18]byte
	Fighter1      int16
	Fighter2      int16
	Fighter3      int16
	Fighter4      int16
	Fighter5      int16
	Fighter6      int16
	Tank1         int16
	Tank2         int16
	Tank3         int16
	Tank4         int16
	Unused        [4]byte // Unused vehicle slots, always 0. These are counted against fleet capacity but are otherwise ignored.
	Car1          int16
	Car2          int16
	Car3          int16
	TotalFighters int16 // (calculated)
	TotalVehicles int16 // (calculated)
	Unknown3      int16
	Unknown4      int16
	Unknown5      int16
	Unknown6      int16
	Unknown7      [5]byte
}

func (r Fleet) Name() string {
	return string(r.Name_[:r.NameLength])
}

func (r Fleet) String() string {
	coords := fmt.Sprintf("(%d, %d)", r.X, r.Y)
	return fmt.Sprintf("%-12s %-16s %-4d %-4d %-4d %-4d %-4d %-4d %-4d %-4d %-4d %-4d %-4d %-4d %-4d %-8d %-8d %-4d",
		r.Name(),
		coords,
		r.Fighter1,
		r.Fighter2,
		r.Fighter3,
		r.Fighter4,
		r.Fighter5,
		r.Fighter6,
		r.Tank1,
		r.Tank2,
		r.Tank3,
		r.Tank4,
		r.Car1,
		r.Car2,
		r.Car3,
		r.TotalFighters,
		r.TotalVehicles,
		r.Unknown5)
}

// Offset: 0x2261 (51 bytes)
type Ship struct {
	NameLength           uint8
	Name_                [12]byte
	Owner                uint8
	TechSubcategory      uint8
	TechSubcategoryIndex uint8
	Slot1                uint8
	Slot2                uint8
	Slot3                uint8
	Slot4                uint8
	Slot5                uint8
	Slot6                uint8
	Slot7                uint8
	Slot8                uint8
	Slot9                uint8
	Slot10               uint8
	Slot1Count           uint8
	Slot2Count           uint8
	Slot3Count           uint8
	Slot4Count           uint8
	Slot5Count           uint8
	Slot6Count           uint8
	Slot7Count           uint8
	Slot8Count           uint8
	Slot9Count           uint8
	Slot10Count          uint8
	Damage1              uint8
	Damage2              uint8
	Damage3              uint8
	Damage4              uint8
	Damage5              uint8
	Damage6              uint8
	Damage7              uint8
	Damage8              uint8
	Damage9              uint8
	Damage10             uint8
	GroundForces         uint16
	Kills                uint16
	Fleet                uint8
	/*
		08 54 72 61 64 65 72 20 31 62 62 48 43 0C 02 01
		01 01 01 01 01 01 00 00 00 00 01 01 01 01 01 01
		00 00 00 00 FF FF FF FF FF FF 00 00 00 00 00 00
		00 00 02
	*/
}

//
// Offset 0xa9078 in MAIN.EXE: Possible start of ship equipment slot definitions
//
// Human Flagship 1 slots
// Module
// Radar
// Guns
// Missiles
// Bombs
// Shield
// Lasers
// Hyperdrive

func (r Ship) String() string {
	if r.TechSubcategory == 3 {
		return fmt.Sprintf("%-12s %-3d %-24s %-16s %-30s %-30s %-30s %-30s %-30s %-30s %-30s %-30s %2d %2d",
			r.Name(),
			r.Fleet,
			labels.Owner(r.Owner),
			r.Class(),
			// only accurate for human Flagship 1
			Slot(Module(r.Slot1), r.Slot1Count, r.Damage1),
			Slot(Radar(r.Slot2), r.Slot2Count, r.Damage2),
			Slot(Gun(r.Slot3), r.Slot3Count, r.Damage3),
			Slot(Missile(r.Slot4), r.Slot4Count, r.Damage4),
			Slot(Missile(r.Slot5), r.Slot5Count, r.Damage5),
			Slot(Shield(r.Slot6), r.Slot6Count, r.Damage6),
			Slot(Laser(r.Slot7), r.Slot7Count, r.Damage7),
			Slot(Hyperdrive(r.Slot8), r.Slot8Count, r.Damage8),
			r.GroundForces,
			r.Kills)
	} else if r.TechSubcategory == 2 {
		return fmt.Sprintf("%-12s %-3d %-24s %-20s %-20s %-20s %-20s %-20s %-20s %-20s %x %x %x %x",
			r.Name(),
			r.Fleet,
			labels.Owner(r.Owner),
			r.Class(),
			// only accurate for human cruiser 2
			Slot(Shield(r.Slot1), r.Slot1Count, r.Damage1),
			Slot(Gun(r.Slot2), r.Slot2Count, r.Damage2),
			Slot(Laser(r.Slot3), r.Slot3Count, r.Damage3),
			Slot(Gun(r.Slot4), r.Slot4Count, r.Damage4),
			Slot(Radar(r.Slot5), r.Slot5Count, r.Damage5),
			Slot(Hyperdrive(r.Slot6), r.Slot6Count, r.Damage6),
			r.Slot7,
			r.Slot8,
			r.Slot9,
			r.Slot10)
	} else {
		return r.Name()
	}
}

func (r Ship) Name() string {
	return string(r.Name_[:r.NameLength])
}

func (r Ship) Class() string {
	return labels.TechnologyName((r.TechSubcategory-1)*6 + r.TechSubcategoryIndex)
}

func Slot(s string, c uint8, d uint8) string {
	if c == 0 {
		return "<empty>"
	}
	return fmt.Sprintf("%s (%d) (%3d%%)", s, c, uint8(float32(d)/float32(255)*100))
}

func Missile(m uint8) string {
	if m == 0 {
		return "n/a"
	}
	return labels.TechnologyName(TechMissiles + m)
}

func Laser(m uint8) string {
	if m == 0 {
		return "n/a"
	}
	return labels.TechnologyName(TechLasers + m)
}

func Gun(m uint8) string {
	if m == 0 {
		return "n/a"
	}
	return labels.TechnologyName(TechGuns + m)
}

func Radar(r uint8) string {
	if r == 0 {
		return "n/a"
	}
	return labels.TechnologyName(TechRadars + r)
}

func Module(m uint8) string {
	if m == 0 {
		return "n/a"
	}
	return labels.TechnologyName(TechModules + m)
}

func Hyperdrive(h uint8) string {
	if h == 0 {
		return "n/a"
	}
	return labels.TechnologyName(TechHyperdrives + h)
}

func Shield(s uint8) string {
	if s == 0 {
		return "n/a"
	}
	return labels.TechnologyName(TechShields + s)
}

func ReadShip(f *os.File) (s Ship) {
	_ = binary.Read(f, binary.LittleEndian, &s)
	return
}

func ReadFleet(f *os.File) (flt Fleet) {
	_ = binary.Read(f, binary.LittleEndian, &flt)
	return
}

/*
FLeet 1: Destroyer 1 "Exalibur" and 2 Fighter 1.

07 46 6C 65 65 74 20 31 00 00 00 00 00 01 01 00
01 01 01 00 00 E0 AB 01 00 08 52 03 00 1C 03 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 02 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 02 00 00 00
07 00 00 00 00 00 00 01 00 00 03 03 03

*/

/*
XYZXYZXYZXYZ: 1 Fighter 1

0C 58 59 5A 58 59 5A 58 59 5A 58 59 5A 01 00 00
01 01 01 00 00 E0 AB 01 00 D8 59 03 00 1C 03 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 01 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00
00 00 00 00 00 00 00 00 00 00 03 03 03

*/

/*
Garthog flt

0B 47 61 72 74 68 6F 67 20 66 6C 74 00 02 0B 05
00 00 01 00 00 60 DA 01 00 B0 71 03 00 20 01 01
3F 00 00 00 00 00 00 00 00 00 00 00 00 00 09 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 09 00 00 00
36 00 00 00 0C 00 02 05 00 00 03 03 03

*/

/*
Thorin Flt Save15

0A 54 68 6F 72 69 6E 20 46 6C 74 71 71 01 01 00
01 01 01 00 00 C8 53 01 00 FC E9 02 00 1C 01 01
3D 00 00 00 98 03 00 00 00 00 71 71 00 00 02 00
09 00 06 00 00 00 05 00 01 00 14 00 00 00 00 00
03 00 00 00 00 00 02 00 01 00 04 00 17 00 1E 00
B8 00 2E 00 1F 00 03 11 00 01 03 03 03
*/
