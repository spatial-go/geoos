package geobuf

import (
	"fmt"

	"github.com/spatial-go/geoos/encoding/geobuf/encode"
	"github.com/spatial-go/geoos/encoding/geojson"
	"github.com/spatial-go/geoos/space"
)

type GeobufEncoder struct {
}

// Encode Returns string of that encode geometry  by codeType.
func (e *GeobufEncoder) Encode(g space.Geometry) []byte {
	//TODO
	gj := &geojson.Geometry{Coordinates: g}

	return []byte(fmt.Sprintf("%v", encode.Encode(gj).String()))
}

// Decode Returns geometry of that decode string by codeType.
func (e *GeobufEncoder) Decode(s []byte) (space.Geometry, error) {
	//TODO
	geom, err := geojson.UnmarshalGeometry(s)
	return geom.Geometry(), err
}
