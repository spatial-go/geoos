// Package planar provides support for the implementation of spatial operations and geometric algorithms.
package planar

import "github.com/spatial-go/geoos/geobase"

// Algorithm is the interface implemented by an object that can implementation
// spatial algorithm.
type Algorithm interface {
	Area(geom geobase.Geometry) (float64, error)

	Boundary(geom geobase.Geometry) (geobase.Geometry, error)

	Buffer(geom geobase.Geometry, width float64, quadsegs int32) geobase.Geometry

	Centroid(geom geobase.Geometry) (geobase.Geometry, error)

	Contains(geom1, geom2 geobase.Geometry) (bool, error)

	ConvexHull(geom geobase.Geometry) (geobase.Geometry, error)

	CoveredBy(geom1, geom2 geobase.Geometry) (bool, error)

	Covers(geom1, geom2 geobase.Geometry) (bool, error)

	Crosses(geom1, geom2 geobase.Geometry) (bool, error)

	Difference(geom1, geom2 geobase.Geometry) (geobase.Geometry, error)

	Disjoint(geom1, geom2 geobase.Geometry) (bool, error)

	Distance(geom1, geom2 geobase.Geometry) (float64, error)

	SphericalDistance(point1, point2 geobase.Point) float64

	Envelope(geom geobase.Geometry) (geobase.Geometry, error)

	Equals(geom1, geom2 geobase.Geometry) (bool, error)

	EqualsExact(geom1, geom2 geobase.Geometry, tolerance float64) (bool, error)

	HasZ(geom geobase.Geometry) (bool, error)

	HausdorffDistance(geom1, geom2 geobase.Geometry) (float64, error)

	HausdorffDistanceDensify(s, d geobase.Geometry, densifyFrac float64) (float64, error)

	Intersection(geom1, geom2 geobase.Geometry) (geobase.Geometry, error)

	Intersects(geom1, geom2 geobase.Geometry) (bool, error)

	IsClosed(geom geobase.Geometry) (bool, error)

	IsEmpty(geom geobase.Geometry) (bool, error)

	IsRing(geom geobase.Geometry) (bool, error)

	IsSimple(geom geobase.Geometry) (bool, error)

	Length(geom geobase.Geometry) (float64, error)

	LineMerge(geom geobase.Geometry) (geobase.Geometry, error)

	NGeometry(geom geobase.Geometry) (int, error)

	Overlaps(geom1, geom2 geobase.Geometry) (bool, error)

	PointOnSurface(geom geobase.Geometry) (geobase.Geometry, error)

	Relate(s, d geobase.Geometry) (string, error)

	SharedPaths(geom1, geom2 geobase.Geometry) (string, error)

	Simplify(geom geobase.Geometry, tolerance float64) (geobase.Geometry, error)

	SimplifyP(geom geobase.Geometry, tolerance float64) (geobase.Geometry, error)

	Snap(input, reference geobase.Geometry, tolerance float64) (geobase.Geometry, error)

	SymDifference(geom1, geom2 geobase.Geometry) (geobase.Geometry, error)

	Touches(geom1, geom2 geobase.Geometry) (bool, error)

	UnaryUnion(geom geobase.Geometry) (geobase.Geometry, error)

	Union(geom1, geom2 geobase.Geometry) (geobase.Geometry, error)

	UniquePoints(geom geobase.Geometry) (geobase.Geometry, error)

	Within(geom1, geom2 geobase.Geometry) (bool, error)
}
