package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/buffer/simplify"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/operation"
	"github.com/spatial-go/geoos/space/spaceerr"
)

// A MultiPoint represents a set of points .
type MultiPoint []Point

// GeoJSONType returns the GeoJSON type for the object.
func (mp MultiPoint) GeoJSONType() string {
	return TypeMultiPoint
}

// Dimensions returns 0 because a MultiPoint is a 0d object.
func (mp MultiPoint) Dimensions() int {
	return 0
}

// Nums num of multiPoint.
func (mp MultiPoint) Nums() int {
	return len(mp)
}

// IsCollection returns true if the Geometry is  collection.
func (mp MultiPoint) IsCollection() bool {
	return true
}

// ToMatrix returns the Steric of a  geometry.
func (mp MultiPoint) ToMatrix() matrix.Steric {
	matr := matrix.Collection{}
	for _, v := range mp {
		matr = append(matr, v.ToMatrix())
	}
	return matr
}

// Bound returns a bound around the points. Uses rectangular coordinates.
func (mp MultiPoint) Bound() Bound {
	if len(mp) == 0 {
		return emptyBound
	}

	b := Bound{mp[0], mp[0]}
	for _, p := range mp {
		b = b.Extend(p)
	}

	return b
}

// EqualsMultiPoint compares two MultiPoint objects. Returns true if lengths are the same
// and all points are Equal, and in the same order.
func (mp MultiPoint) EqualsMultiPoint(multiPoint MultiPoint) bool {
	if len(mp) != len(multiPoint) {
		return false
	}
	for i, v := range mp.ToPointArray() {
		if !v.Equals(Point(multiPoint[i])) {
			return false
		}
	}
	return true
}

// Equals checks if the MultiPoint represents the same Geometry or vector.
func (mp MultiPoint) Equals(g Geometry) bool {
	if g.GeoJSONType() != mp.GeoJSONType() {
		return false
	}
	return mp.EqualsMultiPoint(g.(MultiPoint))
}

// EqualsExact Returns true if the two Geometries are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (mp MultiPoint) EqualsExact(g Geometry, tolerance float64) bool {
	if mp.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	for i, v := range mp {
		if !v.EqualsExact((g.(MultiPoint)[i]), tolerance) {
			return false
		}
	}
	return true
}

// Area returns the area of a polygonal geometry. The area of a multipoint is 0.
func (mp MultiPoint) Area() (float64, error) {
	return 0.0, nil
}

// ToPointArray returns the PointArray
func (mp MultiPoint) ToPointArray() (pa []Point) {
	return []Point(mp)
}

// IsEmpty returns true if the Geometry is empty.
func (mp MultiPoint) IsEmpty() bool {
	return len(mp) == 0
}

// Distance returns distance Between the two Geometry.
func (mp MultiPoint) Distance(g Geometry) (float64, error) {
	return Distance(mp, g, measure.PlanarDistance)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (mp MultiPoint) SpheroidDistance(g Geometry) (float64, error) {
	return Distance(mp, g, measure.SpheroidDistance)
}

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
func (mp MultiPoint) Boundary() (Geometry, error) {
	return nil, spaceerr.ErrNotSupportCollection
}

// Length Returns the length of this geometry
func (mp MultiPoint) Length() float64 {
	return 0.0
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (mp MultiPoint) IsSimple() bool {
	if mp.IsEmpty() {
		return true
	}
	vop := &operation.ValidOP{Steric: mp.ToMatrix()}
	return vop.IsSimple()
}

// Centroid Computes the centroid point of a geometry.
func (mp MultiPoint) Centroid() Point {
	return Centroid(mp)
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (mp MultiPoint) UniquePoints() MultiPoint {
	return mp
}

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (mp MultiPoint) Simplify(tolerance float64) Geometry {
	result := simplify.Simplify(mp.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (mp MultiPoint) SimplifyP(tolerance float64) Geometry {
	tls := &simplify.TopologyPreservingSimplifier{}
	result := tls.Simplify(mp.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// Buffer Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (mp MultiPoint) Buffer(width float64, quadsegs int) Geometry {
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
func (mp MultiPoint) BufferInMeter(width float64, quadsegs int) Geometry {
	return BufferInMeter(mp, width, quadsegs)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (mp MultiPoint) Envelope() Geometry {
	return mp.Bound().ToPolygon()
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (mp MultiPoint) ConvexHull() Geometry {
	result := buffer.ConvexHullWithGeom(mp.ToMatrix()).ConvexHull()
	return TransGeometry(result)
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (mp MultiPoint) PointOnSurface() Geometry {
	m := buffer.InteriorPoint(mp.ToMatrix())
	return Point(m)
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (mp MultiPoint) IsClosed() bool {
	return true
}

// IsRing returns true if the lineal geometry has the ring property.
func (mp MultiPoint) IsRing() bool {
	return mp.IsClosed() && mp.IsSimple()
}

// IsValid returns true if the  geometry is valid.
func (mp MultiPoint) IsValid() bool {
	for _, v := range mp {
		if !v.IsValid() {
			return false
		}
	}
	return true
}

// IsCorrect returns true if the geometry struct is Correct.
func (mp MultiPoint) IsCorrect() bool {
	for _, v := range mp {
		if !v.IsCorrect() {
			return false
		}
	}
	return true
}

// CoordinateSystem return Coordinate System.
func (mp MultiPoint) CoordinateSystem() int {
	return defaultCoordinateSystem()
}

// Filter Performs an operation with the provided .
func (mp MultiPoint) Filter(f matrix.Filter) Geometry {
	f.FilterMatrixes(matrix.TransMatrixes(mp.ToMatrix()))
	if f.IsChanged() {
		mp = mp[:0]
		for _, v := range f.Matrixes() {
			mp = append(mp, Point(v))
		}
		return mp
	}
	return mp
}

// Geom return Geometry without Coordinate System.
func (mp MultiPoint) Geom() Geometry {
	return mp
}
