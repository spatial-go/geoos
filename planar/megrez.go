package planar

import (
	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/geo"
)

// MegrezAlgorithm algorithm implement
type MegrezAlgorithm struct{}

// Area returns the area of a polygonal geometry.
func (g *MegrezAlgorithm) Area(geom geoos.Geometry) (float64, error) {
	return geo.Area(wkt.MarshalString(geom))
}

// Boundary returns the closure of the combinatorial boundary of this geoos.Geometry.
func (g *MegrezAlgorithm) Boundary(geom geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).Boundary(geom)
}

// Buffer sReturns a geometry that represents all points whose distance
// from this geoos.Geometry is less than or equal to distance.
func (g *MegrezAlgorithm) Buffer(geom geoos.Geometry, width float64, quadsegs int32) (geometry geoos.Geometry) {
	return GetStrategy(newGEOAlgorithm).Buffer(geom, width, quadsegs)
}

// Centroid  computes the geometric center of a geometry, or equivalently, the center of mass of the geometry as a POINT.
// For [MULTI]POINTs, this is computed as the arithmetic mean of the input coordinates.
// For [MULTI]LINESTRINGs, this is computed as the weighted length of each line segment.
// For [MULTI]POLYGONs, "weight" is thought in terms of area.
// If an empty geometry is supplied, an empty GEOMETRYCOLLECTION is returned.
// If NULL is supplied, NULL is returned.
// If CIRCULARSTRING or COMPOUNDCURVE are supplied, they are converted to linestring wtih CurveToLine first,
// then same than for LINESTRING
func (g *MegrezAlgorithm) Centroid(geom geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).Centroid(geom)
}

// Contains geoos.Geometry A contains geoos.Geometry B if and only if no points of B lie in the exterior of A,
// and at least one point of the interior of B lies in the interior of A.
// An important subtlety of this definition is that A does not contain its boundary, but A does contain itself.
// Returns TRUE if geometry B is completely inside geometry A.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (g *MegrezAlgorithm) Contains(geom1, geom2 geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).Contains(geom1, geom2)
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (g *MegrezAlgorithm) ConvexHull(geom geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).ConvexHull(geom)
}

// CoveredBy returns TRUE if no point in geoos.Geometry A is outside geoos.Geometry B
func (g *MegrezAlgorithm) CoveredBy(geom1, geom2 geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).CoveredBy(geom1, geom2)
}

// Covers returns TRUE if no point in geoos.Geometry B is outside geoos.Geometry A
func (g *MegrezAlgorithm) Covers(geom1, geom2 geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).Covers(geom1, geom2)
}

// Crosses takes two geometry objects and returns TRUE if their intersection "spatially cross",
// that is, the geometries have some, but not all interior points in common.
// The intersection of the interiors of the geometries must not be the empty set
// and must have a dimensionality less than the maximum dimension of the two input geometries.
// Additionally, the intersection of the two geometries must not equal either of the source geometries.
// Otherwise, it returns FALSE.
func (g *MegrezAlgorithm) Crosses(geom1, geom2 geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).Crosses(geom1, geom2)
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (g *MegrezAlgorithm) Difference(geom1, geom2 geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).Difference(geom1, geom2)
}

// Disjoint Overlaps, Touches, Within all imply geometries are not spatially disjoint.
// If any of the aforementioned returns true, then the geometries are not spatially disjoint.
// Disjoint implies false for spatial intersection.
func (g *MegrezAlgorithm) Disjoint(geom1, geom2 geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).Disjoint(geom1, geom2)
}

// Distance returns the minimum 2D Cartesian (planar) distance between two geometries, in projected units (spatial ref units).
func (g *MegrezAlgorithm) Distance(geom1, geom2 geoos.Geometry) (float64, error) {
	return GetStrategy(newGEOAlgorithm).Distance(geom1, geom2)
}

// SphericalDistance calculates spherical distance
//
// To get real distance in km
func (g *MegrezAlgorithm) SphericalDistance(point1, point2 geoos.Point) float64 {
	return GetStrategy(newGEOAlgorithm).SphericalDistance(point1, point2)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (g *MegrezAlgorithm) Envelope(geom geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).Envelope(geom)
}

// Equals returns TRUE if the given Geometries are "spatially equal".
func (g *MegrezAlgorithm) Equals(geom1, geom2 geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).Equals(geom1, geom2)
}

// EqualsExact returns true if both geometries are Equal, as evaluated by their
// points being within the given tolerance.
func (g *MegrezAlgorithm) EqualsExact(geom1, geom2 geoos.Geometry, tolerance float64) (bool, error) {
	return GetStrategy(newGEOAlgorithm).EqualsExact(geom1, geom2, tolerance)
}

// HasZ returns true if the geometry is 3D
func (g *MegrezAlgorithm) HasZ(geom geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).HasZ(geom)
}

// HausdorffDistance returns the Hausdorff distance between two geometries, a measure of how similar
// or dissimilar 2 geometries are. Implements algorithm for computing a distance metric which can be
// thought of as the "Discrete Hausdorff Distance". This is the Hausdorff distance restricted
// to discrete points for one of the geometries
func (g *MegrezAlgorithm) HausdorffDistance(geom1, geom2 geoos.Geometry) (float64, error) {
	return GetStrategy(newGEOAlgorithm).HausdorffDistance(geom1, geom2)
}

// HausdorffDistanceDensify computes the Hausdorff distance with an additional densification fraction amount
func (g *MegrezAlgorithm) HausdorffDistanceDensify(s, d geoos.Geometry, densifyFrac float64) (float64, error) {
	return GetStrategy(newGEOAlgorithm).HausdorffDistanceDensify(s, d, densifyFrac)
}

// Intersection returns a geometry that represents the point set intersection of the Geometries.
func (g *MegrezAlgorithm) Intersection(geom1, geom2 geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).Intersection(geom1, geom2)
}

// Intersects If a geometry  shares any portion of space then they intersect
func (g *MegrezAlgorithm) Intersects(geom1, geom2 geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).Intersects(geom1, geom2)
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or volumetric (closed).
func (g *MegrezAlgorithm) IsClosed(geom geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).IsClosed(geom)
}

// IsEmpty returns true if this geoos.Geometry is an empty geometry.
// If true, then this geoos.Geometry represents an empty geometry collection, polygon, point etc.
func (g *MegrezAlgorithm) IsEmpty(geom geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).IsEmpty(geom)
}

// IsRing returns true if the lineal geometry has the ring property.
func (g *MegrezAlgorithm) IsRing(geom geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).IsRing(geom)
}

// IsSimple returns true if this geoos.Geometry has no anomalous geometric points, such as self intersection or self tangency.
func (g *MegrezAlgorithm) IsSimple(geom geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).IsSimple(geom)
}

// Length returns the 2D Cartesian length of the geometry if it is a LineString, MultiLineString
func (g *MegrezAlgorithm) Length(geom geoos.Geometry) (float64, error) {
	return GetStrategy(newGEOAlgorithm).Length(geom)
}

// LineMerge returns a (set of) LineString(s) formed by sewing together the constituent line work of a MULTILINESTRING.
func (g *MegrezAlgorithm) LineMerge(geom geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).LineMerge(geom)
}

// NGeometry returns the number of component geometries.
func (g *MegrezAlgorithm) NGeometry(geom geoos.Geometry) (int, error) {
	return GetStrategy(newGEOAlgorithm).NGeometry(geom)
}

// Overlaps returns TRUE if the Geometries "spatially overlap".
// By that we mean they intersect, but one does not completely contain another.
func (g *MegrezAlgorithm) Overlaps(geom1, geom2 geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).Overlaps(geom1, geom2)
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (g *MegrezAlgorithm) PointOnSurface(geom geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).PointOnSurface(geom)
}

// Relate computes the intersection matrix (Dimensionally Extended
// Nine-Intersection Model (DE-9IM) matrix) for the spatial relationship between
// the two geometries.
func (g *MegrezAlgorithm) Relate(s, d geoos.Geometry) (string, error) {
	return GetStrategy(newGEOAlgorithm).Relate(s, d)
}

// SharedPaths returns a collection containing paths shared by the two input geometries.
// Those going in the same direction are in the first element of the collection,
// those going in the opposite direction are in the second element.
// The paths themselves are given in the direction of the first geometry.
func (g *MegrezAlgorithm) SharedPaths(geom1, geom2 geoos.Geometry) (string, error) {
	return GetStrategy(newGEOAlgorithm).SharedPaths(geom1, geom2)
}

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (g *MegrezAlgorithm) Simplify(geom geoos.Geometry, tolerance float64) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).Simplify(geom, tolerance)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (g *MegrezAlgorithm) SimplifyP(geom geoos.Geometry, tolerance float64) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).SimplifyP(geom, tolerance)
}

// Snap the vertices and segments of a geometry to another geoos.Geometry's vertices.
// A snap distance tolerance is used to control where snapping is performed.
// The result geometry is the input geometry with the vertices snapped.
// If no snapping occurs then the input geometry is returned unchanged.
func (g *MegrezAlgorithm) Snap(input, reference geoos.Geometry, tolerance float64) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).Snap(input, reference, tolerance)
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (g *MegrezAlgorithm) SymDifference(geom1, geom2 geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).SymDifference(geom1, geom2)
}

// Touches returns TRUE if the only points in common between geom1 and geom2 lie in the union of the boundaries of geom1 and geom2.
// The ouches relation applies to all Area/Area, Line/Line, Line/Area, Point/Area and Point/Line pairs of relationships,
// but not to the Point/Point pair.
func (g *MegrezAlgorithm) Touches(geom1, geom2 geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).Touches(geom1, geom2)
}

// UnaryUnion does dissolve boundaries between components of a multipolygon (invalid) and does perform union
// between the components of a geometrycollection
func (g *MegrezAlgorithm) UnaryUnion(geom geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).UnaryUnion(geom)
}

// Union returns a new geometry representing all points in this geometry and the other.
func (g *MegrezAlgorithm) Union(geom1, geom2 geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).Union(geom1, geom2)
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (g *MegrezAlgorithm) UniquePoints(geom geoos.Geometry) (geoos.Geometry, error) {
	return GetStrategy(newGEOAlgorithm).UniquePoints(geom)
}

// Within returns TRUE if geometry A is completely inside geometry B.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (g *MegrezAlgorithm) Within(geom1, geom2 geoos.Geometry) (bool, error) {
	return GetStrategy(newGEOAlgorithm).Within(geom1, geom2)
}
