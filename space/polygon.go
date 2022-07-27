package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/buffer/simplify"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/operation"
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

// IsCollection returns true if the Geometry is  collection.
func (p Polygon) IsCollection() bool {
	return false
}

// Bound returns a bound around the polygon.
func (p Polygon) Bound() Bound {
	if len(p) == 0 {
		return emptyBound
	}
	return p.ToRingArray()[0].Bound()
}

// ToMatrix returns the LineMatrix of a Ring geometry.
func (p Polygon) ToMatrix() matrix.Steric {
	return matrix.PolygonMatrix(p)
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

// EqualsPolygon comEqualPolygonpares two polygons. Returns true if lengths are the same
// and all points are Equal.
func (p Polygon) EqualsPolygon(polygon Polygon) bool {
	if len(p) != len(polygon) {
		return false
	}
	for i, v := range p.ToRingArray() {
		if !v.Equals(Ring(polygon[i])) {
			return false
		}
	}
	return true
}

// Equals checks if the Polygon represents the same Geometry or vector.
func (p Polygon) Equals(g Geometry) bool {
	if g.GeoJSONType() != p.GeoJSONType() {
		return false
	}
	return p.EqualsPolygon(g.(Polygon))
}

// EqualsExact Returns true if the two Geometries are exactly equal,
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
		if !LineString(v).EqualsExact(LineString(pol[i]), tolerance) {
			return false
		}
	}
	return true
}

// Area returns the area of a polygonal geometry.
func (p Polygon) Area() (float64, error) {
	return measure.AreaOfPolygon(p.ToMatrix().(matrix.PolygonMatrix)), nil
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
	return len(p) == 0
}

// Distance returns distance Between the two Geometry.
func (p Polygon) Distance(g Geometry) (float64, error) {
	return Distance(p, g, measure.PlanarDistance)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (p Polygon) SpheroidDistance(g Geometry) (float64, error) {
	return Distance(p, g, measure.SpheroidDistance)
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
	if p.IsEmpty() {
		return true
	}
	vop := &operation.ValidOP{Steric: p.ToMatrix()}
	return vop.IsSimple()
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

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (p Polygon) UniquePoints() MultiPoint {
	mp := MultiPoint{}
	for _, v := range p {
		mp = append(mp, Ring(v).UniquePoints()...)
	}
	return mp
}

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (p Polygon) Simplify(tolerance float64) Geometry {
	result := simplify.Simplify(p.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (p Polygon) SimplifyP(tolerance float64) Geometry {
	tls := &simplify.TopologyPreservingSimplifier{}
	result := tls.Simplify(p.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// Buffer Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (p Polygon) Buffer(width float64, quadsegs int) Geometry {
	buff := buffer.Buffer(p.ToMatrix(), width, quadsegs)
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
func (p Polygon) BufferInMeter(width float64, quadsegs int) Geometry {
	return BufferInMeter(p, width, quadsegs)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (p Polygon) Envelope() Geometry {
	return p.Bound().ToPolygon()
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (p Polygon) ConvexHull() Geometry {
	result := buffer.ConvexHullWithGeom(p.ToMatrix()).ConvexHull()
	return TransGeometry(result)
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (p Polygon) PointOnSurface() Geometry {
	m := buffer.InteriorPoint(p.ToMatrix())
	return Point(m)
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (p Polygon) IsClosed() bool {
	return true
}

// IsRing returns true if the lineal geometry has the ring property.
func (p Polygon) IsRing() bool {
	return p.IsClosed() && p.IsSimple()
}

// IsValid returns true if the  geometry is valid.
func (p Polygon) IsValid() bool {
	if p.IsEmpty() {
		return false
	}
	for _, v := range p {
		if !Ring(v).IsValid() {
			return false
		}
	}
	return true
}

// IsCorrect returns true if the geometry struct is Correct.
func (p Polygon) IsCorrect() bool {
	if p.IsEmpty() {
		return false
	}
	for _, v := range p {
		if !Ring(v).IsCorrect() {
			return false
		}
	}
	return true
}

// CoordinateSystem return Coordinate System.
func (p Polygon) CoordinateSystem() int {
	return defaultCoordinateSystem()
}

// Filter Performs an operation with the provided .
func (p Polygon) Filter(f matrix.Filter) Geometry {
	if f.IsChanged() {
		poly := Polygon{}
		for _, v := range p {
			f.Clear()
			r := Ring(v).Filter(f)
			poly = append(poly, r.(Ring))
		}
		return poly
	}
	for _, v := range p {
		_ = Ring(v).Filter(f)
	}
	return p
}

// Geom return Geometry without Coordinate System.
func (p Polygon) Geom() Geometry {
	return p
}
