// Package encode is a library for encoding geobuf into Go structs using
package encode

import (
	"github.com/spatial-go/geoos/geoencoding/geobuf/protogeo"
	"github.com/spatial-go/geoos/geoencoding/geojson"
)

// Encode ...
func Encode(obj interface{}) *protogeo.Data {
	data, err := WithOptions(obj, FromAnalysis(obj))
	if err != nil {
		panic(err)
	}
	return data
}

// WithOptions ...
func WithOptions(obj interface{}, opts ...EncodingOption) (*protogeo.Data, error) {
	cfg := &EncodingConfig{
		Dimension: 2,
		Precision: 100,
		Keys:      protogeo.NewKeyStore(),
	}
	for _, opt := range opts {
		opt(cfg)
	}

	data := &protogeo.Data{
		Keys:       cfg.Keys.Keys(),
		Dimensions: uint32(cfg.Dimension),
		Precision:  protogeo.EncodePrecision(cfg.Precision),
	}

	switch t := obj.(type) {
	case *geojson.FeatureCollection:
		collection, err := FeatureCollection(*t, cfg)
		if err != nil {
			return nil, err
		}
		data.DataType = &protogeo.Data_FeatureCollection_{
			FeatureCollection: collection,
		}
	case *geojson.Feature:
		feature, err := Feature(t, cfg)
		if err != nil {
			return nil, err
		}
		data.DataType = &protogeo.Data_Feature_{
			Feature: feature,
		}
	case *geojson.Geometry:
		data.DataType = &protogeo.Data_Geometry_{
			Geometry: Geometry(t, cfg),
		}
	}

	return data, nil
}
