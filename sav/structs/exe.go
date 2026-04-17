package structs

import (
	"encoding/binary"
	"os"
)

const UnknownStructEntryOffset = 0xa6df5

const UnknownStruct156Offset = 0xa8fdc

type UnknownStruct156 struct {
	Fields [156]byte
}

func ReadUnknownStruct156(f *os.File) (u UnknownStruct156) {
	_ = binary.Read(f, binary.LittleEndian, &u)
	return
}

type UnknownStructEntry22 struct {
	Type       uint8
	NameLength uint8
	Name_      [18]byte
	Indicator  uint8
	U1         [1]byte
}

func (r UnknownStructEntry22) Name() string {
	return string(r.Name_[:r.NameLength])
}

type UnknownStructEntry14 struct {
	NameLength uint8
	Name_      [10]byte
	U0         [3]byte
}

func (r UnknownStructEntry14) Name() string {
	return string(r.Name_[:r.NameLength])
}

type UnknownShipEntry struct {
	U0          [1]byte
	IndexLength uint8
	Index       [3]byte
	U1          [16]byte
	NameLength  uint8
	Name_       [18]byte
	U2          [4]byte
}

func (r UnknownShipEntry) Name() string {
	return string(r.Name_[:r.NameLength])
}

func ReadUnknownShipEntry(f *os.File) (u UnknownShipEntry) {
	_ = binary.Read(f, binary.LittleEndian, &u)
	return
}

func ReadUnknownStructEntry22(f *os.File) (u UnknownStructEntry22) {
	_ = binary.Read(f, binary.LittleEndian, &u)
	return
}

func ReadUnknownStructEntry14(f *os.File) (u UnknownStructEntry14) {
	_ = binary.Read(f, binary.LittleEndian, &u)
	return
}
