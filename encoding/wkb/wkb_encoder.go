package wkb

import "github.com/spatial-go/geoos/space"

type WKBEncoder struct {
}

// Encode Returns string of that encode geometry  by codeType.
func (e *WKBEncoder) Encode(g space.Geometry) []byte {
	s, _ := GeomToWKBHexStr(g)
	return []byte(s)
}

// Decode Returns geometry of that decode string by codeType.
func (e *WKBEncoder) Decode(s []byte) (space.Geometry, error) {
	return GeomFromWKBHexStr(string(s))
}
