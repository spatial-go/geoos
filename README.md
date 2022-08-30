# Geoos
Our organization `spatial-go` is officially established! The first open source project `Geoos`(Using `Golang`) provides spatial data and geometric algorithms.
All comments and suggestions are welcome!

## Guides

http://www.spatial-go.com

## Contents

- [Geoos](#geoos)
  - [Guides](#guides)
  - [Contents](#contents)
  - [Structure](#structure)
  - [Documentation](#documentation)
  - [Maintainer](#maintainer)
  - [Contributing](#contributing)
  - [License](#license)



## Structure
1. `Algorithm` is the definition of spatial operation, which is outside exposing.
2. `strategy.go` defines the implementation of the spatial computing based algorithm.

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