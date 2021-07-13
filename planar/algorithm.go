// Package planar provides support for the implementation of spatial operations and geometric algorithms.
package planar

import (
	"errors"

	"github.com/spatial-go/geoos/space"
)

// ErrNotPolygon UnaryUnion parameter is not polygon
var ErrNotPolygon = errors.New("Geometry is not polygon")

// Algorithm is the interface implemented by an object that can implementation
// spatial algorithm.
type Algorithm interface {
	Area(geom space.Geometry) (float64, error)

	Boundary(geom space.Geometry) (space.Geometry, error)

	Buffer(geom space.Geometry, width float64, quadsegs int32) space.Geometry

	Centroid(geom space.Geometry) (space.Geometry, error)

	Contains(geom1, geom2 space.Geometry) (bool, error)

	ConvexHull(geom space.Geometry) (space.Geometry, error)

	CoveredBy(geom1, geom2 space.Geometry) (bool, error)

	Covers(geom1, geom2 space.Geometry) (bool, error)

	Crosses(geom1, geom2 space.Geometry) (bool, error)

	Difference(geom1, geom2 space.Geometry) (space.Geometry, error)

	Disjoint(geom1, geom2 space.Geometry) (bool, error)

	Distance(geom1, geom2 space.Geometry) (float64, error)

	SphericalDistance(geom1, geom2 space.Geometry) (float64, error)

	Envelope(geom space.Geometry) (space.Geometry, error)

	Equals(geom1, geom2 space.Geometry) (bool, error)

	EqualsExact(geom1, geom2 space.Geometry, tolerance float64) (bool, error)

	HausdorffDistance(geom1, geom2 space.Geometry) (float64, error)

	HausdorffDistanceDensify(s, d space.Geometry, densifyFrac float64) (float64, error)

	Intersection(geom1, geom2 space.Geometry) (space.Geometry, error)

	Intersects(geom1, geom2 space.Geometry) (bool, error)

	IsClosed(geom space.Geometry) (bool, error)

	IsEmpty(geom space.Geometry) (bool, error)

	IsRing(geom space.Geometry) (bool, error)

	IsSimple(geom space.Geometry) (bool, error)

	Length(geom space.Geometry) (float64, error)

	LineMerge(geom space.Geometry) (space.Geometry, error)

	NGeometry(geom space.Geometry) (int, error)

	Overlaps(geom1, geom2 space.Geometry) (bool, error)

	PointOnSurface(geom space.Geometry) (space.Geometry, error)

	Relate(s, d space.Geometry) (string, error)

	SharedPaths(geom1, geom2 space.Geometry) (string, error)

	Simplify(geom space.Geometry, tolerance float64) (space.Geometry, error)

	SimplifyP(geom space.Geometry, tolerance float64) (space.Geometry, error)

	Snap(input, reference space.Geometry, tolerance float64) (space.Geometry, error)

	SymDifference(geom1, geom2 space.Geometry) (space.Geometry, error)

	Touches(geom1, geom2 space.Geometry) (bool, error)

	UnaryUnion(geom space.Geometry) (space.Geometry, error)

	Union(geom1, geom2 space.Geometry) (space.Geometry, error)

	UniquePoints(geom space.Geometry) (space.Geometry, error)

	Within(geom1, geom2 space.Geometry) (bool, error)
}
