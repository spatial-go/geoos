package geom

/*
#include "../geo/geos.h"
*/
import "C"

//Geometry 可代表一下任意一个简单集合类型
//	Point
//	LineString
//	LinearRing
//	Polygon
//	MultiPoint
//	MultiLineString
//	MultiPolygon
//	GeometryCollection
type Geometry struct {
	Geometry *C.GEOSGeometry
}



