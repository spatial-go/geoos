package encode

import (
	"github.com/spatial-go/geoos/geoencoding/geobuf/protogeo"
	"github.com/spatial-go/geoos/geoencoding/geojson"
)

// FeatureEncode ...
type FeatureEncode struct{}

// Feature ...
func Feature(feature *geojson.Feature, cfg *EncodingConfig) (protoFeature *protogeo.Data_Feature, err error) {
	geo := Geometry(&feature.Geometry, cfg)
	if err != nil {
		return
	}
	id, err := protogeo.EncodeID(feature.ID)
	protoFeature = &protogeo.Data_Feature{
		Geometry: geo,
	}
	if err == nil {
		protoFeature.IdType = id
	} else {
		newID, newErr := protogeo.EncodeID(feature.ID)
		if newErr != nil {
			return nil, newErr
		}
		protoFeature.IdType = newID
	}
	properties := make([]uint32, 0, 2*len(feature.Properties))
	values := make([]*protogeo.Data_Value, 0, len(feature.Properties))
	for key, val := range feature.Properties {
		encoded, err2 := protogeo.EncodeValue(val)
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
