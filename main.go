package main

import (
	"fmt"
	"github.com/spatial-go/geos/pattern"
)

func main() {
	wkt := pattern.NewWKT()
	fromWKT := wkt.FromWKTStr("POINT(10 10)")
	fmt.Printf("%p\r\n", &fromWKT)

	s, e := wkt.ToWKTStr(fromWKT)
	if e != nil {
		fmt.Printf("错误%s", e.Error())
	} else {
		fmt.Println(s)
	}
	wkt.Destroy()
}
