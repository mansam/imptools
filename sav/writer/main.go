package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mansam/imptools/sav/structs"
	"gopkg.in/yaml.v3"
)

var (
	SaveFile  string
	Index     int
	InputFile string
	Append    bool
)

func main() {
	if len(os.Args) != 4 {
		println("usage: fleetwriter sav infile idx")
	}
	var err error
	SaveFile = os.Args[1]
	InputFile = os.Args[2]
	save, err := os.OpenFile(SaveFile, os.O_RDWR, 000)
	if err != nil {
		panic(err)
	}
	defer save.Close()

	in, err := os.Open(InputFile)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	numFleets := structs.FleetCount(save)
	fmt.Println(numFleets)
	if strings.Contains(os.Args[3], "a") {
		numFleets++
		Index = numFleets
	} else {
		Index, err = strconv.Atoi(os.Args[3])
		if err != nil {
			panic(err)
		}
	}

	fleet := structs.Fleet{}
	decoder := yaml.NewDecoder(in)
	decoder.KnownFields(true)
	err = decoder.Decode(&fleet)
	if err != nil {
		panic(err)
	}

	_, err = save.Seek(structs.NumFleetsOffset, 0)
	if err != nil {
		panic(err)
	}
	binary.Write(save, binary.LittleEndian, uint16(numFleets))
	_, err = save.Seek(int64(structs.FleetsOffset)+93*max(int64(Index-1), 0), 0)
	if err != nil {
		panic(err)
	}
	err = structs.WriteFleet(fleet, save)
	if err != nil {
		panic(err)
	}

}
