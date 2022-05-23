// Package geobuf  is a library for encoding and decoding geobuf into Go structs using the geometries.
package geobuf

import (
	"github.com/spatial-go/geoos/encoding/geobuf/decode"
	"github.com/spatial-go/geoos/encoding/geobuf/encode"
	"github.com/spatial-go/geoos/space"
)

// Encode Returns string of that encode geometry.
func Encode(g space.Geometry) []byte {
	//TODO
	return []byte(encode.Encode(g).String())
}

// Encode Returns geometry of that decode string by codeType.
func Decode(s []byte) (space.Geometry, error) {
	//TODO
	geom := decode.Decode(encode.Encode(s))
	return geom.(space.Geometry), nil
}
