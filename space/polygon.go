package space

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
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

// IsRectangle returns true if  the polygon is rectangle.
func (p Polygon) IsRectangle() bool {

	if p.IsEmpty() || len(p) > 1 {
		return false
	}
	if len(p[0]) != 5 {
		return false
	}
	// check vertices have correct values
	for i := 0; i < 5; i++ {
		x := p[0][i][0]
		if !(x == p.Bound().Min.X() || x == p.Bound().Max.X()) {
			return false
		}
		y := p[0][i][1]
		if !(y == p.Bound().Min.Y() || y == p.Bound().Max.Y()) {
			return false
		}
	}

	// check vertices are in right order
	for i := 0; i < 4; i++ {
		x0 := p[0][i][0]
		y0 := p[0][i][1]
		x1 := p[0][i+1][0]
		y1 := p[0][i+1][1]
		xChanged := x0 != x1
		yChanged := y0 != y1
		if xChanged == yChanged {
			return false
		}
	}
	return true
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

// EqualsExact Returns true if the two Geometrys are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (p Polygon) EqualsExact(g Geometry, tolerance float64) bool {
	if p.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	pol := g.(Polygon)
	if p.IsEmpty() && g.IsEmpty() {
		return true
	}
	if p.IsEmpty() != g.IsEmpty() {
		return false
	}
	if len(p) != len(pol) {
		return false
	}

	for i, v := range p {
		if LineString(v).EqualsExact(LineString(pol[i]), tolerance) {
			return false
		}
	}
	return true
}

// Area returns the area of a polygonal geometry.
func (p Polygon) Area() (float64, error) {
	return measure.AreaOfPolygon(p.ToMatrix()), nil
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

// IsEmpty returns true if the Geometry is empty.
func (p Polygon) IsEmpty() bool {
	return p == nil || len(p) == 0
}

// Distance returns distance Between the two Geometry.
func (p Polygon) Distance(g Geometry) (float64, error) {
	elem := &Element{p}
	return elem.distanceWithFunc(g, measure.PlanarDistance)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (p Polygon) SpheroidDistance(g Geometry) (float64, error) {
	elem := &Element{p}
	return elem.distanceWithFunc(g, measure.SpheroidDistance)
}

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
func (p Polygon) Boundary() (Geometry, error) {
	if p.IsEmpty() {
		return MultiLineString{}, nil
	}
	if len(p) <= 1 {
		return LineString(p[0]), nil
	}
	rings := MultiLineString{}
	for _, v := range p {
		rings = append(rings, v)
	}
	return rings, nil
}

// Length Returns the length of this Polygon
func (p Polygon) Length() float64 {
	length := 0.0
	for _, v := range p {
		length += LineString(v).Length()
	}
	return length
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (p Polygon) IsSimple() bool {
	elem := ElementValid{p}
	return elem.IsSimple()
}

// Shell returns shell..
func (p Polygon) Shell() Ring {
	return p[0]
}

// Holes returns Holes..
func (p Polygon) Holes() []Ring {
	holes := []Ring{}
	for _, v := range p[1:] {
		holes = append(holes, v)
	}
	return holes
}

// Centroid Computes the centroid point of a geometry.
func (p Polygon) Centroid() Point {
	return Centroid(p)
}
