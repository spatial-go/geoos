package wkt

import (
	"io"

	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

type WKTEncoder struct {
	geojson.BaseEncoder
}

// Encode Returns string of that encode geometry  by codeType.
func (e *WKTEncoder) Encode(g space.Geometry) []byte {
	return []byte(MarshalString(g))
}

// Decode Returns geometry of that decode string by codeType.
func (e *WKTEncoder) Decode(s []byte) (space.Geometry, error) {
	return UnmarshalString(string(s))
}

// Read Returns geometry from reader.
func (e *WKTEncoder) Read(r io.Reader) (space.Geometry, error) {
	if b, err := e.ReadBytes(r); err != nil {
		return nil, err
	} else {
		return e.Decode(b)
	}
}

// Write write geometry to reader.
func (e *WKTEncoder) Write(w io.Writer, g space.Geometry) error {
	b := e.Encode(g)
	return e.WriteBytes(w, b)
}

// WriteGeoJSON write geometry to writer.
func (e *WKTEncoder) WriteGeoJSON(w io.Writer, g *geojson.FeatureCollection) error {
	colls := space.Collection{}
	for _, v := range g.Features {
		colls = append(colls, v.Geometry.Geometry())
	}
	return e.Write(w, colls)
}

// ReadGeoJSON Returns geometry from reader .
func (e *WKTEncoder) ReadGeoJSON(r io.Reader) (*geojson.FeatureCollection, error) {
	if geom, err := e.Read(r); err != nil {
		return nil, err
	} else {
		return geojson.GeometryToFeatureCollection(geom), nil
	}
}
