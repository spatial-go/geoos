package geoos

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Polygon is a closed area. The first LineString is the outer ring.
// The others are the holes. Each LineString is expected to be closed
// ie. the first point matches the last.
type Polygon matrix.PolygonMatrix

// GeoJSONType returns the GeoJSON type for the polygon.
func (p Polygon) GeoJSONType() string {
	return TypePolygon
}

// Dimensions returns 2 because a Polygon is a 2d object.
func (p Polygon) Dimensions() int {
	return 2
}

// Nums num of polygons
func (p Polygon) Nums() int {
	return 1
}

// Bound returns a bound around the polygon.
func (p Polygon) Bound() Bound {
	if len(p) == 0 {
		return emptyBound
	}
	return p.ToRingArray()[0].Bound()
}

// EqualPolygon comEqualPolygonpares two polygons. Returns true if lengths are the same
// and all points are Equal.
func (p Polygon) EqualPolygon(polygon Polygon) bool {
	if len(p) != len(polygon) {
		return false
	}
	for i, v := range p.ToRingArray() {
		if !v.Equal(Ring(polygon[i])) {
			return false
		}
	}
	return true
}

// Equal checks if the Polygon represents the same Geometry or vector.
func (p Polygon) Equal(g Geometry) bool {
	if g.GeoJSONType() != p.GeoJSONType() {
		return false
	}
	return p.EqualPolygon(g.(Polygon))
}

// Area returns the area of a polygonal geometry.
func (p Polygon) Area() (float64, error) {
	return algorithm.AreaOfPolygon(p.ToMatrix()), nil
}

// ToMatrix returns the PolygonMatrix of a polygonal geometry.
func (p Polygon) ToMatrix() matrix.PolygonMatrix {
	return matrix.PolygonMatrix(p)
}

// ToRingArray returns the RingArray
func (p Polygon) ToRingArray() (r []Ring) {
	for _, v := range p {
		r = append(r, v)
	}
	return
}
