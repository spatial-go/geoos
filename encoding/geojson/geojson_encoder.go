package geojson

import "github.com/spatial-go/geoos/space"

type GeojsonEncoder struct {
}

// Encode Returns string of that encode geometry  by codeType.
func (e *GeojsonEncoder) Encode(g space.Geometry) []byte {
	gj := &Geometry{Coordinates: g}
	data, _ := gj.MarshalJSON()
	return data
}

// Decode Returns geometry of that decode string by codeType.
func (e *GeojsonEncoder) Decode(s []byte) (space.Geometry, error) {
	geom, err := UnmarshalGeometry(s)
	return geom.Geometry(), err
}
