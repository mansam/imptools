package structs

import (
	"fmt"
	"github.com/mansam/imptools/reader"
	"os"
)

type Technology struct {
	// 0 Unused, 1 Hidden, 2 Can't Start, 3 Can Start, 4 Started, 5 Complete
	State      int //1
	Count      int //2
	Price      int //4
	Resale     int //2
	Producible int //1 available on the production screen. Buildings and the Thorin = 0
	MinFunding int //4
	MaxFunding int //4
	// PreReqs are 3 bytes indicating
	// 1. The research category (Spaceships, Equipment, Weapons, Buildings 1-4)
	// 2. The research subcategory (Fighters, Destroyers/Cruisers, etc 1-6)
	// 3. The slot number (1-6)
	PreReq1          []byte //3
	PreReq2          []byte //3
	PreReq3          []byte //3
	RankVisible      int    //1
	RankResearchable int    //1
	Civil            int    //1
	Mechanical       int    //1
	Computer         int    //1
	AI               int    //1
	Military         int    //1
	Unknown1         int    //4
	Unknown2         int    //4
}

func (r Technology) String() string {
	return fmt.Sprintf(
		"%-8d \t%-8d",
		r.Unknown1, r.Unknown2)
}

func ReadTechnology(f *os.File) (t Technology) {
	t.State = reader.Btoi(reader.ReadN(f, 1))
	t.Count = reader.Btoi(reader.ReadN(f, 2))
	t.Price = reader.Btoi(reader.ReadN(f, 4))
	t.Resale = reader.Btoi(reader.ReadN(f, 2))
	t.Producible = reader.Btoi(reader.ReadN(f, 1))
	t.MinFunding = reader.Btoi(reader.ReadN(f, 4))
	t.MaxFunding = reader.Btoi(reader.ReadN(f, 4))
	t.PreReq1 = reader.ReadN(f, 3)
	t.PreReq2 = reader.ReadN(f, 3)
	t.PreReq3 = reader.ReadN(f, 3)
	t.RankVisible = reader.Btoi(reader.ReadN(f, 1))
	t.RankResearchable = reader.Btoi(reader.ReadN(f, 1))
	t.Civil = reader.Btoi(reader.ReadN(f, 1))
	t.Mechanical = reader.Btoi(reader.ReadN(f, 1))
	t.Computer = reader.Btoi(reader.ReadN(f, 1))
	t.AI = reader.Btoi(reader.ReadN(f, 1))
	t.Military = reader.Btoi(reader.ReadN(f, 1))
	t.Unknown1 = reader.Btoi(reader.ReadN(f, 4))
	t.Unknown2 = reader.Btoi(reader.ReadN(f, 4))
	return
}
