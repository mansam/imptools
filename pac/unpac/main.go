package main

import (
	"fmt"
	"github.com/mansam/imptools/pac"
	"github.com/mansam/imptools/reader"
	"os"
)

var (
	InputFilePath   string
	OutputDirectory string
)

func main() {
	if len(os.Args) != 3 {
		panic(fmt.Sprintf("usage: %s packed unpacked/", os.Args[0]))
	}
	InputFilePath = os.Args[1]
	OutputDirectory = os.Args[2]

	f, err := os.Open(InputFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = os.Mkdir(OutputDirectory, 0755)
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

	entries := reader.Btoi(buf)
	fmt.Printf("Entries: %d\n", entries)

	for i := 0; i < entries; i++ {
		entry := pac.ReadHeaderEntry(f)
		fmt.Println("\t", entry)
		err = pac.Unpack(f, entry, OutputDirectory)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%d files unpacked successfully.\n", entries)
}
