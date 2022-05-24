package geojson

import (
	"io"

	"github.com/spatial-go/geoos/space"
)

type GeojsonEncoder struct {
	BaseEncoder
}

// Encode Returns string of that encode geometry  by codeType.
func (e *GeojsonEncoder) Encode(g space.Geometry) []byte {
	gj := &Geometry{Coordinates: g}
	data, _ := gj.MarshalJSON()
	return data
}

// Decode Returns geometry of that decode string by codeType.
func (e *GeojsonEncoder) Decode(s []byte) (space.Geometry, error) {
	geom, err := UnmarshalGeometry(s)
	return geom.Geometry(), err
}

// Read Returns geometry from reader.
func (e *GeojsonEncoder) Read(r io.Reader) (space.Geometry, error) {
	if b, err := e.ReadBytes(r); err != nil {
		return nil, err
	} else {
		return e.Decode(b)
	}
}

// Write write geometry to reader.
func (e *GeojsonEncoder) Write(w io.Writer, g space.Geometry) error {
	b := e.Encode(g)
	return e.WriteBytes(w, b)
}

// WriteGeoJSON write geometry to writer  by codeType.
func (e *GeojsonEncoder) WriteGeoJSON(w io.Writer, g *FeatureCollection) error {
	if buf, err := g.MarshalJSON(); err != nil {
		return err
	} else {
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

// ReadGeoJSON Returns geometry from reader by codeType.
func (e *GeojsonEncoder) ReadGeoJSON(r io.Reader) (*FeatureCollection, error) {
	if b, err := e.ReadBytes(r); err != nil {
		return nil, err
	} else {
		return UnmarshalFeatureCollection(b)
	}
}
