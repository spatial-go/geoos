package decode

import (
	"github.com/spatial-go/geoos/geoencoding/geobuf/protogeo"
	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

// Feature ...
func Feature(data *protogeo.Data, feature *protogeo.Data_Feature) *geojson.Feature {
	geo := feature.Geometry
	decodedGeo := Geometry(geo, data.Precision, data.Dimensions)
	var geoFeature *geojson.Feature
	switch decodedGeo.Type {
	case space.TypeCollection:
		collection := make(space.Collection, len(decodedGeo.Geometries))
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
		valIdx := feature.Properties[i+1]
		val := feature.Values[valIdx]
		switch actualVal := val.ValueType.(type) {
		case *protogeo.Data_Value_BoolValue:
			geoFeature.Properties[data.Keys[feature.Properties[i]]] = actualVal.BoolValue
		case *protogeo.Data_Value_DoubleValue:
			geoFeature.Properties[data.Keys[feature.Properties[i]]] = actualVal.DoubleValue
		case *protogeo.Data_Value_StringValue:
			geoFeature.Properties[data.Keys[feature.Properties[i]]] = actualVal.StringValue
		case *protogeo.Data_Value_PosIntValue:
			geoFeature.Properties[data.Keys[feature.Properties[i]]] = uint(actualVal.PosIntValue)
		case *protogeo.Data_Value_NegIntValue:
			geoFeature.Properties[data.Keys[feature.Properties[i]]] = int(actualVal.NegIntValue) * -1
		case *protogeo.Data_Value_JsonValue:
			geoFeature.Properties[data.Keys[feature.Properties[i]]] = actualVal.JsonValue
		}
	}
	switch id := feature.IdType.(type) {
	case *protogeo.Data_Feature_Id:
		geoFeature.ID = id.Id
	case *protogeo.Data_Feature_IntId:
		geoFeature.ID = id.IntId
	}
	return geoFeature
}
