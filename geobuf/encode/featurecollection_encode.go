package encode

import (
	"github.com/spatial-go/geoos/geobuf/proto"
	"github.com/spatial-go/geoos/geojson"
)

func EncodeFeatureCollection(g geojson.FeatureCollection, cfg *EncodingConfig) (*proto.Data_FeatureCollection, error) {
	features := make([]*proto.Data_Feature, len(g.Features))
	for i, feature := range g.Features {
		encoded, err := EncodeFeature(feature, cfg)
		if err != nil {
			return nil, err
		}
		features[i] = encoded
	}
	return &proto.Data_FeatureCollection{
		Features: features,
	}, nil
}
