//Package topograph the topological relations between points, line and surface entities in DE-9IM model are given.
package topological

import (
	"sync"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
	"github.com/spatial-go/geoos/space"
	"github.com/spatial-go/geoos/space/spaceerr"
	"github.com/spatial-go/geoos/space/topograph"
)

var topog topograph.Relationship
var once sync.Once

// NewTopological returns Relationship that is Topological.
func NewTopological() topograph.Relationship {
	once.Do(func() {
		topog = &Topological{}
	})
	return topog
}

// Topological Computes the Intersection Matrix for the spatial relationship
// between two geometries ,  using the default (OGC SFS) Boundary Node Rule
type Topological struct {
}

// Relate Computes the  Intersection Matrix for the spatial relationship
// between two geometries, using the default (OGC SFS) Boundary Node Rule
func (t *Topological) Relate(A, B space.Geometry) (string, error) {
	if A.IsCollection() || B.IsCollection() {
		return "", spaceerr.ErrNotSupportCollection
	}
	rel := &relate.Relationship{Arg: []matrix.Steric{A.ToMatrix(), B.ToMatrix()},
		IntersectBound: A.Bound().IntersectsBound(B.Bound())}
	im := rel.IntersectionMatrix()
	return im.ToString(), nil
}

// Within returns TRUE if geometry A is completely inside geometry B.
// For this func (t *Topological)tion to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (t *Topological) Within(A, B space.Geometry) (bool, error) {
	isIntersect, isAInB, isSure := topograph.AInB(A, B)
	if isSure {
		return isAInB, nil
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), isIntersect)
	return im.IsWithin(), nil
}

// Contains space.Geometry A contains space.Geometry B if and only if no points of B lie in the exterior of A,
// and at least one point of the interior of B lies in the interior of A.
// An important subtlety of this definition is that A does not contain its boundary, but A does contain itself.
// Returns TRUE if geometry B is completely inside geometry A.
// For this func (t *Topological)tion to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (t *Topological) Contains(A, B space.Geometry) (bool, error) {
	isIntersect, isAInB, isSure := topograph.AInB(B, A)
	if isSure {
		return isAInB, nil
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), isIntersect)
	return im.IsContains(), nil
}

// Covers returns TRUE if no point in space.Geometry B is outside space.Geometry A
func (t *Topological) Covers(A, B space.Geometry) (bool, error) {
	isIntersect, isAInB, isSure := topograph.AInB(B, A)
	if isSure {
		return isAInB, nil
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), isIntersect)
	return im.IsCovers(), nil
}

// CoveredBy returns TRUE if no point in space.Geometry A is outside space.Geometry B
func (t *Topological) CoveredBy(A, B space.Geometry) (bool, error) {
	isIntersect, isAInB, isSure := topograph.AInB(A, B)
	if isSure {
		return isAInB, nil
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), isIntersect)
	return im.IsCoveredBy(), nil
}

// Crosses takes two geometry objects and returns TRUE if their intersection "spatially cross",
// that is, the geometries have some, but not all interior points in common.
// The intersection of the interiors of the geometries must not be the empty set
// and must have a dimensionality less than the maximum dimension of the two input geometries.
// Additionally, the intersection of the two geometries must not equal either of the source geometries.
// Otherwise, it returns FALSE.
func (t *Topological) Crosses(A, B space.Geometry) (bool, error) {
	intersectBound := B.Bound().IntersectsBound(A.Bound())
	if B.Bound().ContainsBound(A.Bound()) || A.Bound().ContainsBound(B.Bound()) {
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
func (t *Topological) Disjoint(A, B space.Geometry) (bool, error) {
	intersectBound := B.Bound().IntersectsBound(A.Bound())
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsDisjoint(), nil
}

// Intersects If a geometry  shares any portion of space then they intersect
func (t *Topological) Intersects(A, B space.Geometry) (bool, error) {
	intersectBound := B.Bound().IntersectsBound(A.Bound())
	if B.Bound().ContainsBound(A.Bound()) || A.Bound().ContainsBound(B.Bound()) {
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
func (t *Topological) Touches(A, B space.Geometry) (bool, error) {
	intersectBound := B.Bound().IntersectsBound(A.Bound())
	if B.Bound().ContainsBound(A.Bound()) || A.Bound().ContainsBound(B.Bound()) {
		intersectBound = true
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsTouches(A.Dimensions(), B.Dimensions()), nil
}

// Overlaps returns TRUE if the Geometries "spatially overlap".
// By that we mean they intersect, but one does not completely contain another.
func (t *Topological) Overlaps(A, B space.Geometry) (bool, error) {
	intersectBound := A.Bound().IntersectsBound(B.Bound())
	if A.Bound().ContainsBound(B.Bound()) || B.Bound().ContainsBound(A.Bound()) {
		intersectBound = true
	}
	im := relate.IM(A.ToMatrix(), B.ToMatrix(), intersectBound)
	return im.IsOverlaps(A.Dimensions(), B.Dimensions()), nil
}
