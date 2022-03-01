package encode

import (
	"github.com/spatial-go/geoos/geobuf/proto"
	"github.com/spatial-go/geoos/geobuf/utils"
	"github.com/spatial-go/geoos/geojson"
)

// Encode ...
func Encode(obj interface{}) *proto.Data {
	data, err := WithOptions(obj, FromAnalysis(obj))
	if err != nil {
		panic(err)
	}
	return data
}

// WithOptions ...
func WithOptions(obj interface{}, opts ...EncodingOption) (*proto.Data, error) {
	cfg := &EncodingConfig{
		Dimension: 2,
		Precision: 100,
		Keys:      utils.NewKeyStore(),
	}
	for _, opt := range opts {
		opt(cfg)
	}

	data := &proto.Data{
		Keys:       cfg.Keys.Keys(),
		Dimensions: uint32(cfg.Dimension),
		Precision:  utils.EncodePrecision(cfg.Dimension),
	}

	switch t := obj.(type) {
	case *geojson.FeatureCollection:
		collection, err := FeatureCollection(*t, cfg)
		if err != nil {
			return nil, err
		}
		data.DataType = &proto.Data_FeatureCollection_{
			FeatureCollection: collection,
		}
	case *geojson.Feature:
		feature, err := Feature(t, cfg)
		if err != nil {
			return nil, err
		}
		data.DataType = &proto.Data_Feature_{
			Feature: feature,
		}
	case *geojson.Geometry:
		data.DataType = &proto.Data_Geometry_{
			Geometry: Geometry(t, cfg),
		}
	}

	return data, nil
}
