package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/buffer/simplify"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/space/spaceerr"
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

// IsCollection returns true if the Geometry is  collection.
func (c Collection) IsCollection() bool {
	return true
}

// ToMatrix returns the Steric of a  geometry.
func (c Collection) ToMatrix() matrix.Steric {
	matr := matrix.Collection{}
	for _, v := range c {
		matr = append(matr, v.ToMatrix())
	}
	return matr
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

// EqualsCollection compares two collections. Returns true if lengths are the same
// and all the sub geometries are the same and in the same order.
func (c Collection) EqualsCollection(collection Collection) bool {
	if len(c) != len(collection) {
		return false
	}
	for i, g := range c {
		if !g.Equals(collection[i]) {
			return false
		}
	}
	return true
}

// Equals checks if the Collection represents the same Geometry or vector.
func (c Collection) Equals(g Geometry) bool {
	if g.GeoJSONType() != c.GeoJSONType() {
		return false
	}
	return c.EqualsCollection(g.(Collection))
}

// EqualsExact Returns true if the two Geometries are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (c Collection) EqualsExact(g Geometry, tolerance float64) bool {
	if c.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	for i, v := range c {
		if !v.EqualsExact((g.(Collection)[i]), tolerance) {
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
	return len(c) == 0
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (c Collection) SpheroidDistance(g Geometry) (float64, error) {
	if c.IsEmpty() && g.IsEmpty() {
		return 0, nil
	}
	if c.IsEmpty() != g.IsEmpty() {
		return 0, spaceerr.ErrNilGeometry
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
		return 0, spaceerr.ErrNilGeometry
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
	return nil, spaceerr.ErrNotSupportCollection
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
	elem := &GeometryValid{c, GCJ02}
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

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (c Collection) Simplify(tolerance float64) Geometry {
	result := simplify.Simplify(c.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (c Collection) SimplifyP(tolerance float64) Geometry {
	tls := &simplify.TopologyPreservingSimplifier{}
	result := tls.Simplify(c.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// Buffer Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (c Collection) Buffer(width float64, quadsegs int) Geometry {
	buff := buffer.Buffer(c.ToMatrix(), width, quadsegs)
	switch b := buff.(type) {
	case matrix.LineMatrix:
		return LineString(b)
	case matrix.PolygonMatrix:
		return Polygon(b)
	}
	return nil
}

// BufferInMeter Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (c Collection) BufferInMeter(width float64, quadsegs int) Geometry {
	return BufferInMeter(c, width, quadsegs)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (c Collection) Envelope() Geometry {
	return c.Bound().ToPolygon()
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (c Collection) ConvexHull() Geometry {
	result := buffer.ConvexHullWithGeom(c.ToMatrix()).ConvexHull()
	return TransGeometry(result)
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (c Collection) PointOnSurface() Geometry {
	m := buffer.InteriorPoint(c.ToMatrix())
	return Point(m)
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (c Collection) IsClosed() bool {
	return true
}

// IsRing returns true if the lineal geometry has the ring property.
func (c Collection) IsRing() bool {
	return c.IsClosed() && c.IsSimple()
}

// IsValid returns true if the  geometry is valid.
func (c Collection) IsValid() bool {
	for _, v := range c {
		if !v.IsValid() {
			return false
		}
	}
	return true
}

// IsCorrect returns true if the geometry struct is Correct.
func (c Collection) IsCorrect() bool {
	for _, v := range c {
		if !v.IsCorrect() {
			return false
		}
	}
	return true
}

// CoordinateSystem return Coordinate System.
func (c Collection) CoordinateSystem() int {
	return defaultCoordinateSystem()
}

// Filter Performs an operation with the provided .
func (c Collection) Filter(f matrix.Filter) Geometry {
	if f.IsChanged() {
		mc := c[:0]
		for _, v := range c {
			f.Clear()
			g := v.Filter(f)
			mc = append(mc, g)
		}
		return mc
	}
	for _, v := range c {
		_ = v.Filter(f)
	}
	return c
}

// Geom return Geometry without Coordinate System.
func (c Collection) Geom() Geometry {
	return c
}
