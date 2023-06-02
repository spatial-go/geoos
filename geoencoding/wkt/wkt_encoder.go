package wkt

import (
	"io"

	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

// Encoder defines wkt encoder.
type Encoder struct {
	geojson.BaseEncoder
}

// Encode Returns string of that encode geometry  by codeType.
func (e *Encoder) Encode(g space.Geometry) []byte {
	return []byte(MarshalString(g))
}

// Decode Returns geometry of that decode string by codeType.
func (e *Encoder) Decode(s []byte) (space.Geometry, error) {
	return UnmarshalString(string(s))
}

// Read Returns geometry from reader.
func (e *Encoder) Read(r io.Reader) (space.Geometry, error) {
	b, err := e.ReadBytes(r)
	if err != nil {
		return nil, err
	}
	return e.Decode(b)
}

// Write write geometry to reader.
func (e *Encoder) Write(w io.Writer, g space.Geometry) error {
	b := e.Encode(g)
	return e.WriteBytes(w, b)
}

// WriteGeoJSON write geometry to writer.
func (e *Encoder) WriteGeoJSON(w io.Writer, g *geojson.FeatureCollection) error {
	colls := space.Collection{}
	for _, v := range g.Features {
		colls = append(colls, v.Geometry.Geometry())
	}
	return e.Write(w, colls)
}

// ReadGeoJSON Returns geometry from reader .
func (e *Encoder) ReadGeoJSON(r io.Reader) (*geojson.FeatureCollection, error) {
	geom, err := e.Read(r)
	if err != nil {
		return nil, err
	}
	return geojson.GeometryToFeatureCollection(geom), nil
}
