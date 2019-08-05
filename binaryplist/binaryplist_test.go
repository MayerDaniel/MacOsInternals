package binaryplist

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
	if !isPlist(pl) {
		t.Errorf("Incorrectly identified a bplist as non bplist")
	}
	if isPlist(npl) {
		t.Errorf("Incorrectly identified a non bplist as bplist")
	}

}

func TestParsePlist(t *testing.T) {
	bplistPath := "/github.com/mayerdaniel/MacOSInternals/test/bin.plist"
	bp := parsePlist(bplistPath)
	var flag bool
	flag = true
	if bp.SortVersion != 0 {
		flag = false
	}
	//complete

}
