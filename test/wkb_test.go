package test

import (
	"github.com/spatial-go/geos/coder"
	"testing"
)

func TestWkb(t *testing.T){

	var testByte = []byte{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 64, 93, 192, 0, 0, 0, 0, 0, 128, 65, 64}
	var testPoint = "POINT(-117 35)"
	var testHex = "01010000000000000000405DC00000000000804140"

	geometry, e := coder.FromWKB(testByte)
	if e != nil {
		t.Error(e.Error())
	}
	s, _ := coder.ToWKTStr(geometry)
	bytes, _ := coder.ToWKB(geometry)
	t.Log(s)
	t.Log(testPoint)
	t.Log(bytes)

	geosGeometry, _ := coder.FromWKBHex(testHex)
	wkbstr, _ := coder.ToWKB(geosGeometry)
	t.Log(wkbstr)
	hex, _ := coder.ToWKBHex(geosGeometry)
	t.Log(hex)

}