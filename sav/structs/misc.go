package structs

import (
	"encoding/binary"
	"os"
)

// SpecialMessage is a longer scrolling message that starts after the save game name. Is this possibly
// the active scrolling message? or something else?
type SpecialMessage struct {
	Length  uint8
	Message [102]byte
}

func (r SpecialMessage) String() string {
	return string(r.Message[:r.Length])
}

type ScrollingMessage struct {
	Length  uint8
	Message [75]byte
}

func (r ScrollingMessage) String() string {
	return string(r.Message[:r.Length])
}

func WriteScrollingMessage(s ScrollingMessage, f *os.File) (err error) {
	err = binary.Write(f, binary.LittleEndian, s)
	return
}

func ReadScrollingMessage(f *os.File) (s ScrollingMessage) {
	_ = binary.Read(f, binary.LittleEndian, &s)
	return
}

func ReadSpecialMessage(f *os.File) (s SpecialMessage) {
	_ = binary.Read(f, binary.LittleEndian, &s)
	return
}
