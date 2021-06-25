package geoos

// MultiPolygon is a set of polygons.
type MultiPolygon []Polygon

// GeoJSONType returns the GeoJSON type for the object.
func (mp MultiPolygon) GeoJSONType() string {
	return TypeMultiPolygon
}

// Dimensions returns 0 because a MultiPoint is a 0d object.
func (mp MultiPolygon) Dimensions() int {
	return 2
}

// Nums num of multiPolygons
func (mp MultiPolygon) Nums() int {
	return len(mp)
}

// // Area Returns the area of this polygonal geometry
// func (mp MultiPolygon) Area() (float64, error) {
// 	s := NormalStrategy()
// 	return s.Area(mp)
// }

// Bound returns a bound around the multi-polygon.
func (mp MultiPolygon) Bound() Bound {
	if len(mp) == 0 {
		return emptyBound
	}
	bound := mp[0].Bound()
	for i := 1; i < len(mp); i++ {
		bound = bound.Union(mp[i].Bound())
	}

	return bound
}

// Equal compares two multi-polygons.
func (mp MultiPolygon) Equal(multiPolygon MultiPolygon) bool {
	if len(mp) != len(multiPolygon) {
		return false
	}

	for i, p := range mp {
		if !p.Equal(multiPolygon[i]) {
			return false
		}
	}

	return true
}

// Area returns the area of a polygonal geometry.
func (mp MultiPolygon) Area() (float64, error) {
	area := 0.0
	for _, polygon := range mp {
		if areaOfPolygon, err := polygon.Area(); err == nil {
			area += areaOfPolygon
		} else {
			return 0, nil
		}
	}
	return area, nil
}
