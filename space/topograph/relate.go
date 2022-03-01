package topograph

import (
	"github.com/spatial-go/geoos/algorithm/graph"
	"github.com/spatial-go/geoos/space"
)

// Topograph Computes the Intersection Matrix for the spatial relationship
// between two geometries , using the graph.
type Topograph struct {
}

// Relate Computes the  Intersection Matrix for the spatial relationship
// between two geometries, using the default (OGC SFS) Boundary Node Rule
func (t *Topograph) Relate(A, B space.Geometry) (string, error) {
	return graph.Relate(A.ToMatrix(), B.ToMatrix()), nil

}

// Within returns TRUE if geometry A is completely inside geometry B.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (t *Topograph) Within(A, B space.Geometry) (bool, error) {
	_, isAInB, isSure := AInB(A, B)
	if isSure {
		return isAInB, nil
	}
	im := graph.IM(A.ToMatrix(), B.ToMatrix())
	return im.IsWithin(), nil
}

// Contains space.Geometry A contains space.Geometry B if and only if no points of B lie in the exterior of A,
// and at least one point of the interior of B lies in the interior of A.
// An important subtlety of this definition is that A does not contain its boundary, but A does contain itself.
// Returns TRUE if geometry B is completely inside geometry A.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (t *Topograph) Contains(A, B space.Geometry) (bool, error) {
	_, isAInB, isSure := AInB(B, A)
	if isSure {
		return isAInB, nil
	}
	im := graph.IM(A.ToMatrix(), B.ToMatrix())
	return im.IsContains(), nil
}

// Covers returns TRUE if no point in space.Geometry B is outside space.Geometry A
func (t *Topograph) Covers(A, B space.Geometry) (bool, error) {
	_, isAInB, isSure := AInB(B, A)
	if isSure {
		return isAInB, nil
	}
	im := graph.IM(A.ToMatrix(), B.ToMatrix())
	return im.IsCovers(), nil
}

// CoveredBy returns TRUE if no point in space.Geometry A is outside space.Geometry B
func (t *Topograph) CoveredBy(A, B space.Geometry) (bool, error) {
	_, isAInB, isSure := AInB(A, B)
	if isSure {
		return isAInB, nil
	}
	im := graph.IM(A.ToMatrix(), B.ToMatrix())
	return im.IsCoveredBy(), nil
}

// Crosses takes two geometry objects and returns TRUE if their intersection "spatially cross",
// that is, the geometries have some, but not all interior points in common.
// The intersection of the interiors of the geometries must not be the empty set
// and must have a dimensionality less than the maximum dimension of the two input geometries.
// Additionally, the intersection of the two geometries must not equal either of the source geometries.
// Otherwise, it returns FALSE.
func (t *Topograph) Crosses(A, B space.Geometry) (bool, error) {
	intersectBound := B.Bound().IntersectsBound(A.Bound())
	if B.Bound().ContainsBound(A.Bound()) || A.Bound().ContainsBound(B.Bound()) {
		intersectBound = true
	}
	if !intersectBound {
		return false, nil
	}
	im := graph.IM(A.ToMatrix(), B.ToMatrix())
	return im.IsCrosses(A.Dimensions(), B.Dimensions()), nil
}

// Disjoint Overlaps, Touches, Within all imply geometries are not spatially disjoint.
// If any of the aforementioned returns true, then the geometries are not spatially disjoint.
// Disjoint implies false for spatial intersection.
func (t *Topograph) Disjoint(A, B space.Geometry) (bool, error) {
	im := graph.IM(A.ToMatrix(), B.ToMatrix())
	return im.IsDisjoint(), nil
}

// Intersects If a geometry  shares any portion of space then they intersect
func (t *Topograph) Intersects(A, B space.Geometry) (bool, error) {
	intersectBound := B.Bound().IntersectsBound(A.Bound())
	if B.Bound().ContainsBound(A.Bound()) || A.Bound().ContainsBound(B.Bound()) {
		intersectBound = true
	}
	if !intersectBound {
		return false, nil
	}
	im := graph.IM(A.ToMatrix(), B.ToMatrix())
	return im.IsIntersects(), nil
}

// Touches returns TRUE if the only points in common between geom1 and geom2 lie in the union of the boundaries of geom1 and geom2.
// The ouches relation applies to all Area/Area, Line/Line, Line/Area, Point/Area and Point/Line pairs of relationships,
// but not to the Point/Point pair.
func (t *Topograph) Touches(A, B space.Geometry) (bool, error) {
	im := graph.IM(A.ToMatrix(), B.ToMatrix())
	return im.IsTouches(A.Dimensions(), B.Dimensions()), nil
}

// Overlaps returns TRUE if the Geometries "spatially overlap".
// By that we mean they intersect, but one does not completely contain another.
func (t *Topograph) Overlaps(A, B space.Geometry) (bool, error) {
	im := graph.IM(A.ToMatrix(), B.ToMatrix())
	return im.IsOverlaps(A.Dimensions(), B.Dimensions()), nil
}

// AInB Returns isIntersect, isAInB, isSure.
func AInB(A, B space.Geometry) (isIntersect, isAInB, isSure bool) {

	// optimization - lower dimension cannot contain areas
	if A.Dimensions() == 2 && B.Dimensions() < 2 {
		isIntersect, isAInB, isSure = false, false, true
		return
	}
	// optimization - P cannot contain a non-zero-length L
	// Note that a point can contain a zero-length lineal geometry,
	// since the line has no boundary due to Mod-2 Boundary Rule
	if A.Dimensions() == 1 && B.Dimensions() < 1 && A.Length() > 0.0 {
		isIntersect, isAInB, isSure = false, false, true
		return
	}
	// optimization - envelope test
	if A.Bound().ContainsBound(B.Bound()) {
		isIntersect, isAInB, isSure = false, false, true
		return
	}
	//optimization for rectangle arguments
	if B.GeoJSONType() == space.TypePolygon && B.(space.Polygon).IsRectangle() {
		isIntersect, isAInB, isSure = false, B.Bound().ContainsBound(A.Bound()), true
		return
	}

	isIntersect, isAInB, isSure = B.Bound().IntersectsBound(A.Bound()), false, false

	return

}
