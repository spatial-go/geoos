package decode

import (
	"github.com/spatial-go/geoos/geobuf/proto"
	"github.com/spatial-go/geoos/geojson"
)

func Decode(msg *proto.Data) interface{} {
	switch v := msg.DataType.(type) {
	case *proto.Data_Geometry_:
		geo := v.Geometry
		return DecodeGeometry(geo, msg.Precision, msg.Dimensions)
	case *proto.Data_Feature_:
		return DecodeFeature(msg, v.Feature)
	case *proto.Data_FeatureCollection_:
		collection := geojson.NewFeatureCollection()
		for _, feature := range v.FeatureCollection.Features {
			collection.Append(DecodeFeature(msg, feature))
		}
		return collection
	}
	return struct{}{}
}
