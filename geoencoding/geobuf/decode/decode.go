// Package encode is a library for decoding geobuf into Go structs using
package decode

import (
	"github.com/spatial-go/geoos/geoencoding/geobuf/protogeo"
	"github.com/spatial-go/geoos/geoencoding/geojson"
)

// Decode ...
func Decode(msg *protogeo.Data) interface{} {
	switch v := msg.DataType.(type) {
	case *protogeo.Data_Geometry_:
		geo := v.Geometry
		return Geometry(geo, msg.Precision, msg.Dimensions)
	case *protogeo.Data_Feature_:
		return Feature(msg, v.Feature)
	case *protogeo.Data_FeatureCollection_:
		collection := geojson.NewFeatureCollection()
		for _, feature := range v.FeatureCollection.Features {
			collection.Append(Feature(msg, feature))
		}
		return collection
	}
	return struct{}{}
}
