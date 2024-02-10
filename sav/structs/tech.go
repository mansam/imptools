package structs

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Technology struct {
	// 0 Unused, 1 Hidden, 2 Can't Start, 3 Can Start, 4 Started, 5 Complete
	State      uint8  //1
	Count      uint16 //2
	Price      uint32 //4
	Resale     uint16 //2
	Producible bool   //1 available on the production screen. Buildings and the Thorin = 0
	MinFunding uint32 //4
	MaxFunding uint32 //4
	// PreReqs are 3 bytes indicating
	// 1. The research category (Spaceships, Equipment, Weapons, Buildings 1-4)
	// 2. The research subcategory (Fighters, Destroyers/Cruisers, etc 1-6)
	// 3. The slot number (1-6)
	PreReq1          [3]byte //3
	PreReq2          [3]byte //3
	PreReq3          [3]byte //3
	RankVisible      uint8   //1
	RankResearchable uint8   //1
	Civil            uint8   //1
	Mechanical       uint8   //1
	Computer         uint8   //1
	AI               uint8   //1
	Military         uint8   //1
	Unknown1         uint32  //4
	Unknown2         uint32  //4
}

func (r Technology) String() string {
	return fmt.Sprintf(
		"%d %d %d %d %v %d %d %d %d %d",
		r.State, r.Count, r.Price, r.Resale, r.Producible, r.Civil, r.Mechanical, r.Computer, r.AI, r.Military)
}

func WriteTechnology(t Technology, f *os.File) (err error) {
	err = binary.Write(f, binary.LittleEndian, t)
	return
}

func ReadTechnology(f *os.File) (t Technology) {
	_ = binary.Read(f, binary.LittleEndian, &t)
	return
}
