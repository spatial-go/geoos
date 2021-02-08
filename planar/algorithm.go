// Package planar provides support for the implementation of spatial operations and geometric algorithms.
package planar

import (
	"github.com/spatial-go/geoos"
)

// Algorithm is the interface implemented by an object that can implementation
// spatial algorithm.
type Algorithm interface {
	Area(geom geoos.Geometry) (float64, error)

	Boundary(geom geoos.Geometry) (geoos.Geometry, error)

	Buffer(geom geoos.Geometry, width float64, quadsegs int32) geoos.Geometry

	Centroid(geom geoos.Geometry) (geoos.Geometry, error)

	Contains(geom1, geom2 geoos.Geometry) (bool, error)

	ConvexHull(geom geoos.Geometry) (geoos.Geometry, error)

	CoveredBy(geom1, geom2 geoos.Geometry) (bool, error)

	Covers(geom1, geom2 geoos.Geometry) (bool, error)

	Crosses(geom1, geom2 geoos.Geometry) (bool, error)

	Difference(geom1, geom2 geoos.Geometry) (geoos.Geometry, error)

	Disjoint(geom1, geom2 geoos.Geometry) (bool, error)

	Distance(geom1, geom2 geoos.Geometry) (float64, error)

	SphericalDistance(point1, point2 geoos.Point) float64

	Envelope(geom geoos.Geometry) (geoos.Geometry, error)

	Equals(geom1, geom2 geoos.Geometry) (bool, error)

	EqualsExact(geom1, geom2 geoos.Geometry, tolerance float64) (bool, error)

	HasZ(geom geoos.Geometry) (bool, error)

	HausdorffDistance(geom1, geom2 geoos.Geometry) (float64, error)

	HausdorffDistanceDensify(s, d geoos.Geometry, densifyFrac float64) (float64, error)

	Intersection(geom1, geom2 geoos.Geometry) (geoos.Geometry, error)

	Intersects(geom1, geom2 geoos.Geometry) (bool, error)

	IsClosed(geom geoos.Geometry) (bool, error)

	IsEmpty(geom geoos.Geometry) (bool, error)

	IsRing(geom geoos.Geometry) (bool, error)

	IsSimple(geom geoos.Geometry) (bool, error)

	Length(geom geoos.Geometry) (float64, error)

	LineMerge(geom geoos.Geometry) (geoos.Geometry, error)

	NGeometry(geom geoos.Geometry) (int, error)

	Overlaps(geom1, geom2 geoos.Geometry) (bool, error)

	PointOnSurface(geom geoos.Geometry) (geoos.Geometry, error)

	Relate(s, d geoos.Geometry) (string, error)

	SharedPaths(geom1, geom2 geoos.Geometry) (string, error)

	Simplify(geom geoos.Geometry, tolerance float64) (geoos.Geometry, error)

	SimplifyP(geom geoos.Geometry, tolerance float64) (geoos.Geometry, error)

	Snap(input, reference geoos.Geometry, tolerance float64) (geoos.Geometry, error)

	SymDifference(geom1, geom2 geoos.Geometry) (geoos.Geometry, error)

	Touches(geom1, geom2 geoos.Geometry) (bool, error)

	UnaryUnion(geom geoos.Geometry) (geoos.Geometry, error)

	Union(geom1, geom2 geoos.Geometry) (geoos.Geometry, error)

	UniquePoints(geom geoos.Geometry) (geoos.Geometry, error)

	Within(geom1, geom2 geoos.Geometry) (bool, error)
}
