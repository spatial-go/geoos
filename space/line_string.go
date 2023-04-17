package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/buffer/simplify"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/operation"
	"github.com/spatial-go/geoos/space/spaceerr"
)

// LineString represents a set of points to be thought of as a polyline.
type LineString matrix.LineMatrix

// GeoJSONType returns the GeoJSON type for the linestring.
func (ls LineString) GeoJSONType() string {
	return TypeLineString
}

// Dimensions returns 1 because a LineString is a 1d object.
func (ls LineString) Dimensions() int {
	return 1
}

// Nums num of linestrings
func (ls LineString) Nums() int {
	return 1
}

// IsCollection returns true if the Geometry is  collection.
func (ls LineString) IsCollection() bool {
	return false
}

// ToMatrix returns the LineMatrix of a Ring geometry.
func (ls LineString) ToMatrix() matrix.Steric {
	return matrix.LineMatrix(ls)
}

// Bound returns a rect around the line string. Uses rectangular coordinates.
func (ls LineString) Bound() Bound {
	if len(ls) == 0 {
		return emptyBound
	}

	b := Bound{ls[0], ls[0]}
	for _, p := range ls {
		b = b.Extend(p)
	}

	return b
}

// EqualsLineString compares two line strings. Returns true if lengths are the same
// and all points are Equal.
func (ls LineString) EqualsLineString(lineString LineString) bool {
	if len(ls) != len(lineString) {
		return false
	}
	for i, v := range ls.ToPointArray() {
		if !v.Equals(Point(lineString[i])) {
			return false
		}
	}
	return true
}

// Equals checks if the LineString represents the same Geometry or vector.
func (ls LineString) Equals(g Geometry) bool {
	if g.GeoJSONType() != ls.GeoJSONType() {
		return false
	}
	return ls.EqualsLineString(g.(LineString))
}

// EqualsExact Returns true if the two Geometries are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (ls LineString) EqualsExact(g Geometry, tolerance float64) bool {
	if ls.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	line := g.(LineString)
	if ls.IsEmpty() && g.IsEmpty() {
		return true
	}
	if ls.IsEmpty() != g.IsEmpty() {
		return false
	}
	if len(ls) != len(line) {
		return false
	}

	for i, v := range ls {
		if !Point(v).EqualsExact(Point(line[i]), tolerance) {
			return false
		}
	}
	return true
}

// Area returns the area of a polygonal geometry. The area of a LineString is 0.
func (ls LineString) Area() (float64, error) {
	return 0.0, nil
}

// ToPointArray returns the PointArray
func (ls LineString) ToPointArray() (la []Point) {
	for _, v := range ls {
		la = append(la, v)
	}
	return
}

// ToLineArray returns the LineArray
func (ls LineString) ToLineArray() (lines []Line) {
	for i := 0; i < len(ls)-1; i++ {
		lines = append(lines, Line{Point(ls[i]), Point(ls[i+1])})
	}
	return
}

// IsEmpty returns true if the Geometry is empty.
func (ls LineString) IsEmpty() bool {
	return len(ls) == 0
}

// Distance returns distance Between the two Geometry.
func (ls LineString) Distance(g Geometry) (float64, error) {
	return Distance(ls, g, measure.PlanarDistance)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (ls LineString) SpheroidDistance(g Geometry) (float64, error) {
	return Distance(ls, g, measure.SpheroidDistance)
}

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
// The boundary of a lineal geometry is always a zero-dimensional geometry (which may be empty).
func (ls LineString) Boundary() (Geometry, error) {
	if ls.IsClosed() {
		return nil, spaceerr.ErrBoundBeNil
	}
	return MultiPoint{ls[0], ls[len(ls)-1]}, nil
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (ls LineString) IsClosed() bool {
	return Point(ls[0]).Equals(Point(ls[len(ls)-1]))
}

// Length Returns the length of this LineString
func (ls LineString) Length() float64 {
	return measure.OfLine(matrix.LineMatrix(ls))
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (ls LineString) IsSimple() bool {
	if ls.IsEmpty() {
		return true
	}
	vop := &operation.ValidOP{Steric: ls.ToMatrix()}
	return vop.IsSimple()
}

// Centroid Computes the centroid point of a geometry.
func (ls LineString) Centroid() Point {
	return Centroid(ls)
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (ls LineString) UniquePoints() MultiPoint {
	mp := MultiPoint{}
	for _, v := range ls {
		mp = append(mp, v)
	}
	return mp
}

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (ls LineString) Simplify(tolerance float64) Geometry {
	result := simplify.Simplify(ls.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (ls LineString) SimplifyP(tolerance float64) Geometry {
	tls := &simplify.TopologyPreservingSimplifier{}
	result := tls.Simplify(ls.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// Buffer Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (ls LineString) Buffer(width float64, quadsegs int) Geometry {
	buff := buffer.Buffer(ls.ToMatrix(), width, quadsegs)
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
func (ls LineString) BufferInMeter(width float64, quadsegs int) Geometry {
	return BufferInMeter(ls, width, quadsegs)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (ls LineString) Envelope() Geometry {
	return ls.Bound().ToPolygon()
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (ls LineString) ConvexHull() Geometry {
	result := buffer.ConvexHullWithGeom(ls.ToMatrix()).ConvexHull()
	return TransGeometry(result)
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (ls LineString) PointOnSurface() Geometry {
	m := buffer.InteriorPoint(ls.ToMatrix())
	return Point(m)
}

// IsRing returns true if the lineal geometry has the ring property.
func (ls LineString) IsRing() bool {
	return ls.IsClosed() && ls.IsSimple()
}

// IsValid returns true if the  geometry is valid.
func (ls LineString) IsValid() bool {
	if ls.IsEmpty() {
		return false
	}
	for _, v := range ls {
		if !Point(v).IsValid() {
			return false
		}
	}
	return true
}

// IsCorrect returns true if the geometry struct is Correct.
func (ls LineString) IsCorrect() bool {
	if ls.IsEmpty() {
		return false
	}
	for _, v := range ls {
		if !Point(v).IsCorrect() {
			return false
		}
	}
	return true
}

// CoordinateSystem return Coordinate System.
func (ls LineString) CoordinateSystem() int {
	return defaultCoordinateSystem()
}

// Filter Performs an operation with the provided .
func (ls LineString) Filter(f matrix.Filter) Geometry {
	f.FilterMatrixes(matrix.TransMatrixes(ls.ToMatrix()))
	if f.IsChanged() {
		ls = ls[:0]
		for _, v := range f.Matrixes() {
			ls = append(ls, v)
		}
		return ls
	}
	return ls
}

// Geom return Geometry without Coordinate System.
func (ls LineString) Geom() Geometry {
	return ls
}
