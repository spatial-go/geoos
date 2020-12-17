package geos

// Algorithm ...
type Algorithm interface {
	Area(g Geometry) (float64, error)

	Buffer(g Geometry, width float64, quadsegs int32) Geometry

	EqualsExact(g1 Geometry, g2 Geometry, tolerance float64) (bool, error)

	IsSimple(g Geometry) (bool, error)

	Length(g Geometry) (float64, error)

	Distance(g1 Geometry, g2 Geometry) (float64, error)

	HausdorffDistance(g1 Geometry, g2 Geometry) (float64, error)

	HausdorffDistanceDensify(s Geometry, d Geometry, densifyFrac float64) (float64, error)

	Relate(s Geometry, d Geometry)

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

	Intersection(other *Geometry) (*Geometry, error)

	Difference(other *Geometry) (*Geometry, error)

	SymDifference(other *Geometry) (*Geometry, error)

	Union(other *Geometry) (*Geometry, error)

	Disjoint(other *Geometry) (bool, error)

	Touches(other *Geometry) (bool, error)

	Intersects(other *Geometry) (bool, error)

	Crosses(g1 Geometry, g2 Geometry) (bool, error)

	Within(g1 Geometry, g2 Geometry) (bool, error)

	Contains(g1 Geometry, g2 Geometry) (bool, error)

	Overlaps(other *Geometry) (bool, error)

	Equals(other *Geometry) (bool, error)

	Covers(other *Geometry) (bool, error)

	CoveredBy(other *Geometry) (bool, error)

	IsEmpty(g Geometry) (bool, error)

	IsRing() (bool, error)

	HasZ() (bool, error)

	IsClosed() (bool, error)

	NGeometry(g Geometry) (int, error)
}
