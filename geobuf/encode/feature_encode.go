package encode

import (
	"github.com/spatial-go/geoos/geobuf/proto"
	"github.com/spatial-go/geoos/geobuf/utils"
	"github.com/spatial-go/geoos/geojson"
)

type FeatureEncode struct{}

func EncodeFeature(feature *geojson.Feature, cfg *EncodingConfig) (protoFeature *proto.Data_Feature, err error) {
	geo := EncodeGeometry(&feature.Geometry, cfg)
	if err != nil {
		return
	}
	id, err := utils.EncodeId(feature.ID)
	protoFeature = &proto.Data_Feature{
		Geometry: geo,
	}
	if err == nil {
		protoFeature.IdType = id
	} else {
		newId, newErr := utils.EncodeId(feature.ID)
		if newErr != nil {
			return nil, newErr
		}
		protoFeature.IdType = newId
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
