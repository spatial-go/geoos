package wkb

import (
	"io"

	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

type WKBEncoder struct {
	geojson.BaseEncoder
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

// Read Returns geometry from reader.
func (e *WKBEncoder) Read(r io.Reader) (space.Geometry, error) {
	if b, err := e.ReadBytes(r); err != nil {
		return nil, err
	} else {
		return e.Decode(b)
	}
}

// Write write geometry to reader.
func (e *WKBEncoder) Write(w io.Writer, g space.Geometry) error {
	b := e.Encode(g)
	return e.WriteBytes(w, b)
}

// WriteGeoJSON write geometry to writer.
func (e *WKBEncoder) WriteGeoJSON(w io.Writer, g *geojson.FeatureCollection) error {
	colls := space.Collection{}
	for _, v := range g.Features {
		colls = append(colls, v.Geometry.Geometry())
	}
	return e.Write(w, colls)
}

// ReadGeoJSON Returns geometry from reader .
func (e *WKBEncoder) ReadGeoJSON(r io.Reader) (*geojson.FeatureCollection, error) {
	if geom, err := e.Read(r); err != nil {
		return nil, err
	} else {
		return geojson.GeometryToFeatureCollection(geom), nil
	}
}
