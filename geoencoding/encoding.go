// Package geoencoding  is a library for encoding and decoding into Go structs using the geometries.
package geoencoding

import (
	"io"

	"github.com/spatial-go/geoos/geoencoding/geobuf"
	"github.com/spatial-go/geoos/geoencoding/geocsv"
	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/geoencoding/wkb"
	"github.com/spatial-go/geoos/geoencoding/wkt"
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
	// Encode Returns string of that encode geometry.
	Encode(g space.Geometry) []byte
	// Decode Returns geometry of that decode string.
	Decode(s []byte) (space.Geometry, error)

	// Read Returns geometry from reader.
	Read(r io.Reader) (space.Geometry, error)

	// Write write geometry to writer.
	Write(w io.Writer, g space.Geometry) error

	// Read Returns geometry from reader.
	ReadGeoJSON(r io.Reader) (*geojson.FeatureCollection, error)

	// Write write geometry to writer.
	WriteGeoJSON(w io.Writer, g *geojson.FeatureCollection) error
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

// Write write geometry to writer.  by codeType.
func Write(w io.Writer, g space.Geometry, codeType int) error {
	encode := getEncoder(codeType)
	return encode.Write(w, g)
}

// Read Returns geometry from reader by codeType.
func Read(r io.Reader, codeType int) (space.Geometry, error) {
	encode := getEncoder(codeType)
	return encode.Read(r)
}

// WriteGeoJSON write geometry to writer  by codeType.
func WriteGeoJSON(w io.Writer, g *geojson.FeatureCollection, codeType int) error {
	encode := getEncoder(codeType)
	return encode.WriteGeoJSON(w, g)
}

// ReadGeoJSON Returns geometry from reader by codeType.
func ReadGeoJSON(r io.Reader, codeType int) (*geojson.FeatureCollection, error) {
	encode := getEncoder(codeType)
	return encode.ReadGeoJSON(r)
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
		encode = &geojson.BaseEncoder{}
	}
	return encode
}
