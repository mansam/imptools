package main

import (
	"encoding/binary"
	"fmt"
	"github.com/mansam/imptools/pac"
	"math"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 3 {
		panic(fmt.Sprintf("usage: %s unpacked/ packed", os.Args[0]))
	}
	unpackedFileDir := os.Args[1]
	outfile, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	entries, err := os.ReadDir(unpackedFileDir)
	if err != nil {
		panic(err)
	}

	var badNames []string
	var anticipatedSize int64
	var fileEntries int
	for _, entry := range entries {
		if entry.IsDir() {
			// only pack files in the root
			continue
		}
		if len(entry.Name()) > 13 {
			badNames = append(badNames, entry.Name())
		}
		fileInfo, rErr := entry.Info()
		if rErr != nil {
			panic(rErr)
		}
		anticipatedSize += fileInfo.Size()
		fileEntries++
	}
	if fileEntries > math.MaxUint16 || len(badNames) > 0 || anticipatedSize > math.MaxUint32 {
		if len(badNames) > 0 {
			println("Filenames too long to pac:")
			for _, name := range badNames {
				println("\t" + name)
			}
		}
		if fileEntries > math.MaxUint16 {
			println("Too many files to pac.")
			println("\t Total number of files: ", fileEntries)
			println("\t Overage: ", fileEntries-math.MaxUint16)
		}
		if anticipatedSize > math.MaxUint32 {
			println("Total length of packed archive would exceed uint32:")
			println("\tTotal anticipated size: ", anticipatedSize)
			println("\tOverage: ", anticipatedSize-math.MaxUint32)
		}
		os.Exit(1)
	}

	// offset the data region by the 2 bytes for the number of entries, plus
	// the number of entries * 20 bytes to account for the header.
	dataOffset := fileEntries*pac.HeaderEntryLength + 2
	var headers []pac.HeaderEntry
	var data []byte
	for _, entry := range entries {
		if entry.IsDir() {
			// only pack files in the root
			continue
		}
		entryFileName := path.Join(unpackedFileDir, entry.Name())
		bytes, rErr := os.ReadFile(entryFileName)
		if rErr != nil {
			panic(err)
		}

		nameLength := len(entry.Name())
		header := pac.HeaderEntry{}
		var nameBytes [13]byte
		for i := 0; i < nameLength; i++ {
			nameBytes[i] = entry.Name()[i]
		}
		header.FilenameLength = byte(nameLength)
		header.Filename = nameBytes
		header.Length = uint16(len(bytes))
		header.Offset = uint32(len(data) + dataOffset)
		data = append(data, bytes...)
		headers = append(headers, header)
	}
	err = binary.Write(outfile, binary.LittleEndian, uint16(len(headers)))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Entries: %d\n", len(headers))
	for _, header := range headers {
		err = pac.WriteHeaderEntry(header, outfile)
		if err != nil {
			panic(err)
		}
		fmt.Println("\t", header)
	}
	n, err := outfile.Write(data)
	if err != nil {
		panic(err)
	}
	if n != len(data) {
		fmt.Printf("Only wrote %d of %d data bytes.\n", n, len(data))
		os.Exit(1)
	}
	fmt.Printf("%d files packed successfully.\n", len(headers))
}
