// Package geobuf is a library for encoding and decoding geobuf into Go structs using
package geobuf

import (
	"io"

	"github.com/spatial-go/geoos/geoencoding/geobuf/decode"
	"github.com/spatial-go/geoos/geoencoding/geobuf/encode"
	"github.com/spatial-go/geoos/geoencoding/geobuf/protogeo"
	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
	"google.golang.org/protobuf/proto"
)

type GeobufEncoder struct {
	geojson.BaseEncoder
}

// Encode Returns string of that encode geometry  by codeType.
func (e *GeobufEncoder) Encode(g space.Geometry) []byte {
	gj := geojson.NewGeometry(g)
	protoGeo := encode.Encode(gj)

	b, _ := proto.Marshal(protoGeo)
	return b
}

// Decode Returns geometry of that decode string by codeType.
func (e *GeobufEncoder) Decode(s []byte) (space.Geometry, error) {
	protoGeo := &protogeo.Data{}
	_ = proto.Unmarshal(s, protoGeo)
	geom := decode.Decode(protoGeo)
	switch gj := geom.(type) {
	case *geojson.FeatureCollection:
		colls := space.Collection{}
		for _, v := range gj.Features {
			colls = append(colls, v.Geometry.Geometry())
		}
		return colls, nil
	case *geojson.Feature:
		return gj.Geometry.Geometry(), nil
	case *geojson.Geometry:
		return gj.Geometry(), nil
	}
	return nil, nil
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

	protoGeo := encode.Encode(g)

	b, _ := proto.Marshal(protoGeo)
	return e.WriteBytes(w, b)
}

// ReadGeoJSON Returns geometry from reader .
func (e *GeobufEncoder) ReadGeoJSON(r io.Reader) (*geojson.FeatureCollection, error) {
	if b, err := e.ReadBytes(r); err != nil {
		return nil, err
	} else {
		protoGeo := &protogeo.Data{}
		_ = proto.Unmarshal(b, protoGeo)
		geom := decode.Decode(protoGeo)
		switch gj := geom.(type) {
		case *geojson.FeatureCollection:
			return gj, nil
		case *geojson.Feature:
			fc := geojson.NewFeatureCollection()
			features := []*geojson.Feature{}
			features = append(features, gj)
			fc.Features = features
			return fc, nil
		case *geojson.Geometry:
			fc := geojson.NewFeatureCollection()
			features := []*geojson.Feature{}
			feature := geojson.NewFeature(*gj)
			features = append(features, feature)
			fc.Features = features
			return fc, nil
		}
		return nil, nil
	}
}
