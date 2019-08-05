package binaryplist

import (
	"encoding/binary"
	"log"
	"os"
)

//BinaryPlist - Binary plist struct
type BinaryPlist struct {
	SortVersion      uint8
	OffsetSize       uint8
	RefSize          uint8
	NumObjects       uint64
	TopObjectOffset  uint64
	OffsetTableStart uint64
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

//BinaryPlist - Constructor that takes in a file and returns a BinaryPlist struct
func NewBinaryPlist(file string) BinaryPlist {
	//instance vars
	var (
		bp        BinaryPlist
		fileStats os.FileInfo
		err       error
	)
	trailer := make([]byte, 26)

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
	bp.SortVersion = uint8(trailer[5])
	bp.OffsetSize = uint8(trailer[6])
	bp.RefSize = uint8(trailer[7])
	bp.NumObjects = binary.BigEndian.Uint64(trailer[8:15])
	bp.TopObjectOffset = binary.BigEndian.Uint64(trailer[16:23])
	bp.OffsetTableStart = binary.BigEndian.Uint64(trailer[24:31])
	bp.ObjectTable = trailer[bp.TopObjectOffset : bp.OffsetTableStart-1]
	bp.OffsetTable = trailer[bp.OffsetTableStart : trailerOffset-1]

	return bp
}
