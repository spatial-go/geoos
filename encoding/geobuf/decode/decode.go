package decode

import (
	"github.com/spatial-go/geoos/encoding/geobuf/proto"
	"github.com/spatial-go/geoos/encoding/geojson"
)

// Decode ...
func Decode(msg *proto.Data) interface{} {
	switch v := msg.DataType.(type) {
	case *proto.Data_Geometry_:
		geo := v.Geometry
		return Geometry(geo, msg.Precision, msg.Dimensions)
	case *proto.Data_Feature_:
		return Feature(msg, v.Feature)
	case *proto.Data_FeatureCollection_:
		collection := geojson.NewFeatureCollection()
		for _, feature := range v.FeatureCollection.Features {
			collection.Append(Feature(msg, feature))
		}
		return collection
	}
	return struct{}{}
}
