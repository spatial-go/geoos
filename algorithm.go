// Package geoos provides support for the implementation of spatial operations and geometric algorithms.
package geoos

// Algorithm is the interface implemented by an object that can implementation
// spatial algorithm.
type Algorithm interface {
	Area(geom Geometry) (float64, error)

	Boundary(geom Geometry) (Geometry, error)

	Buffer(geom Geometry, width float64, quadsegs int32) Geometry

	Centroid(geom Geometry) (Geometry, error)

	Contains(geom1, geom2 Geometry) (bool, error)

	ConvexHull(geom Geometry) (Geometry, error)

	CoveredBy(geom1, geom2 Geometry) (bool, error)

	Covers(geom1, geom2 Geometry) (bool, error)

	Crosses(geom1, geom2 Geometry) (bool, error)

	Difference(geom1, geom2 Geometry) (Geometry, error)

	Disjoint(geom1, geom2 Geometry) (bool, error)

	Distance(geom1, geom2 Geometry) (float64, error)

	Envelope(geom Geometry) (Geometry, error)

	Equals(geom1, geom2 Geometry) (bool, error)

	EqualsExact(geom1, geom2 Geometry, tolerance float64) (bool, error)

	HasZ(geom Geometry) (bool, error)

	HausdorffDistance(geom1, geom2 Geometry) (float64, error)

	HausdorffDistanceDensify(s, d Geometry, densifyFrac float64) (float64, error)

	Intersection(geom1, geom2 Geometry) (Geometry, error)

	Intersects(geom1, geom2 Geometry) (bool, error)

	IsClosed(geom Geometry) (bool, error)

	IsEmpty(geom Geometry) (bool, error)

	IsRing(geom Geometry) (bool, error)

	IsSimple(geom Geometry) (bool, error)

	Length(geom Geometry) (float64, error)

	LineMerge(geom Geometry) (Geometry, error)

	NGeometry(geom Geometry) (int, error)

	Overlaps(geom1, geom2 Geometry) (bool, error)

	PointOnSurface(geom Geometry) (Geometry, error)

	Relate(s, d Geometry) (string, error)

	SharedPaths(geom1, geom2 Geometry) (string, error)

	Simplify(geom Geometry, tolerance float64) (Geometry, error)

	SimplifyP(geom Geometry, tolerance float64) (Geometry, error)

	Snap(input, reference Geometry, tolerance float64) (Geometry, error)

	SymDifference(geom1, geom2 Geometry) (Geometry, error)

	Touches(geom1, geom2 Geometry) (bool, error)

	UnaryUnion(geom Geometry) (Geometry, error)

	Union(geom1, geom2 Geometry) (Geometry, error)

	UniquePoints(geom Geometry) (Geometry, error)

	Within(geom1, geom2 Geometry) (bool, error)
}
