package space

import (
	"math/rand"
	"reflect"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/space/spaceerr"
)

// Point describes a geographic point
type Point matrix.Matrix

// GeoJSONType returns GeoJSON type for the point
func (p Point) GeoJSONType() string {
	return TypePoint
}

// Dimensions returns 0 because a point is a 0d object.
func (p Point) Dimensions() int {
	return 0
}

// Bound returns a single point bound of the point.
func (p Point) Bound() Bound {
	return Bound{p, p}
}

// Nums num of points
func (p Point) Nums() int {
	return 1
}

// IsCollection returns true if the Geometry is  collection.
func (p Point) IsCollection() bool {
	return false
}

// Lat returns the vertical, latitude coordinate of the point.
func (p Point) Lat() float64 {
	return p[1]
}

// Lon returns the horizontal, longitude coordinate of the point.
func (p Point) Lon() float64 {
	return p[0]
}

// Y returns the vertical coordinate of the point.
func (p Point) Y() float64 {
	return p[1]
}

// X returns the horizontal coordinate of the point.
func (p Point) X() float64 {
	return p[0]
}

// EqualsPoint checks if the point represents the same point or vector.
func (p Point) EqualsPoint(point Point) bool {
	return matrix.Matrix(p).Equals(matrix.Matrix(point))
}

// Equals checks if the point represents the same Geometry or vector.
func (p Point) Equals(g Geometry) bool {
	if g.GeoJSONType() != p.GeoJSONType() {
		return false
	}
	return p.EqualsPoint(g.(Point))
}

// EqualsExact Returns true if the two Geometrys are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (p Point) EqualsExact(g Geometry, tolerance float64) bool {
	if g.GeoJSONType() != p.GeoJSONType() {
		return false
	}
	if p.IsEmpty() && g.IsEmpty() {
		return true
	}
	if p.IsEmpty() != g.IsEmpty() {
		return false
	}
	if tolerance == 0 {
		return p.Equals(g)
	}
	return measure.PlanarDistance(matrix.Matrix(p), matrix.Matrix(g.(Point))) <= tolerance
}

// Generate implements the Generator interface for Points
func (p Point) Generate(r *rand.Rand, _ int) reflect.Value {
	for i := range p {
		p[i] = r.Float64()
	}
	return reflect.ValueOf(p)
}

// ToMatrix returns the Steric of a  geometry.
func (p Point) ToMatrix() matrix.Steric {
	return matrix.Matrix(p)
}

// Area returns the area of a polygonal geometry. The area of a point is 0.
func (p Point) Area() (float64, error) {
	return 0.0, nil
}

// IsEmpty returns true if the Geometry is empty.
func (p Point) IsEmpty() bool {
	return p == nil || len(p) == 0
}

// Distance returns distance Between the two Geometry.
func (p Point) Distance(g Geometry) (float64, error) {
	return Distance(p, g, measure.PlanarDistance)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (p Point) SpheroidDistance(g Geometry) (float64, error) {
	return Distance(p, g, measure.SpheroidDistance)
}

// Boundary returns the closure of the combinatorial boundary of this Geometry.
func (p Point) Boundary() (Geometry, error) {
	return nil, spaceerr.ErrBoundBeNil
}

// IsSimple returns true if this Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (p Point) IsSimple() bool {
	return true
}

// Length Returns the length of this geometry
func (p Point) Length() float64 {
	return 0.0
}

// Centroid Computes the centroid point of a geometry.
func (p Point) Centroid() Point {
	return p
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (p Point) UniquePoints() MultiPoint {
	return MultiPoint{p}
}
