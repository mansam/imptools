package structs

import (
	"encoding/binary"
	"fmt"
	"github.com/mansam/imptools/sav/labels"
	"os"
)

type Header struct {
	NameLength uint8
	SaveName   [26]byte
	Unknown    byte
	Rank       uint16
	Year       uint16
	Month      uint16
	Day        uint16
	Hour       uint16
	Minute     uint16
	Unknown2   [5]byte
	Money      int32
}

func (r Header) String() string {
	return fmt.Sprintf("%-30s %-2x %-13s %s %x %d", r.SaveName[:r.NameLength], r.Unknown, labels.Rank(r.Rank), r.Date(), r.Unknown2, r.Money)
}

func (r Header) Date() string {
	return fmt.Sprintf("%s %d %04d %02d:%02d", labels.Month(r.Month), r.Day, r.Year, r.Hour, r.Minute)
}

func ReadHeader(f *os.File) (h Header) {
	binary.Read(f, binary.LittleEndian, &h)
	return
}
