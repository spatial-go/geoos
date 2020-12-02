package pattern

import (
	"fmt"
	"testing"
)

func TestFromWkt(t *testing.T){
	wkt := NewWKT()
	fromWKT := wkt.FromWKTStr("POINT(10 10)")
	fmt.Printf("%p\r\n", &fromWKT)

	s, e := wkt.ToWKTStr(fromWKT)
	if e != nil {
		t.Error( e.Error())
	} else {
		t.Log(s)
	}
	wkt.Destroy()

}