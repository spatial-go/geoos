// Package geojson is a library for encoding and decoding GeoJSON into Go structs using
// the geometries. Supports both the json.Marshaler and json.Unmarshaler
// interfaces as well as helper functions such as `UnmarshalFeatureCollection` and `UnmarshalFeature`.
package geojson

import (
	"encoding/json"
	"fmt"

	"github.com/spatial-go/geoos/space"
)

const featureCollection = "FeatureCollection"

// A FeatureCollection correlates to a GeoJSON feature collection.
type FeatureCollection struct {
	Type     string     `json:"type"`
	BBox     BBox       `json:"bbox,omitempty"`
	Features []*Feature `json:"features"`
}

// NewFeatureCollection creates and initializes a new feature collection.
func NewFeatureCollection() *FeatureCollection {
	return &FeatureCollection{
		Type:     featureCollection,
		Features: []*Feature{},
	}
}

// Append appends a feature to the collection.
func (fc *FeatureCollection) Append(feature *Feature) *FeatureCollection {
	fc.Features = append(fc.Features, feature)
	return fc
}

// MarshalJSON converts the feature collection object into the proper JSON.
// It will handle the encoding of all the child features and geometries.
// Alternately one can call json.Marshal(fc) directly for the same result.
func (fc FeatureCollection) MarshalJSON() ([]byte, error) {
	type tempFC FeatureCollection

	c := tempFC{
		Type:     featureCollection,
		BBox:     fc.BBox,
		Features: fc.Features,
	}

	if c.Features == nil {
		c.Features = []*Feature{}
	}
	return json.Marshal(c)
}

// String returns string.
func (fc *FeatureCollection) String() string {
	if buf, err := fc.MarshalJSON(); err == nil {
		return string(buf)
	} else {
		return err.Error()
	}
}

// UnmarshalFeatureCollection decodes the data into a GeoJSON feature collection.
// Alternately one can call json.Unmarshal(fc) directly for the same result.
func UnmarshalFeatureCollection(data []byte) (*FeatureCollection, error) {
	fc := &FeatureCollection{}
	err := json.Unmarshal(data, fc)
	if err != nil {
		return nil, err
	}

	if fc.Type != featureCollection {
		return nil, fmt.Errorf("geojson: not a feature collection: type=%s", fc.Type)
	}
	for _, v := range fc.Features {
		if poly, ok := v.Geometry.Geometry().(space.Polygon); ok {
			for i, ring := range poly {
				if !space.Ring(ring).IsClosed() {
					poly[i] = append(ring, ring[0])
				}
			}
		} else if mult, ok := v.Geometry.Geometry().(space.MultiPolygon); ok {
			for _, poly := range mult {
				for i, ring := range poly {
					if !space.Ring(ring).IsClosed() {
						poly[i] = append(ring, ring[0])
					}
				}
			}
		}

		if !v.Geometry.Geometry().IsCorrect() {
			return nil, ErrInvalidGeometry
		}
	}

	return fc, nil
}
