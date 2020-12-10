package test

import (
	"github.com/spatial-go/geos/geo"
	"testing"
)

func TestWkb(t *testing.T){

	var testByte = []byte{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 64, 93, 192, 0, 0, 0, 0, 0, 128, 65, 64}
	var testPoint = "POINT(-117 35)"
	var testHex = "01010000000000000000405DC00000000000804140"

	geometry, e := geo.GeomFromWKBStr(testByte)
	if e != nil {
		t.Error(e.Error())
	}
	s, _ := geo.ToWKTStr(geometry)
	bytes, _ := geo.ToWKB(geometry)
	t.Log(s)
	t.Log(testPoint)
	t.Log(bytes)

	geosGeometry, _ := geo.GeomFromWKBHexStr(testHex)
	wkbstr, _ := geo.ToWKB(geosGeometry)
	t.Log(wkbstr)
	hex, _ := geo.ToWKBHex(geosGeometry)
	t.Log(hex)

}