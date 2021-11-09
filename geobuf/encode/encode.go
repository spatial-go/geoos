package encode

import (
	"github.com/spatial-go/geoos/geobuf/proto"
	"github.com/spatial-go/geoos/geobuf/utils"
	"github.com/spatial-go/geoos/geojson"
)

func Encode(obj interface{}) *proto.Data {
	data, err := EncodeWithOptions(obj, FromAnalysis(obj))
	if err != nil {
		panic(err)
	}
	return data
}

func EncodeWithOptions(obj interface{}, opts ...EncodingOption) (*proto.Data, error) {
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
		collection, err := EncodeFeatureCollection(*t, cfg)
		if err != nil {
			return nil, err
		}
		data.DataType = &proto.Data_FeatureCollection_{
			FeatureCollection: collection,
		}
	case *geojson.Feature:
		feature, err := EncodeFeature(t, cfg)
		if err != nil {
			return nil, err
		}
		data.DataType = &proto.Data_Feature_{
			Feature: feature,
		}
	case *geojson.Geometry:
		data.DataType = &proto.Data_Geometry_{
			Geometry: EncodeGeometry(t, cfg),
		}
	}

	return data, nil
}
