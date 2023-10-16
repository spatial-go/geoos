package geojson

import (
	"io"

	"github.com/spatial-go/geoos/space"
)

// BaseEncoder defines base encoder.
type BaseEncoder struct {
}

// Encode Returns string of that encode geometry.
func (e *BaseEncoder) Encode(g space.Geometry) []byte {
	return []byte{}
}

// Decode Returns geometry of that decode string.
func (e *BaseEncoder) Decode(s []byte) (space.Geometry, error) {
	return nil, nil
}

// Read Returns geometry from reader.
func (e *BaseEncoder) Read(r io.Reader) (space.Geometry, error) {
	b, err := e.ReadBytes(r)
	if err != nil {
		return nil, err
	}
	return e.Decode(b)
}

// ReadBytes Returns geometry from reader.
func (e *BaseEncoder) ReadBytes(r io.Reader) ([]byte, error) {
	buf := []byte{}
	b := make([]byte, 4096)
	for {
		n, err := r.Read(b) 
        if err != nil && err != io.EOF {
            return nil,err
        }
        buf = append(buf, b[:n]...)
        if err == io.EOF {
            break
        }
	}
	return buf, nil
}

// Write write geometry to reader.
func (e *BaseEncoder) Write(w io.Writer, g space.Geometry) error {
	b := e.Encode(g)
	return e.WriteBytes(w, b)
}

// WriteBytes write geometry to writer.
func (e *BaseEncoder) WriteBytes(w io.Writer, buf []byte) error {
	if _, err := w.Write(buf); err != nil {
		return err
	}
	return nil
}

// WriteGeoJSON write geometry to writer .
func (e *BaseEncoder) WriteGeoJSON(w io.Writer, g *FeatureCollection) error {
	buf, err := g.MarshalJSON()
	if err != nil {
		return err
	}
	if _, err := w.Write(buf); err != nil {
		return err
	}
	return nil
}

// ReadGeoJSON Returns geometry from reader .
func (e *BaseEncoder) ReadGeoJSON(r io.Reader) (*FeatureCollection, error) {
	b, err := e.ReadBytes(r)
	if err != nil {
		return nil, err
	}
	return UnmarshalFeatureCollection(b)
}

// GeometryToFeatureCollection Returns feature collection from reader geometry .
func GeometryToFeatureCollection(geom space.Geometry) *FeatureCollection {
	fc := NewFeatureCollection()
	if geom == nil {
		return fc
	}
	switch geom.GeoJSONType() {
	case space.TypeCollection:
		features := []*Feature{}
		for _, v := range geom.Geom().(space.Collection) {
			geometry := NewGeometry(v)
			feature := NewFeature(*geometry)
			features = append(features, feature)
			fc.Features = features
		}
	default:
		features := []*Feature{}

		geometry := NewGeometry(geom.Geom())
		feature := NewFeature(*geometry)
		features = append(features, feature)
		fc.Features = features
	}
	return fc
}
