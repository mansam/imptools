package pac

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/mansam/imptools/reader"
	"os"
)

// A HeaderEntry is 20 bytes
type HeaderEntry struct {
	FilenameLength byte
	Filename       [13]byte
	Length         uint16
	Offset         uint32
}

func (h HeaderEntry) String() string {
	return fmt.Sprintf("%d %s %d %d", h.FilenameLength, h.Filename, h.Length, h.Offset)
}

func WriteHeaderEntry(h HeaderEntry, f *os.File) (err error) {
	err = binary.Write(f, binary.LittleEndian, h)
	return
}

func ReadHeaderEntry(f *os.File) (h HeaderEntry) {
	_ = binary.Read(f, binary.LittleEndian, &h)
	return
}

func Unpack(f *os.File, h HeaderEntry) (err error) {
	out, err := os.Create(string(h.Filename[:h.FilenameLength]))
	if err != nil {
		return
	}
	defer out.Close()
	buf := reader.ReadNAt(f, int(h.Length), int(h.Offset))
	written, err := out.Write(buf)
	if err != nil {
		return
	}
	if written != int(h.Length) {
		err = errors.New(fmt.Sprintf("couldn't write %d bytes (wrote %d bytes)", h.Length, written))
	}
	return
}
