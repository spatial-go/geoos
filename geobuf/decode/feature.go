package decode

import (
	"github.com/spatial-go/geoos/geobuf/proto"
	"github.com/spatial-go/geoos/geojson"
	geoos "github.com/spatial-go/geoos/space"
)

// Feature ...
func Feature(data *proto.Data, feature *proto.Data_Feature) *geojson.Feature {
	geo := feature.Geometry
	decodedGeo := Geometry(geo, data.Precision, data.Dimensions)
	var geoFeature *geojson.Feature
	switch decodedGeo.Type {
	case geoos.TypeCollection:
		collection := make(geoos.Collection, len(decodedGeo.Geometries))
		for i, child := range decodedGeo.Geometries {
			collection[i] = child.Coordinates
		}
		geoFeature = geojson.NewFeature(*geojson.NewGeometry(collection))
	default:
		geoFeature = geojson.NewFeature(geojson.Geometry{
			Coordinates: decodedGeo.Coordinates,
		})
	}

	for i := 0; i < len(feature.Properties); i = i + 2 {
		keyIdx := feature.Properties[i]
		valIdx := feature.Properties[i+1]
		val := feature.Values[valIdx]
		switch actualVal := val.ValueType.(type) {
		case *proto.Data_Value_BoolValue:
			geoFeature.Properties[data.Keys[keyIdx]] = actualVal.BoolValue
		case *proto.Data_Value_DoubleValue:
			geoFeature.Properties[data.Keys[keyIdx]] = actualVal.DoubleValue
		case *proto.Data_Value_StringValue:
			geoFeature.Properties[data.Keys[keyIdx]] = actualVal.StringValue
		case *proto.Data_Value_PosIntValue:
			geoFeature.Properties[data.Keys[keyIdx]] = uint(actualVal.PosIntValue)
		case *proto.Data_Value_NegIntValue:
			geoFeature.Properties[data.Keys[keyIdx]] = int(actualVal.NegIntValue) * -1
		case *proto.Data_Value_JsonValue:
			geoFeature.Properties[data.Keys[keyIdx]] = actualVal.JsonValue
		}
	}
	switch id := feature.IdType.(type) {
	case *proto.Data_Feature_Id:
		geoFeature.ID = id.Id
	case *proto.Data_Feature_IntId:
		geoFeature.ID = id.IntId
	}
	return geoFeature
}
