# Geoos
Our organization `spatial-go` is officially established! The first open source project `Geoos`(Using `Golang`) provides spatial data and geometric algorithms.
All comments and suggestions are welcome!

## Guides

http://www.spatial-go.com

## Contents

- [Installation](#Installation)
- [Structure](#Structure)
- [Documentation](#Documentation)
- [Maintainer](#Maintainer)
- [Contributing](#Contributing)
- [Copying](#Copying)


## Installation

The project depends on [geos](https://github.com/libgeos/geos) (GEOS is a C++ port of the ​JTS Topology Suite), you need to complete the installation of `geos` first. The installation of `geos`:

1. Mac OS X(via brew)
```sh
$ brew install geos
```
2. Ubuntu or Debian
```sh
$ apt-get install libgeos-dev
```
3. Build from source code
```sh
$ wget http://download.osgeo.org/geos/geos-3.9.0.tar.bz2
$ tar xvfj geos-3.9.0.tar.bz2
$ cd geos-3.9.0
$ ./configure
$ make
$ sudo make install
```

## Structure
1. `Geo` package contains references and calls to the `GEOS C` library to implement spatial operations.
2. `Algorithm` is the definition of spatial operation, which is outside exposing.
3. `strategy.go` defines the implementation of the spatial computing based algorithm.

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

Thanks to [OSGeo](https://www.osgeo.org/), we will also uphold the concept of "openness, co-creation, and win-win" to contribute in the field of space computing.

We use GitHub issues, but that is not a requirement. Contact us via email is also fine. Please use the GitHub issues only for actual issues. If you are not 100% sure that your problem is an issue, please first discuss this via email. If you report an issue, please describe exactly how to reproduce it.

Email Address： [geoos@changjing.ai](geoos@changjing.ai)

## Copying
`Geoos` is licensed under the:
[LGPL-2.1 ](LICENSE)
