package encode

import (
	"github.com/spatial-go/geoos/encoding/geobuf/proto"
	"github.com/spatial-go/geoos/encoding/geojson"
)

// FeatureCollection ...
func FeatureCollection(g geojson.FeatureCollection, cfg *EncodingConfig) (*proto.Data_FeatureCollection, error) {
	features := make([]*proto.Data_Feature, len(g.Features))
	for i, feature := range g.Features {
		encoded, err := Feature(feature, cfg)
		if err != nil {
			return nil, err
		}
		features[i] = encoded
	}
	return &proto.Data_FeatureCollection{
		Features: features,
	}, nil
}
