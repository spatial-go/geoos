# Geoos
Our organization `spatial-go` is officially established! The first open source project `Geoos`(Using `Golang`) provides spatial data and geometric algorithms.
All comments and suggestions are welcome!

## Guides

http://www.spatial-go.com

## Contents

- [Structure](#Structure)
- [Documentation](#Documentation)
- [Maintainer](#Maintainer)
- [Contributing](#Contributing)
- [License](#License)



## Structure
1. `Algorithm` is the definition of spatial operation, which is outside exposing.
2. `strategy.go` defines the implementation of the spatial computing based algorithm.

## Documentation
How to use `Geoos`:
Example: Calculating `area` via `Geoos`
```
package main

import (
  "encoding/json"
  "fmt"
  "github.com/spatial-go/geoos"
  "github.com/spatial-go/geoos/encoding/wkt"
  "github.com/spatial-go/geoos/geojson"
  "github.com/spatial-go/geoos/planar"
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
    fmt.Printf(e.Error())
  }
  fmt.Printf("%f", area)
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
  println("%p", fc)

  // Geometry will be unmarshalled into the correct geo.Geometry type.
  point := fc.Features[0].Geometry.(geoos.Point)
  println("%p", &point)

}

```

## Maintainer

[@spatial-go](https://github.com/spatial-go)。

## Contributing

We will also uphold the concept of "openness, co-creation, and win-win" to contribute in the field of space computing.

Welcome to join us ！[please report an issue](https://github.com/spatial-go/geos/issues/new)

Email Address： [geoos@changjing.ai](geoos@changjing.ai)

## License
`Geoos` is licensed under the:
[LGPL-2.1 ](LICENSE)