package geos

// Algorithm is the interface implemented by an object that can implementation
// spatial algorithm.
type Algorithm interface {
	Area(g Geometry) (float64, error)

	Buffer(g Geometry, width float64, quadsegs int32) Geometry

	EqualsExact(g1 Geometry, g2 Geometry, tolerance float64) (bool, error)

	IsSimple(g Geometry) (bool, error)

	Length(g Geometry) (float64, error)

	Distance(g1 Geometry, g2 Geometry) (float64, error)

	HausdorffDistance(g1 Geometry, g2 Geometry) (float64, error)

	HausdorffDistanceDensify(s Geometry, d Geometry, densifyFrac float64) (float64, error)

	Relate(s Geometry, d Geometry) (string, error)

	Envelope(g Geometry) (Geometry, error)

	ConvexHull(g Geometry) (Geometry, error)

	Boundary(g Geometry) (Geometry, error)

	UnaryUnion(g Geometry) (Geometry, error)

	PointOnSurface(g Geometry) (Geometry, error)

	Centroid(g Geometry) (Geometry, error)

	LineMerge(g Geometry) (Geometry, error)

	Simplify(g Geometry, tolerance float64) (Geometry, error)

	SimplifyP(g Geometry, tolerance float64) (Geometry, error)

	UniquePoints(g Geometry) (Geometry, error)

	SharedPaths(g1 Geometry, g2 Geometry) (string, error)

	Snap(input Geometry, reference Geometry, tolerance float64) (Geometry, error)

	Intersection(g1 Geometry, g2 Geometry) (Geometry, error)

	Difference(g1 Geometry, g2 Geometry) (Geometry, error)

	SymDifference(g1 Geometry, g2 Geometry) (Geometry, error)

	Union(g1 Geometry, g2 Geometry) (Geometry, error)

	Disjoint(g1 Geometry, g2 Geometry) (bool, error)

	Touches(g1 Geometry, g2 Geometry) (bool, error)

	Intersects(g1 Geometry, g2 Geometry) (bool, error)

	Crosses(g1 Geometry, g2 Geometry) (bool, error)

	Within(g1 Geometry, g2 Geometry) (bool, error)

	Contains(g1 Geometry, g2 Geometry) (bool, error)

	Overlaps(g1 Geometry, g2 Geometry) (bool, error)

	Equals(g1 Geometry, g2 Geometry) (bool, error)

	Covers(g1 Geometry, g2 Geometry) (bool, error)

	CoveredBy(g1 Geometry, g2 Geometry) (bool, error)

	IsEmpty(g Geometry) (bool, error)

	IsRing(g Geometry) (bool, error)

	HasZ(g Geometry) (bool, error)

	IsClosed(g Geometry) (bool, error)

	NGeometry(g Geometry) (int, error)
}
