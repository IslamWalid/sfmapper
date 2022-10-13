package main

import (
	"fmt"
	"os"
	"time"

    "github.com/IslamWalid/struct-file-mapper"
)

type structure struct {
	String       string
	Int          int
	Bool         bool
	SubStructure subStructure
}

type subStructure struct {
	Float float32
}

func Routine(input *structure) {
	time.Sleep(500 * time.Millisecond)
	input.String = "new string"
} 

func main() {
	var err error
	input := &structure{
		String: "str",
		Int:    18,
		Bool:   true,
		SubStructure: subStructure{
			Float: 1.3,
		},
	}

	err = os.MkdirAll("mnt", 0777)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	go Routine(input)
	err = fs_mapper.Mount("mnt", input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
