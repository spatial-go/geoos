package space

import (
	"errors"
)

// A Collection is a collection of geometries that is also a Geometry.
type Collection []Geometry

// GeoJSONType returns the geometry collection type.
func (c Collection) GeoJSONType() string {
	return TypeCollection
}

// Dimensions returns the max of the dimensions of the collection.
func (c Collection) Dimensions() int {
	max := -1
	for _, g := range c {
		if d := g.Dimensions(); d > max {
			max = d
		}
	}
	return max
}

// Nums ...
func (c Collection) Nums() int {
	return len(c)
}

// Bound returns the bounding box of all the Geometries combined.
func (c Collection) Bound() Bound {
	if len(c) == 0 {
		return emptyBound
	}

	var b Bound
	start := -1

	for i, g := range c {
		if g != nil {
			start = i
			b = g.Bound()
			break
		}
	}

	if start == -1 {
		return emptyBound
	}

	for i := start + 1; i < len(c); i++ {
		if c[i] == nil {
			continue
		}

		b = b.Union(c[i].Bound())
	}

	return b
}

// EqualCollection compares two collections. Returns true if lengths are the same
// and all the sub geometries are the same and in the same order.
func (c Collection) EqualCollection(collection Collection) bool {
	if len(c) != len(collection) {
		return false
	}
	for i, g := range c {
		if !g.Equal(collection[i]) {
			return false
		}
	}
	return true
}

// Equal checks if the Collection represents the same Geometry or vector.
func (c Collection) Equal(g Geometry) bool {
	if g.GeoJSONType() != c.GeoJSONType() {
		return false
	}
	return c.EqualCollection(g.(Collection))
}

// EqualsExact Returns true if the two Geometrys are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (c Collection) EqualsExact(g Geometry, tolerance float64) bool {
	if c.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	for i, v := range c {
		if v.EqualsExact((g.(Collection)[i]), tolerance) {
			return false
		}
	}
	return true
}

// Area returns the area of a Collection geometry.
func (c Collection) Area() (float64, error) {
	area := 0.0
	for _, g := range c {
		switch g.GeoJSONType() {
		case TypePolygon:
			if areaOfPolygon, err := g.(Polygon).Area(); err == nil {
				area += areaOfPolygon
			} else {
				return 0, nil
			}
		case TypeMultiPolygon:
			if areaOfMultiPolygon, err := g.(MultiPolygon).Area(); err == nil {
				area += areaOfMultiPolygon
			} else {
				return 0, nil
			}
		default:

		}
	}
	return area, nil
}

// IsEmpty returns true if the Geometry is empty.
func (c Collection) IsEmpty() bool {
	return c == nil || len(c) == 0
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (c Collection) SpheroidDistance(g Geometry) (float64, error) {
	if c.IsEmpty() && g.IsEmpty() {
		return 0, nil
	}
	if c.IsEmpty() != g.IsEmpty() {
		return 0, errors.New("Geometry is nil")
	}
	var dist float64
	for _, v := range c {
		if distP, _ := v.SpheroidDistance(g); dist > distP {
			dist = distP
		}
	}
	return dist, nil
}

// Distance returns distance Between the two Geometry.
func (c Collection) Distance(g Geometry) (float64, error) {
	if c.IsEmpty() && g.IsEmpty() {
		return 0, nil
	}
	if c.IsEmpty() != g.IsEmpty() {
		return 0, errors.New("Geometry is nil")
	}
	var dist float64
	for _, v := range c {
		if distP, _ := v.Distance(g); dist > distP {
			dist = distP
		}
	}
	return dist, nil
}

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
func (c Collection) Boundary() (Geometry, error) {
	return nil, errors.New("Operation does not support GeometryCollection arguments")
}

// Length Returns the length of this Collection
func (c Collection) Length() float64 {
	length := 0.0
	for _, v := range c {
		length += v.Length()
	}
	return length
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (c Collection) IsSimple() bool {
	elem := ElementValid{c}
	return elem.IsSimple()
}

// Centroid Computes the centroid point of a geometry.
func (c Collection) Centroid() Point {
	return Centroid(c)
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (c Collection) UniquePoints() MultiPoint {
	mult := MultiPoint{}
	for _, v := range c {
		mult = append(mult, v.UniquePoints()...)
	}
	return mult
}
