package encode

import (
	"github.com/spatial-go/geoos/encoding/geobuf/protogeo"
	"github.com/spatial-go/geoos/encoding/geojson"
)

// FeatureCollection ...
func FeatureCollection(g geojson.FeatureCollection, cfg *EncodingConfig) (*protogeo.Data_FeatureCollection, error) {
	features := make([]*protogeo.Data_Feature, len(g.Features))
	for i, feature := range g.Features {
		encoded, err := Feature(feature, cfg)
		if err != nil {
			return nil, err
		}
		features[i] = encoded
	}
	return &protogeo.Data_FeatureCollection{
		Features: features,
	}, nil
}
