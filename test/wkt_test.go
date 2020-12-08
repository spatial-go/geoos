package test

import (
	"fmt"
	"github.com/spatial-go/geos/coder"
	"testing"
)

func TestFromWkt(t *testing.T){
	fromWKT := coder.FromWKTStr("POINT(10 10)")
	fmt.Printf("%p\r\n", &fromWKT)

	s, e := coder.ToWKTStr(fromWKT)
	if e != nil {
		t.Error( e.Error())
	} else {
		t.Log(s)
	}

}