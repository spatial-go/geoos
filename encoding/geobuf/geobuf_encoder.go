package geobuf

import (
	"fmt"
	"io"

	"github.com/spatial-go/geoos/encoding/geobuf/encode"
	"github.com/spatial-go/geoos/encoding/geojson"
	"github.com/spatial-go/geoos/space"
)

type GeobufEncoder struct {
	geojson.BaseEncoder
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

// Read Returns geometry from reader.
func (e *GeobufEncoder) Read(r io.Reader) (space.Geometry, error) {
	if b, err := e.ReadBytes(r); err != nil {
		return nil, err
	} else {
		return e.Decode(b)
	}
}

// Write write geometry to reader.
func (e *GeobufEncoder) Write(w io.Writer, g space.Geometry) error {
	b := e.Encode(g)
	return e.WriteBytes(w, b)
}

// WriteGeoJSON write geometry to writer.
func (e *GeobufEncoder) WriteGeoJSON(w io.Writer, g *geojson.FeatureCollection) error {
	colls := space.Collection{}
	for _, v := range g.Features {
		colls = append(colls, v.Geometry.Geometry())
	}
	return e.Write(w, colls)
}

// ReadGeoJSON Returns geometry from reader .
func (e *GeobufEncoder) ReadGeoJSON(r io.Reader) (*geojson.FeatureCollection, error) {
	if geom, err := e.Read(r); err != nil {
		return nil, err
	} else {
		return geojson.GeometryToFeatureCollection(geom), nil
	}
}
