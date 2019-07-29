package main

import (
	"os"
	"testing"
)

func TestIsPlist(t *testing.T) {
	bplistPath := "/Users/danniboi/Desktop/bin.plist"
	xmlplistPath := "/Users/danniboi/Desktop/xml.plist"

	pl, err := os.Open(bplistPath)
	if err != nil {
		t.Errorf("error importing file from path %s", bplistPath)
	}
	npl, err := os.Open(xmlplistPath)
	if err != nil {
		t.Errorf("error importing file from path %s", xmlplistPath)
	}
	if isPlist(pl) != true {
		t.Errorf("Incorrectly identified a bplist as non bplist")
	}
	if isPlist(npl) != false {
		t.Errorf("Incorrectly identified a non bplist as bplist")
	}

}
