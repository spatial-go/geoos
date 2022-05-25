package geojson

import (
	"encoding/json"
	"fmt"
)

// A Feature corresponds to GeoJSON feature object
type Feature struct {
	ID         interface{} `json:"id,omitempty"`
	Type       string      `json:"type"`
	BBox       BBox        `json:"bbox,omitempty"`
	Geometry   Geometry    `json:"geometry"`
	Properties Properties  `json:"properties"`
}

// NewFeature creates and initializes a GeoJSON feature given the required attributes.
func NewFeature(geometry Geometry) *Feature {
	return &Feature{
		Type:       "Feature",
		Geometry:   geometry,
		Properties: make(map[string]interface{}),
	}
}

// MarshalJSON converts the feature object into the proper JSON.
// It will handle the encoding of all the child geometries.
// Alternately one can call json.Marshal(f) directly for the same result.
func (f Feature) MarshalJSON() ([]byte, error) {
	jf := &jsonFeature{
		ID:         f.ID,
		Type:       "Feature",
		Properties: f.Properties,
		BBox:       f.BBox,
		Geometry:   NewGeometry(f.Geometry.Geometry()),
	}

	if len(jf.Properties) == 0 {
		jf.Properties = nil
	}

	return json.Marshal(jf)
}

// UnmarshalFeature decodes the data into a GeoJSON feature.
// Alternately one can call json.Unmarshal(f) directly for the same result.
func UnmarshalFeature(data []byte) (*Feature, error) {
	f := &Feature{}
	err := json.Unmarshal(data, f)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// UnmarshalJSON handles the correct unmarshalling of the data
// into the geoos.Geometry types.
func (f *Feature) UnmarshalJSON(data []byte) error {
	jf := &jsonFeature{}
	err := json.Unmarshal(data, &jf)
	if err != nil {
		return err
	}

	if jf.Type != "Feature" {
		return fmt.Errorf("geojson: not a feature: type=%s", jf.Type)
	}

	if jf.Geometry == nil || (jf.Geometry.Coordinates == nil && jf.Geometry.Geometries == nil) {
		return ErrInvalidGeometry
	}

	*f = Feature{
		ID:         jf.ID,
		Type:       jf.Type,
		Properties: jf.Properties,
		BBox:       jf.BBox,
		Geometry:   *jf.Geometry,
	}

	return nil
}

type jsonFeature struct {
	ID         interface{} `json:"id,omitempty"`
	Type       string      `json:"type"`
	BBox       BBox        `json:"bbox,omitempty"`
	Geometry   *Geometry   `json:"geometry"`
	Properties Properties  `json:"properties"`
}
