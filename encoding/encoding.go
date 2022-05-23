// Package encoding  is a library for encoding and decoding into Go structs using the geometries.
package encoding

import (
	"github.com/spatial-go/geoos/encoding/geobuf"
	"github.com/spatial-go/geoos/encoding/geocsv"
	"github.com/spatial-go/geoos/encoding/geojson"
	"github.com/spatial-go/geoos/encoding/wkb"
	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/space"
)

// encode type
const (
	WKT = iota
	WKB
	GeoJSON
	GeoCSV
	Geobuf
)

// Encoder defines encoder for encoding and decoding into Go structs using the geometries.
type Encoder interface {
	// Encode Returns string of that encode geometry  by codeType.
	Encode(g space.Geometry, codeType int) []byte
	// Decode Returns geometry of that decode string by codeType.
	Decode(s []byte, codeType int) (space.Geometry, error)
}

// Encode Returns string of that encode geometry  by codeType.
func Encode(g space.Geometry, codeType int) []byte {
	//TODO
	switch codeType {
	case WKT:
		return []byte(wkt.MarshalString(g))
	case WKB:
		s, _ := wkb.GeomToWKBHexStr(g)
		return []byte(s)
	case GeoJSON:
		g := &geojson.Geometry{Coordinates: g}
		data, _ := g.MarshalJSON()
		return data
	case GeoCSV:
		return geocsv.Encode(g)
	case Geobuf:
		return geobuf.Encode(g)
	default:
		return []byte{}
	}
}

// Decode Returns geometry of that decode string by codeType.
func Decode(s []byte, codeType int) (space.Geometry, error) {
	// TODO
	switch codeType {
	case WKT:
		return wkt.UnmarshalString(string(s))
	case WKB:
		return wkb.GeomFromWKBHexStr(string(s))
	case GeoJSON:
		geom, err := geojson.UnmarshalGeometry(s)
		return geom.Geometry(), err
	case GeoCSV:
		return geocsv.Decode(s)
	case Geobuf:
		return geobuf.Decode(s)
	default:
		return nil, nil
	}
}
