package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/buffer/simplify"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/operation"
)

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

// IsCollection returns true if the Geometry is  collection.
func (mp MultiPolygon) IsCollection() bool {
	return true
}

// ToMatrix returns the Steric of a  geometry.
func (mp MultiPolygon) ToMatrix() matrix.Steric {
	matr := matrix.Collection{}
	for _, v := range mp {
		matr = append(matr, v.ToMatrix())
	}
	return matr
}

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

// EqualsMultiPolygon compares two multi-polygons.
func (mp MultiPolygon) EqualsMultiPolygon(multiPolygon MultiPolygon) bool {
	if len(mp) != len(multiPolygon) {
		return false
	}

	for i, p := range mp {
		if !p.Equals(multiPolygon[i]) {
			return false
		}
	}

	return true
}

// Equals checks if the MultiPolygon represents the same Geometry or vector.
func (mp MultiPolygon) Equals(g Geometry) bool {
	if g.GeoJSONType() != mp.GeoJSONType() {
		return false
	}
	return mp.EqualsMultiPolygon(g.(MultiPolygon))
}

// EqualsExact Returns true if the two Geometries are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (mp MultiPolygon) EqualsExact(g Geometry, tolerance float64) bool {
	if mp.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	for i, v := range mp {
		if !v.EqualsExact((g.(MultiPolygon)[i]), tolerance) {
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

// IsEmpty returns true if the Geometry is empty.
func (mp MultiPolygon) IsEmpty() bool {
	return len(mp) == 0
}

// Distance returns distance Between the two Geometry.
func (mp MultiPolygon) Distance(g Geometry) (float64, error) {
	return Distance(mp, g, measure.PlanarDistance)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (mp MultiPolygon) SpheroidDistance(g Geometry) (float64, error) {
	return Distance(mp, g, measure.SpheroidDistance)
}

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
func (mp MultiPolygon) Boundary() (Geometry, error) {
	if mp.IsEmpty() {
		return MultiLineString{}, nil
	}
	rings := MultiLineString{}
	for _, p := range mp {
		if r, err := p.Boundary(); err == nil {
			rings = append(rings, r.(MultiLineString)...)
		}
	}
	return rings, nil
}

// Length Returns the length of this MultiPolygon
func (mp MultiPolygon) Length() float64 {
	length := 0.0
	for _, v := range mp {
		length += v.Length()
	}
	return length
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (mp MultiPolygon) IsSimple() bool {
	if mp.IsEmpty() {
		return true
	}
	vop := &operation.ValidOP{Steric: mp.ToMatrix()}
	return vop.IsSimple()
}

// Centroid Computes the centroid point of a geometry.
func (mp MultiPolygon) Centroid() Point {
	return Centroid(mp)
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (mp MultiPolygon) UniquePoints() MultiPoint {
	mult := MultiPoint{}
	for _, v := range mp {
		mult = append(mult, v.UniquePoints()...)
	}
	return mult
}

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (mp MultiPolygon) Simplify(tolerance float64) Geometry {
	result := simplify.Simplify(mp.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (mp MultiPolygon) SimplifyP(tolerance float64) Geometry {
	tls := &simplify.TopologyPreservingSimplifier{}
	result := tls.Simplify(mp.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// Buffer Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (mp MultiPolygon) Buffer(width float64, quadsegs int) Geometry {
	buff := buffer.Buffer(mp.ToMatrix(), width, quadsegs)
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
func (mp MultiPolygon) BufferInMeter(width float64, quadsegs int) Geometry {
	return BufferInMeter(mp, width, quadsegs)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (mp MultiPolygon) Envelope() Geometry {
	return mp.Bound().ToPolygon()
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (mp MultiPolygon) ConvexHull() Geometry {
	result := buffer.ConvexHullWithGeom(mp.ToMatrix()).ConvexHull()
	return TransGeometry(result)
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (mp MultiPolygon) PointOnSurface() Geometry {
	m := buffer.InteriorPoint(mp.ToMatrix())
	return Point(m)
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (mp MultiPolygon) IsClosed() bool {
	return true
}

// IsRing returns true if the lineal geometry has the ring property.
func (mp MultiPolygon) IsRing() bool {
	return mp.IsClosed() && mp.IsSimple()
}

// IsValid returns true if the  geometry is valid.
func (mp MultiPolygon) IsValid() bool {
	for _, v := range mp {
		if !v.IsValid() {
			return false
		}
	}
	return true
}

// IsCorrect returns true if the geometry struct is Correct.
func (mp MultiPolygon) IsCorrect() bool {
	for _, v := range mp {
		if !v.IsCorrect() {
			return false
		}
	}
	return true
}

// CoordinateSystem return Coordinate System.
func (mp MultiPolygon) CoordinateSystem() int {
	return defaultCoordinateSystem()
}

// Filter Performs an operation with the provided .
func (mp MultiPolygon) Filter(f matrix.Filter) Geometry {
	if f.IsChanged() {
		mPoly := mp[:0]
		for _, v := range mp {
			p := Polygon(v).Filter(f)
			mPoly = append(mPoly, p.(Polygon))
		}
		return mPoly
	}
	for _, v := range mp {
		_ = Polygon(v).Filter(f)
	}
	return mp
}

// Geom return Geometry without Coordinate System.
func (mp MultiPolygon) Geom() Geometry {
	return mp
}
