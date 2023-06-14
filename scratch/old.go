package main

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 3 {
		panic(fmt.Sprintf("usage: %s packed unpacked/", os.Args[0]))
	}
	filepath := os.Args[1]
	outpath := os.Args[2]
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = os.Mkdir(outpath, 0755)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 2)
	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	if n != 2 {
		panic("couldn't read number of entries")
	}

	//entries := int(buf[1])*256 + int(buf[0])
	entries := btoi(buf)
	fmt.Printf("Entries: %d\n", entries)

	for i := 0; i < entries; i++ {
		buf = readn(f, 1)
		nameLen := int(buf[0])
		buf = readn(f, 13)
		packedName := string(buf[:nameLen])
		packedLength := btoi(readn(f, 2))
		packedOffset := btoi(readn(f, 4))
		fmt.Printf("%d %s %d %d\n", nameLen, packedName, packedLength, packedOffset)
		unpack(f, path.Join(outpath, packedName), packedLength, packedOffset)
	}
}

func unpack(f *os.File, path string, n, offset int) {
	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	buf := readnat(f, n, offset)
	written, err := out.Write(buf)
	if err != nil {
		panic(err)
	}
	if written != n {
		err = errors.New(fmt.Sprintf("couldn't write %d bytes (wrote %d bytes)", n, written))
		panic(err)
	}

}

func btoi(b []byte) (x int) {
	factor := 1
	for i := 0; i < len(b); i++ {
		x += int(b[i]) * factor
		factor *= 256
	}
	return
}

func readnat(f *os.File, n, offset int) (b []byte) {
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

func readn(f *os.File, n int) (b []byte) {
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
