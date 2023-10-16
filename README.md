# Geoos
Our organization `spatial-go` is officially established! The first open source project `Geoos`(Using `Golang`) provides spatial data and geometric algorithms.
All comments and suggestions are welcome!

## Guides

http://www.spatial-go.com

## Contents
  - [Guides](#guides)
  - [Contents](#contents)
  - [Structure](#structure)
  - [Documentation](#documentation)
  - [Maintainer](#maintainer)
  - [Contributing](#contributing)
  - [License](#license)



## Structure
1.algorithm
Package algorithm defines Specifies Computational Geometric and algorithm err.
2.clusters
Package clusters is a spatial clustering algorithm.
3.coordtransform
Package coordtransform is for transform coord.
4.example
Example This is an example .
5.geoencoding
Package geoencoding is a library for encoding and decoding into Go structs using the geometries.
6.grid
Package grid is used to generate grid data.
7.index
Package index define spatial index interface.
8.planar
Package planar provides support for the implementation of spatial operations and geometric algorithms.
9.space
Package space A representation of a linear vector geometry.
10.utils
Package utils A functions of utils.

## Documentation
How to use `Geoos`:
Example: Calculating `area` via `Geoos`
```
package main

import (
	"bytes"
	"fmt"

	"github.com/spatial-go/geoos/geoencoding"
	"github.com/spatial-go/geoos/planar"
)

func main() {
	// First, choose the default algorithm.
	strategy := planar.NormalStrategy()
	// Secondly, manufacturing test data and convert it to geometry
	const polygon = `POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))`
	// geometry, _ := wkt.UnmarshalString(polygon)

	buf0 := new(bytes.Buffer)
	buf0.Write([]byte(polygon))
	geometry, _ := geoencoding.Read(buf0, geoencoding.WKT)

	// Last， call the Area () method and get result.
	area, e := strategy.Area(geometry)
	if e != nil {
		fmt.Printf(e.Error())
	}
	fmt.Printf("%f", area)
	// get result 4.0
}
```
Example: geoencoding
[example_encoding.go](https://github.com/spatial-go/geoos/example/example_encoding.go)

## Maintainer

[@spatial-go](https://github.com/spatial-go)。

## Contributing

We will also uphold the concept of "openness, co-creation, and win-win" to contribute in the field of space computing.

Welcome to join us ！[please report an issue](https://github.com/spatial-go/geoos/issues/new)

Email Address： [geoos@changjing.ai](mailto:geoos@changjing.ai)

## License
`Geoos` is licensed under the:
[LGPL-2.1 ](LICENSE)