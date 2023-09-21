package geojson

import (
	"io"
	"log"
	"strings"

	"github.com/spatial-go/geoos/space"
)

// Encoder defines geojson encoder.
type Encoder struct {
	BaseEncoder
}

// Encode Returns string of that encode geometry  by codeType.
func (e *Encoder) Encode(g space.Geometry) []byte {
	gj := &Geometry{Coordinates: g}
	data, _ := gj.MarshalJSON()
	return data
}

// Decode Returns geometry of that decode string by codeType.
func (e *Encoder) Decode(s []byte) (space.Geometry, error) {
	if strings.Contains(string(s), "\"type\":\"FeatureCollection\"") {
		colls, err := UnmarshalFeatureCollection(s)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		geom := space.Collection{}
		for _, v := range colls.Features {
			geom = append(geom, v.Geometry.Geometry())
		}
		return geom, nil
		} else if strings.Contains(string(s), "\"type\":\"Feature\"") {
		feat, err := UnmarshalFeature(s)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		geom := feat.Geometry.Geometry()
		return geom, nil
	}
	geom, err := UnmarshalGeometry(s)
	if err != nil {
		return nil, err
	}
	return geom.Geometry(), err
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

// WriteGeoJSON write geometry to writer  by codeType.
func (e *Encoder) WriteGeoJSON(w io.Writer, g *FeatureCollection) error {
	buf, err := g.MarshalJSON()
	if err != nil {
		return err
	}
	if _, err := w.Write(buf); err != nil {
		return err
	}
	return nil
}

// ReadGeoJSON Returns geometry from reader by codeType.
func (e *Encoder) ReadGeoJSON(r io.Reader) (*FeatureCollection, error) {
	b, err := e.ReadBytes(r)
	if err != nil {
		return nil, err
	}
	return UnmarshalFeatureCollection(b)
}
