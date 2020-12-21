// Package geoos provides support for the implementation of spatial operations and geometric algorithms.
package geoos

// Algorithm is the interface implemented by an object that can implementation
// spatial algorithm.
type Algorithm interface {
	Area(geom Geometry) (float64, error)

	Buffer(geom Geometry, width float64, quadsegs int32) Geometry

	EqualsExact(geom1, geom2 Geometry, tolerance float64) (bool, error)

	IsSimple(geom Geometry) (bool, error)

	Length(geom Geometry) (float64, error)

	Distance(geom1, geom2 Geometry) (float64, error)

	HausdorffDistance(geom1, geom2 Geometry) (float64, error)

	HausdorffDistanceDensify(s, d Geometry, densifyFrac float64) (float64, error)

	Relate(s, d Geometry) (string, error)

	Envelope(geom Geometry) (Geometry, error)

	ConvexHull(geom Geometry) (Geometry, error)

	Boundary(geom Geometry) (Geometry, error)

	UnaryUnion(geom Geometry) (Geometry, error)

	PointOnSurface(geom Geometry) (Geometry, error)

	Centroid(geom Geometry) (Geometry, error)

	LineMerge(geom Geometry) (Geometry, error)

	Simplify(geom Geometry, tolerance float64) (Geometry, error)

	SimplifyP(geom Geometry, tolerance float64) (Geometry, error)

	UniquePoints(geom Geometry) (Geometry, error)

	SharedPaths(geom1, geom2 Geometry) (string, error)

	Snap(input, reference Geometry, tolerance float64) (Geometry, error)

	Intersection(geom1, geom2 Geometry) (Geometry, error)

	Difference(geom1, geom2 Geometry) (Geometry, error)

	SymDifference(geom1, geom2 Geometry) (Geometry, error)

	Union(geom1, geom2 Geometry) (Geometry, error)

	Disjoint(geom1, geom2 Geometry) (bool, error)

	Touches(geom1, geom2 Geometry) (bool, error)

	Intersects(geom1, geom2 Geometry) (bool, error)

	Crosses(geom1, geom2 Geometry) (bool, error)

	Within(geom1, geom2 Geometry) (bool, error)

	Contains(geom1, geom2 Geometry) (bool, error)

	Overlaps(geom1, geom2 Geometry) (bool, error)

	Equals(geom1, geom2 Geometry) (bool, error)

	Covers(geom1, geom2 Geometry) (bool, error)

	CoveredBy(geom1, geom2 Geometry) (bool, error)

	IsEmpty(geom Geometry) (bool, error)

	IsRing(geom Geometry) (bool, error)

	HasZ(geom Geometry) (bool, error)

	IsClosed(geom Geometry) (bool, error)

	NGeometry(geom Geometry) (int, error)
}
