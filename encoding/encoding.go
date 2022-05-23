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
	Encode(g space.Geometry) []byte
	// Decode Returns geometry of that decode string by codeType.
	Decode(s []byte) (space.Geometry, error)
}

type BaseEncode struct {
}

// Encode Returns string of that encode geometry  by codeType.
func (e *BaseEncode) Encode(g space.Geometry) []byte {
	return []byte{}
}

// Decode Returns geometry of that decode string by codeType.
func (e *BaseEncode) Decode(s []byte) (space.Geometry, error) {
	return nil, nil
}

// Encode Returns string of that encode geometry  by codeType.
func Encode(g space.Geometry, codeType int) []byte {
	encode := getEncoder(codeType)
	return encode.Encode(g)
}

// Decode Returns geometry of that decode string by codeType.
func Decode(s []byte, codeType int) (space.Geometry, error) {
	encode := getEncoder(codeType)
	return encode.Decode(s)
}

func getEncoder(codeType int) Encoder {
	var encode Encoder
	switch codeType {
	case WKT:
		encode = &wkt.WKTEncoder{}
	case WKB:
		encode = &wkb.WKBEncoder{}
	case GeoJSON:
		encode = &geojson.GeojsonEncoder{}
	case GeoCSV:
		encode = &geocsv.GeocsvEncoder{}
	case Geobuf:
		encode = &geobuf.GeobufEncoder{}
	default:
		encode = &BaseEncode{}
	}
	return encode
}
