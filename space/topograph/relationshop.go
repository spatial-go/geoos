//Package topograph the topological relations between points, line and surface entities in DE-9IM model are given.
package topograph

import (
	"sync"

	"github.com/spatial-go/geoos/space"
)

var topog Relationship
var once sync.Once

type newRelationship func() Relationship

// NormalRelationship returns normal Relationship.
func NormalRelationship() Relationship {
	return GetRelationship(NewTopograph)
}

// GetRelationship returns  algorithm by new Relationship.
func GetRelationship(f newRelationship) Relationship {
	return f()
}

// NewTopograph returns Relationship that is Topograph.
func NewTopograph() Relationship {
	once.Do(func() {
		topog = &Topograph{}
	})
	return topog
}

// NewTopological returns Relationship that is Topograph.
func NewTopological() Relationship {
	once.Do(func() {
		topog = &Topograph{}
	})
	return topog
}

// Relationship Computes the Intersection Matrix for the spatial relationship
// between two geometries
type Relationship interface {
	// Relate Computes the  Intersection Matrix for the spatial relationship
	// between two geometries, using the default (OGC SFS) Boundary Node Rule
	Relate(A, B space.Geometry) (string, error)

	// Within returns TRUE if geometry A is completely inside geometry B.
	// For this function to make sense, the source geometries must both be of the same coordinate projection,
	// having the same SRID.
	Within(A, B space.Geometry) (bool, error)

	// Contains space.Geometry A contains space.Geometry B if and only if no points of B lie in the exterior of A,
	// and at least one point of the interior of B lies in the interior of A.
	// An important subtlety of this definition is that A does not contain its boundary, but A does contain itself.
	// Returns TRUE if geometry B is completely inside geometry A.
	// For this function to make sense, the source geometries must both be of the same coordinate projection,
	// having the same SRID.
	Contains(A, B space.Geometry) (bool, error)

	// Covers returns TRUE if no point in space.Geometry B is outside space.Geometry A
	Covers(A, B space.Geometry) (bool, error)

	// CoveredBy returns TRUE if no point in space.Geometry A is outside space.Geometry B
	CoveredBy(A, B space.Geometry) (bool, error)

	// Crosses takes two geometry objects and returns TRUE if their intersection "spatially cross",
	// that is, the geometries have some, but not all interior points in common.
	// The intersection of the interiors of the geometries must not be the empty set
	// and must have a dimensionality less than the maximum dimension of the two input geometries.
	// Additionally, the intersection of the two geometries must not equal either of the source geometries.
	// Otherwise, it returns FALSE.
	Crosses(A, B space.Geometry) (bool, error)

	// Disjoint Overlaps, Touches, Within all imply geometries are not spatially disjoint.
	// If any of the aforementioned returns true, then the geometries are not spatially disjoint.
	// Disjoint implies false for spatial intersection.
	Disjoint(A, B space.Geometry) (bool, error)

	// Intersects If a geometry  shares any portion of space then they intersect
	Intersects(A, B space.Geometry) (bool, error)

	// Touches returns TRUE if the only points in common between geom1 and geom2 lie in the union of the boundaries of geom1 and geom2.
	// The ouches relation applies to all Area/Area, Line/Line, Line/Area, Point/Area and Point/Line pairs of relationships,
	// but not to the Point/Point pair.
	Touches(A, B space.Geometry) (bool, error)

	// Overlaps returns TRUE if the Geometries "spatially overlap".
	// By that we mean they intersect, but one does not completely contain another.
	Overlaps(A, B space.Geometry) (bool, error)
}
