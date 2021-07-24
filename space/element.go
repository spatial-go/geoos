package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/operation"
	"github.com/spatial-go/geoos/algorithm/relate"
	"github.com/spatial-go/geoos/space/spaceerr"
)

// Line  straight line  .
type Line struct {
	Start, End Point
}

// ElementValid describes a geographic Element Valid
type ElementValid struct {
	Geometry
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (el *ElementValid) IsClosed() bool {
	switch el.GeoJSONType() {
	case TypeLineString:
		return el.Geometry.(LineString).IsClosed()
	case TypeMultiLineString:
		return el.Geometry.(MultiLineString).IsClosed()
	}
	return true
}

// IsSimple Computes simplicity for geometries.
func (el *ElementValid) IsSimple() bool {
	if el.IsEmpty() {
		return true
	}
	vop := &operation.ValidOP{Steric: el.ToMatrix()}
	return vop.IsSimple()
}

// Centroid Computes the centroid point of a geometry.
func Centroid(geom Geometry) Point {
	cent := &buffer.Centroid{}

	if geom == nil || geom.IsEmpty() {
		return nil
	}
	cent.Add(geom.ToMatrix())
	m := cent.GetCentroid()
	return Point(m)
}

// Relate Computes the  Intersection Matrix for the spatial relationship
// between two Sterics, using the default (OGC SFS) Boundary Node Rule
func Relate(a, b Geometry) (string, error) {
	if a.IsCollection() || b.IsCollection() {
		return "", spaceerr.ErrNotSupportCollection
	}
	rel := &relate.Relationship{Arg: []matrix.Steric{a.ToMatrix(), b.ToMatrix()},
		IntersectBound: a.Bound().IntersectsBound(b.Bound())}
	im := rel.IntersectionMatrix()
	return im.ToString(), nil
}

// Distance returns distance Between the two Geometry.
func Distance(from, to Geometry, f measure.Distance) (float64, error) {
	if from == nil || from.IsEmpty() ||
		to == nil || to.IsEmpty() {
		return 0, nil
	}
	elem := &measure.ElementDistance{From: from.ToMatrix(), To: to.ToMatrix(), F: f}
	return elem.Distance()
}

// Within returns TRUE if geometry A is completely inside geometry B.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func Within(A, B Geometry) (bool, error) {
	intersectBound := true
	if inter, ret := aInB(A, B); !ret {
		intersectBound = inter
	} else {
		return inter, nil
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsWithin(), nil
}

// Contains space.Geometry A contains space.Geometry B if and only if no points of B lie in the exterior of A,
// and at least one point of the interior of B lies in the interior of A.
// An important subtlety of this definition is that A does not contain its boundary, but A does contain itself.
// Returns TRUE if geometry B is completely inside geometry A.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func Contains(A, B Geometry) (bool, error) {
	intersectBound := true
	if inter, ret := aInB(B, A); !ret {
		intersectBound = inter
	} else {
		return inter, nil
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsContains(), nil
}

// Covers returns TRUE if no point in space.Geometry B is outside space.Geometry A
func Covers(A, B Geometry) (bool, error) {
	intersectBound := true
	if inter, ret := aInB(B, A); !ret {
		intersectBound = inter
	} else {
		return inter, nil
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsCovers(), nil
}

// CoveredBy returns TRUE if no point in space.Geometry A is outside space.Geometry B
func CoveredBy(A, B Geometry) (bool, error) {
	intersectBound := true
	if inter, ret := aInB(A, B); !ret {
		intersectBound = inter
	} else {
		return inter, nil
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsCoveredBy(), nil
}

// Crosses takes two geometry objects and returns TRUE if their intersection "spatially cross",
// that is, the geometries have some, but not all interior points in common.
// The intersection of the interiors of the geometries must not be the empty set
// and must have a dimensionality less than the maximum dimension of the two input geometries.
// Additionally, the intersection of the two geometries must not equal either of the source geometries.
// Otherwise, it returns FALSE.
func Crosses(A, B Geometry) (bool, error) {
	intersectBound := B.Bound().IntersectsBound(A.Bound())
	if B.Bound().ContainsBound(A.Bound()) || B.Bound().ContainsBound(A.Bound()) {
		intersectBound = true
	}
	if !intersectBound {
		return false, nil
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsCrosses(A.Dimensions(), B.Dimensions()), nil
}

// Disjoint Overlaps, Touches, Within all imply geometries are not spatially disjoint.
// If any of the aforementioned returns true, then the geometries are not spatially disjoint.
// Disjoint implies false for spatial intersection.
func Disjoint(A, B Geometry) (bool, error) {
	intersectBound := B.Bound().IntersectsBound(A.Bound())
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsDisjoint(), nil
}

// Intersects If a geometry  shares any portion of space then they intersect
func Intersects(A, B Geometry) (bool, error) {
	intersectBound := B.Bound().IntersectsBound(A.Bound())
	if B.Bound().ContainsBound(A.Bound()) || B.Bound().ContainsBound(A.Bound()) {
		intersectBound = true
	}
	if !intersectBound {
		return false, nil
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsIntersects(), nil
}

// Touches returns TRUE if the only points in common between geom1 and geom2 lie in the union of the boundaries of geom1 and geom2.
// The ouches relation applies to all Area/Area, Line/Line, Line/Area, Point/Area and Point/Line pairs of relationships,
// but not to the Point/Point pair.
func Touches(A, B Geometry) (bool, error) {
	intersectBound := B.Bound().IntersectsBound(A.Bound())
	if B.Bound().ContainsBound(A.Bound()) || B.Bound().ContainsBound(A.Bound()) {
		intersectBound = true
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsTouches(A.Dimensions(), B.Dimensions()), nil
}

// Overlaps returns TRUE if the Geometries "spatially overlap".
// By that we mean they intersect, but one does not completely contain another.
func Overlaps(A, B Geometry) (bool, error) {
	intersectBound := A.Bound().IntersectsBound(B.Bound())
	if A.Bound().ContainsBound(B.Bound()) || A.Bound().ContainsBound(B.Bound()) {
		intersectBound = true
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsOverlaps(A.Dimensions(), B.Dimensions()), nil
}

func aInB(A, B Geometry) (bool, bool) {

	// optimization - lower dimension cannot contain areas
	if A.Dimensions() == 2 && B.Dimensions() < 2 {
		return false, true
	}
	// optimization - P cannot contain a non-zero-length L
	// Note that a point can contain a zero-length lineal geometry,
	// since the line has no boundary due to Mod-2 Boundary Rule
	if A.Dimensions() == 1 && B.Dimensions() < 1 && A.Length() > 0.0 {
		return false, true
	}
	// optimization - envelope test
	if A.Bound().ContainsBound(B.Bound()) {
		return false, true
	}
	// optimization for rectangle arguments
	if B.GeoJSONType() == TypePolygon && B.(Polygon).IsRectangle() {
		return B.Bound().ContainsBound(A.Bound()), true
	}

	intersectBound := B.Bound().IntersectsBound(A.Bound())
	if B.Bound().ContainsBound(A.Bound()) || B.Bound().ContainsBound(A.Bound()) {
		intersectBound = true
	}
	return intersectBound, false

}
