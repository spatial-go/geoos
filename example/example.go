package main

import (
	"encoding/json"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/geojson"
)

func main() {
	//// First, choose the default algorithm.
	//strategy := geoos.NormalStrategy()
	//// Secondly, manufacturing test data and convert it to geometry
	//const wkt = `POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))`
	//geometry, _ := geoos.UnmarshalString(wkt)
	//// Last， call the Area () method and get result.
	//area, e := strategy.Area(geometry)
	//if e != nil {
	//	fmt.Printf(e.Error())
	//}
	//fmt.Printf("%f", area)
	//// get result 4.0

	test()
}

func test() {
	rawJSON := []byte(`
  { "type": "FeatureCollection",
	"features": [
	  { "type": "Feature",
		"geometry": {"type": "Point", "coordinates": [102.0, 0.5]},
		"properties": {"prop0": "value0"}
	  }
	]
  }`)

	//fc, _ := geojson.UnmarshalFeatureCollection(rawJSON)
	//println("%p", fc)
	// or

	fc := geojson.NewFeatureCollection()
	_ = json.Unmarshal(rawJSON, &fc)
	println("%p", fc)

	// Geometry will be unmarshalled into the correct geo.Geometry type.
	point := fc.Features[0].Geometry.(geoos.Point)
	println("%p", &point)
}
