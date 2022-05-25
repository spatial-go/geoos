// Example  This is an example .
package main

import (
	"encoding/json"
	"fmt"

	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/geoencoding/wkt"
	"github.com/spatial-go/geoos/planar"
	"github.com/spatial-go/geoos/space"
)

func main() {
	// First, choose the default algorithm.
	strategy := planar.NormalStrategy()
	// Secondly, manufacturing test data and convert it to geometry
	const polygon = `POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))`
	geometry, _ := wkt.UnmarshalString(polygon)
	// Last， call the Area () method and get result.
	area, e := strategy.Area(geometry)
	if e != nil {
		fmt.Print(e.Error())
	}
	fmt.Printf("polygon area:%f", area)
	// get result 4.0

	rawJSON := []byte(`
  { "type": "FeatureCollection",
	"features": [
	  { "type": "Feature",
		"geometry": {"type": "Point", "coordinates": [102.0, 0.5]},
		"properties": {"prop0": "value0"}
	  }
	]
  }`)

	fc := geojson.NewFeatureCollection()
	_ = json.Unmarshal(rawJSON, &fc)
	fmt.Printf("\n%p", fc)

	// Geometry will be unmarshalled into the correct geo.Geometry type.
	point := fc.Features[0].Geometry.Coordinates.(space.Point)
	fmt.Printf("\n%p", &point)

}
