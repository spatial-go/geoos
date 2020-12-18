package main

import "github.com/spatial-go/geos"

func main() {
	const wkt = `POLYGON((1 2, 3 4, 5 6, 1 2))`
	geometry, _ := geos.UnmarshalString(wkt)
	strategy := geos.NormalStrategy()
	b, _ := strategy.IsSimple(geometry)
	if b {
		println(b)
	}
}
