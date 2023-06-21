package structs

import (
	"encoding/binary"
	"fmt"
	"os"
)

// Offset: 0xEA53 (93 bytes)
type Fleet struct {
	NameLength    uint8
	Name_         [12]byte
	Unknown       [9]byte
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
	Unused        int32 // Unused vehicle slots, always 0. These are counted against fleet capacity but are otherwise ignored.
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
	NameLength uint8
	Name_      [12]byte
	Unknown    [37]byte
	Fleet      uint8
	/*
		08 54 72 61 64 65 72 20 31 62 62 48 43 0C 02 01
		01 01 01 01 01 01 00 00 00 00 01 01 01 01 01 01
		00 00 00 00 FF FF FF FF FF FF 00 00 00 00 00 00
		00 00 02
	*/
}

func (r Ship) String() string {
	return fmt.Sprintf("%-12s %-3d", r.Name(), r.Fleet)
}

func (r Ship) Name() string {
	return string(r.Name_[:r.NameLength])
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
