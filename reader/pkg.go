package reader

import (
	"errors"
	"fmt"
	"os"
)

func Unpack(f *os.File, path string, n, offset int) {
	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	buf := ReadNAt(f, n, offset)
	written, err := out.Write(buf)
	if err != nil {
		panic(err)
	}
	if written != n {
		err = errors.New(fmt.Sprintf("couldn't write %d bytes (wrote %d bytes)", n, written))
		panic(err)
	}

}

func Btoi(b []byte) (x int) {
	factor := 1
	for i := 0; i < len(b); i++ {
		x += int(b[i]) * factor
		factor *= 256
	}
	return
}

func Bool(b byte) bool {
	return b > 0
}

func ReadNAt(f *os.File, n, offset int) (b []byte) {
	b = make([]byte, n)
	read, err := f.ReadAt(b, int64(offset))
	if err != nil {
		panic(err)
	}
	if read != n {
		err = errors.New(fmt.Sprintf("couldn't read %d bytes at offset %d", n, offset))
		panic(err)
	}
	return
}

func ReadN(f *os.File, n int) (b []byte) {
	b = make([]byte, n)
	read, err := f.Read(b)
	if err != nil {
		panic(err)
	}
	if read != n {
		err = errors.New(fmt.Sprintf("couldn't read %d bytes", n))
		panic(err)
	}
	return
}
