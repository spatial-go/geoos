package geoos

import (
	"math/rand"
	"reflect"
)

// Point describes a geographic point
type Point [2]float64

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

// Equal checks if the point represents the same point or vector.
func (p Point) Equal(point Point) bool {
	return p[0] == point[0] && p[1] == point[1]
}

// Generate implements the Generator interface for Points
func (p Point) Generate(r *rand.Rand, _ int) reflect.Value {
	for i := range p {
		p[i] = r.Float64()
	}
	return reflect.ValueOf(p)
}
