package test

import (
	"fmt"
	"github.com/spatial-go/geos/geo"
	"testing"
)

func TestFromWkt(t *testing.T){
	fromWKT := geo.GeomFromWKTStr("POINT(10 10)")
	fmt.Printf("%p\r\n", &fromWKT)

	s, e := geo.ToWKTStr(fromWKT)
	if e != nil {
		t.Error( e.Error())
	} else {
		t.Log(s)
	}

}