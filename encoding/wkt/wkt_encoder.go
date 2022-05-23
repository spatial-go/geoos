package wkt

import "github.com/spatial-go/geoos/space"

type WKTEncoder struct {
}

// Encode Returns string of that encode geometry  by codeType.
func (e *WKTEncoder) Encode(g space.Geometry) []byte {
	return []byte(MarshalString(g))
}

// Decode Returns geometry of that decode string by codeType.
func (e *WKTEncoder) Decode(s []byte) (space.Geometry, error) {
	return UnmarshalString(string(s))
}
