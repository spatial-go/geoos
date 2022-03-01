package planar

import (
	"github.com/spatial-go/geoos/space"
)

// Contains space.Geometry A contains space.Geometry B if and only if no points of B lie in the exterior of A,
// and at least one point of the interior of B lies in the interior of A.
// An important subtlety of this definition is that A does not contain its boundary, but A does contain itself.
// Returns TRUE if geometry B is completely inside geometry A.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (g *megrezAlgorithm) Contains(A, B space.Geometry) (bool, error) {
	return g.topog.Contains(A, B)
}

// CoveredBy returns TRUE if no point in space.Geometry A is outside space.Geometry B
func (g *megrezAlgorithm) CoveredBy(A, B space.Geometry) (bool, error) {
	return g.topog.CoveredBy(A, B)
}

// Covers returns TRUE if no point in space.Geometry B is outside space.Geometry A
func (g *megrezAlgorithm) Covers(A, B space.Geometry) (bool, error) {
	return g.topog.Covers(A, B)
}

// Crosses takes two geometry objects and returns TRUE if their intersection "spatially cross",
// that is, the geometries have some, but not all interior points in common.
// The intersection of the interiors of the geometries must not be the empty set
// and must have a dimensionality less than the maximum dimension of the two input geometries.
// Additionally, the intersection of the two geometries must not equal either of the source geometries.
// Otherwise, it returns FALSE.
func (g *megrezAlgorithm) Crosses(A, B space.Geometry) (bool, error) {
	return g.topog.Crosses(A, B)
}

// Disjoint Overlaps, Touches, Within all imply geometries are not spatially disjoint.
// If any of the aforementioned returns true, then the geometries are not spatially disjoint.
// Disjoint implies false for spatial intersection.
func (g *megrezAlgorithm) Disjoint(A, B space.Geometry) (bool, error) {
	return g.topog.Disjoint(A, B)
}

// Intersects If a geometry  shares any portion of space then they intersect
func (g *megrezAlgorithm) Intersects(A, B space.Geometry) (bool, error) {
	return g.topog.Intersects(A, B)
}

// Overlaps returns TRUE if the Geometries "spatially overlap".
// By that we mean they intersect, but one does not completely contain another.
func (g *megrezAlgorithm) Overlaps(A, B space.Geometry) (bool, error) {
	return g.topog.Overlaps(A, B)
}

// Relate computes the intersection matrix (Dimensionally Extended
// Nine-Intersection Model (DE-9IM) matrix) for the spatial relationship between
// the two geometries.
func (g *megrezAlgorithm) Relate(s, d space.Geometry) (string, error) {
	return g.topog.Relate(s, d)
}

// Touches returns TRUE if the only points in common between A and B lie in the union of the boundaries of A and B.
// The touches relation applies to all Area/Area, Line/Line, Line/Area, Point/Area and Point/Line pairs of relationships,
// but not to the Point/Point pair.
func (g *megrezAlgorithm) Touches(A, B space.Geometry) (bool, error) {
	return g.topog.Touches(A, B)
}

// Within returns TRUE if geometry A is completely inside geometry B.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (g *megrezAlgorithm) Within(A, B space.Geometry) (bool, error) {
	return g.topog.Within(A, B)
}
