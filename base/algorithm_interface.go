package base

import "C"

type Algorithm interface {
	Buffer(g Geometry, width float64, quadsegs int32) Geometry
	// 如果两个几何图形相等，则EqualsExact将返回true，因为它们的点在给定公差内。
	EqualsExact(s Geometry, d Geometry, tolerance float64) bool

	// 其他功能
	// Area returns the area of the geometry, which must be a areal geometry like
	// a polygon or multipolygon.
	Area(g Geometry) (float64, error)

	// Length returns the length of the geometry, which must be a lineal geometry
	// like a linestring or linear ring.
	Length() (float64, error)

	// Distance returns the Cartesian distance between the two geometries.
	Distance(s Geometry, d Geometry) (float64, error)

	// HausdorffDistance computes the maximum distance of the geometry to the nearest
	// point in the other geometry (i.e., considers the whole shape and position of
	// the geometries).
	HausdorffDistance(s Geometry, d Geometry) (float64, error)

	// HausdorffDistanceDensify computes the Hausdorff distance (see
	// HausdorffDistance) with an additional densification fraction amount.
	HausdorffDistanceDensify(s Geometry, d Geometry, densifyFrac float64) (float64, error)

	// Relate computes the intersection matrix (Dimensionally Extended
	// Nine-Intersection Model (DE-9IM) matrix) for the spatial relationship between
	// the two geometries.
	Relate(s Geometry, d Geometry)

	// Unary topology functions

	// Envelope is the bounding box of a geometry, as a polygon.
	Envelope() (*Geometry, error)

	// ConvexHull computes the smallest convex geometry that contains all the points
	// of the geometry.
	ConvexHull() (*Geometry, error)

	// 计算几何图形的组合边界的闭合值
	Boundary(g Geometry) (*Geometry, error)

	// UnaryUnion computes the union of all the constituent geometries of the
	// geometry.
	UnaryUnion() (*Geometry, error)

	// PointOnSurface computes a point geometry guaranteed to be on the surface of
	// the geometry.
	PointOnSurface() (*Geometry, error)

	// Centroid is the center point of the geometry.
	Centroid() (*Geometry, error)

	// LineMerge will merge together a collection of LineStrings where they touch
	// only at their start and end points. The LineStrings must be fully noded. The
	// resulting geometry is a new collection.
	LineMerge() (*Geometry, error)

	// Simplify returns a geometry simplified by amount given by tolerance.
	// May not preserve topology -- see SimplifyP.
	Simplify(tolerance float64) (*Geometry, error)

	// SimplifyP returns a geometry simplified by amount given by tolerance.
	// Unlike Simplify, SimplifyP guarantees it will preserve topology.
	SimplifyP(tolerance float64) (*Geometry, error)

	// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
	UniquePoints() (*Geometry, error)

	// SharedPaths finds paths shared between the two given lineal geometries.
	// Returns a GeometryCollection having two elements:
	//	- first element is a MultiLineString containing shared paths having the _same_ direction on both inputs
	//	- second element is a MultiLineString containing shared paths having the _opposite_ direction on the two inputs
	SharedPaths(other *Geometry) (*Geometry, error)

	// Snap returns a new geometry where the geometry is snapped to the given
	// geometry by given tolerance.
	Snap(other *Geometry, tolerance float64) (*Geometry, error)

	// Binary topology functions

	// Intersection returns a new geometry representing the points shared by this
	// geometry and the other.
	Intersection(other *Geometry) (*Geometry, error)

	// Difference returns a new geometry representing the points making up this
	// geometry that do not make up the other.
	Difference(other *Geometry) (*Geometry, error)

	// SymDifference returns a new geometry representing the set combining the
	// points in this geometry not in the other, and the points in the other
	// geometry and not in this.
	SymDifference(other *Geometry) (*Geometry, error)

	// Union returns a new geometry representing all points in this geometry and the
	// other.
	Union(other *Geometry) (*Geometry, error)

	// Binary predicate functions

	// Disjoint returns true if the two geometries have no point in common.
	Disjoint(other *Geometry) (bool, error)

	// Touches returns true if the two geometries have at least one point in common,
	// but their interiors do not intersect.
	Touches(other *Geometry) (bool, error)

	// Intersects returns true if the two geometries have at least one point in
	// common.
	Intersects(other *Geometry) (bool, error)

	// Crosses returns true if the two geometries have some but not all interior
	// points in common.
	Crosses(other *Geometry) (bool, error)

	// Within returns true if every point of this geometry is a point of the other,
	// and the interiors of the two geometries have at least one point in common.
	Within(other *Geometry) (bool, error)

	// Contains returns true if every point of the other is a point of this geometry,
	// and the interiors of the two geometries have at least one point in common.
	Contains(other *Geometry) (bool, error)

	// Overlaps returns true if the geometries have some but not all points in
	// common, they have the same dimension, and the intersection of the interiors
	// of the two geometries has the same dimension as the geometries themselves.
	Overlaps(other *Geometry) (bool, error)

	// Equals returns true if the two geometries have at least one point in common,
	// and no point of either geometry lies in the exterior of the other geometry.
	Equals(other *Geometry) (bool, error)

	// Covers returns true if every point of the other geometry is a point of this
	// geometry.
	Covers(other *Geometry) (bool, error)

	// CoveredBy returns true if every point of this geometry is a point of the
	// other geometry.
	CoveredBy(other *Geometry) (bool, error)

	// Unary predicate functions

	// IsEmpty returns true if the set of points of this geometry is empty (i.e.,
	// the empty geometry).
	IsEmpty() (bool, error)

	// IsSimple returns true iff the only self-intersections are at boundary points.
	IsSimple() (bool, error)
	// IsRing returns true if the lineal geometry has the ring property.
	IsRing() (bool, error)

	// HasZ returns true if the geometry is 3D.
	HasZ() (bool, error)
	// IsClosed returns true if the geometry is closed (i.e., start & end points
	// equal).
	IsClosed() (bool, error)

	// SRID returns the geometry's SRID, if set.
	SRID() (int, error)

	// SetSRID sets the geometry's SRID.
	SetSRID(srid int)
	// NGeometry returns the number of component geometries (eg., for
	// a collection).
	NGeometry() (int, error)
}
