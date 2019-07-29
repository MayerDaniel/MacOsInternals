package main

import (
	"fmt"
	"log"
	"os"
)

//BinaryPlist - Binary plist struct
type BinaryPlist struct {
	OffsetSize       byte
	RefSize          byte
	NumObjects       uint8
	TopObjectOffset  uint8
	OffsetTableStart uint8
	ObjectTable      []byte
	OffsetTable      []byte
}

func isPlist(file *os.File) bool { //turn in to get header
	bytes := make([]byte, 8)

	_, err := file.Read(bytes)
	if err != nil {
		panic(err)
	}

	if string(bytes) == "bplist00" {
		return true
	}
	return false
}

//make get trailer

func parsePlist(file string) BinaryPlist {
	//instance vars
	var (
		bp        BinaryPlist
		fileStats os.FileInfo
		err       error
	)
	trailer := make([]byte, 32)

	//open file and grab its stats
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	fileStats, err = f.Stat()
	if err != nil {
		panic(err)
	}

	//checks to see if the file is in fact a binary plist
	if !isPlist(f) {
		log.Fatal("File is not a binary plist: ", fileStats.Name())
		os.Exit(0)
	}

	//parse the fixed length trailer
	trailerOffset := fileStats.Size() - 32
	_, err = f.ReadAt(trailer, trailerOffset)
	return bp
}

func main() {
	fmt.Println("hello world")
}
