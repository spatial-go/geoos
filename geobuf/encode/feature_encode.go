package encode

import (
	"github.com/spatial-go/geoos/geobuf/proto"
	"github.com/spatial-go/geoos/geobuf/utils"
	"github.com/spatial-go/geoos/geojson"
)

// FeatureEncode ...
type FeatureEncode struct{}

// Feature ...
func Feature(feature *geojson.Feature, cfg *EncodingConfig) (protoFeature *proto.Data_Feature, err error) {
	geo := Geometry(&feature.Geometry, cfg)
	if err != nil {
		return
	}
	id, err := utils.EncodeID(feature.ID)
	protoFeature = &proto.Data_Feature{
		Geometry: geo,
	}
	if err == nil {
		protoFeature.IdType = id
	} else {
		newID, newErr := utils.EncodeID(feature.ID)
		if newErr != nil {
			return nil, newErr
		}
		protoFeature.IdType = newID
	}
	properties := make([]uint32, 0, 2*len(feature.Properties))
	values := make([]*proto.Data_Value, 0, len(feature.Properties))
	for key, val := range feature.Properties {
		encoded, err2 := utils.EncodeValue(val)
		if err2 != nil {
			return
		}
		idx := cfg.Keys.IndexOf(key)
		values = append(values, encoded)
		properties = append(properties, uint32(idx))
		properties = append(properties, uint32(len(values)-1))
	}
	protoFeature.Values = values
	protoFeature.Properties = properties
	return
}
