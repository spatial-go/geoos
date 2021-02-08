package planar

import (
	"math"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/common"
	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/geo"
)

// GEOAlgorithm algorithm implement by geos
type GEOAlgorithm struct{}

// Area returns the area of a polygonal geometry.
func (g *GEOAlgorithm) Area(geom geoos.Geometry) (float64, error) {
	return geo.Area(wkt.MarshalString(geom))
}

// Boundary returns the closure of the combinatorial boundary of this geoos.Geometry.
func (g *GEOAlgorithm) Boundary(geom geoos.Geometry) (geoos.Geometry, error) {
	aa := wkt.MarshalString(geom)
	boundary, err := geo.Boundary(aa)
	if err != nil {
		return nil, err
	}
	geometry, err := wkt.UnmarshalString(boundary)
	if err != nil {
		return nil, err
	}
	return geometry, nil
}

// Buffer sReturns a geometry that represents all points whose distance
// from this geoos.Geometry is less than or equal to distance.
func (g *GEOAlgorithm) Buffer(geom geoos.Geometry, width float64, quadsegs int32) (geometry geoos.Geometry) {
	result, err := geo.Buffer(wkt.MarshalString(geom), width, quadsegs)
	if err != nil {
		return
	}
	geometry, _ = wkt.UnmarshalString(result)
	return
}

// Centroid  computes the geometric center of a geometry, or equivalently, the center of mass of the geometry as a POINT.
// For [MULTI]POINTs, this is computed as the arithmetic mean of the input coordinates.
// For [MULTI]LINESTRINGs, this is computed as the weighted length of each line segment.
// For [MULTI]POLYGONs, "weight" is thought in terms of area.
// If an empty geometry is supplied, an empty GEOMETRYCOLLECTION is returned.
// If NULL is supplied, NULL is returned.
// If CIRCULARSTRING or COMPOUNDCURVE are supplied, they are converted to linestring wtih CurveToLine first,
// then same than for LINESTRING
func (g *GEOAlgorithm) Centroid(geom geoos.Geometry) (geoos.Geometry, error) {
	result, err := geo.Centroid(wkt.MarshalString(geom))
	if err != nil {
		return nil, err
	}
	geometry, err := wkt.UnmarshalString(result)
	if err != nil {
		return nil, err
	}
	return geometry, nil
}

// Contains geoos.Geometry A contains geoos.Geometry B if and only if no points of B lie in the exterior of A,
// and at least one point of the interior of B lies in the interior of A.
// An important subtlety of this definition is that A does not contain its boundary, but A does contain itself.
// Returns TRUE if geometry B is completely inside geometry A.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (g *GEOAlgorithm) Contains(geom1, geom2 geoos.Geometry) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.Contains(ms1, ms2)
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (g *GEOAlgorithm) ConvexHull(geom geoos.Geometry) (geoos.Geometry, error) {
	result, err := geo.ConvexHull(wkt.MarshalString(geom))
	if err != nil {
		return nil, err
	}
	return wkt.UnmarshalString(result)
}

// CoveredBy returns TRUE if no point in geoos.Geometry A is outside geoos.Geometry B
func (g *GEOAlgorithm) CoveredBy(geom1, geom2 geoos.Geometry) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.CoversBy(ms1, ms2)
}

// Covers returns TRUE if no point in geoos.Geometry B is outside geoos.Geometry A
func (g *GEOAlgorithm) Covers(geom1, geom2 geoos.Geometry) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.Covers(ms1, ms2)
}

// Crosses takes two geometry objects and returns TRUE if their intersection "spatially cross",
// that is, the geometries have some, but not all interior points in common.
// The intersection of the interiors of the geometries must not be the empty set
// and must have a dimensionality less than the maximum dimension of the two input geometries.
// Additionally, the intersection of the two geometries must not equal either of the source geometries.
// Otherwise, it returns FALSE.
func (g *GEOAlgorithm) Crosses(geom1, geom2 geoos.Geometry) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.Crosses(ms1, ms2)
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (g *GEOAlgorithm) Difference(geom1, geom2 geoos.Geometry) (geoos.Geometry, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	result, err := geo.Difference(ms1, ms2)
	if err != nil {
		return nil, err
	}
	geometry, err := wkt.UnmarshalString(result)
	if err != nil {
		return nil, err
	}
	return geometry, nil
}

// Disjoint Overlaps, Touches, Within all imply geometries are not spatially disjoint.
// If any of the aforementioned returns true, then the geometries are not spatially disjoint.
// Disjoint implies false for spatial intersection.
func (g *GEOAlgorithm) Disjoint(geom1, geom2 geoos.Geometry) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.Disjoint(ms1, ms2)
}

// Distance returns the minimum 2D Cartesian (planar) distance between two geometries, in projected units (spatial ref units).
func (g *GEOAlgorithm) Distance(geom1, geom2 geoos.Geometry) (float64, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.Distance(ms1, ms2)
}

// SphericalDistance calculates spherical distance
//
// To get real distance in km
func (g *GEOAlgorithm) SphericalDistance(point1, point2 geoos.Point) float64 {
	radius := common.EarthR
	rad := common.DegreeRad
	lat0 := point1[1] * rad
	lng0 := point1[0] * rad
	lat1 := point2[1] * rad
	lng1 := point2[0] * rad
	theta := lng1 - lng0
	dist := math.Acos(math.Sin(lat0)*math.Sin(lat1) + math.Cos(lat0)*math.Cos(lat1)*math.Cos(theta))
	return dist * radius
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (g *GEOAlgorithm) Envelope(geom geoos.Geometry) (geoos.Geometry, error) {
	result, err := geo.Envelope(wkt.MarshalString(geom))
	if err != nil {
		return nil, err
	}
	return wkt.UnmarshalString(result)
}

// Equals returns TRUE if the given Geometries are "spatially equal".
func (g *GEOAlgorithm) Equals(geom1, geom2 geoos.Geometry) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.Equals(ms1, ms2)
}

// EqualsExact returns true if both geometries are Equal, as evaluated by their
// points being within the given tolerance.
func (g *GEOAlgorithm) EqualsExact(geom1, geom2 geoos.Geometry, tolerance float64) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.EqualsExact(ms1, ms2, tolerance)
}

// HasZ returns true if the geometry is 3D
func (g *GEOAlgorithm) HasZ(geom geoos.Geometry) (bool, error) {
	return geo.HasZ(wkt.MarshalString(geom))
}

// HausdorffDistance returns the Hausdorff distance between two geometries, a measure of how similar
// or dissimilar 2 geometries are. Implements algorithm for computing a distance metric which can be
// thought of as the "Discrete Hausdorff Distance". This is the Hausdorff distance restricted
// to discrete points for one of the geometries
func (g *GEOAlgorithm) HausdorffDistance(geom1, geom2 geoos.Geometry) (float64, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.HausdorffDistance(ms1, ms2)
}

// HausdorffDistanceDensify computes the Hausdorff distance with an additional densification fraction amount
func (g *GEOAlgorithm) HausdorffDistanceDensify(s, d geoos.Geometry, densifyFrac float64) (float64, error) {
	ms1, ms2 := convertGeomToWKT(s, d)
	return geo.HausdorffDistanceDensify(ms1, ms2, densifyFrac)
}

// Intersection returns a geometry that represents the point set intersection of the Geometries.
func (g *GEOAlgorithm) Intersection(geom1, geom2 geoos.Geometry) (geoos.Geometry, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	result, err := geo.Intersection(ms1, ms2)
	if err != nil {
		return nil, err
	}
	geometry, err := wkt.UnmarshalString(result)
	if err != nil {
		return nil, err
	}
	return geometry, nil
}

// Intersects If a geometry  shares any portion of space then they intersect
func (g *GEOAlgorithm) Intersects(geom1, geom2 geoos.Geometry) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.Intersects(ms1, ms2)
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or volumetric (closed).
func (g *GEOAlgorithm) IsClosed(geom geoos.Geometry) (bool, error) {
	return geo.IsClosed(wkt.MarshalString(geom))
}

// IsEmpty returns true if this geoos.Geometry is an empty geometry.
// If true, then this geoos.Geometry represents an empty geometry collection, polygon, point etc.
func (g *GEOAlgorithm) IsEmpty(geom geoos.Geometry) (bool, error) {
	return geo.IsEmpty(wkt.MarshalString(geom))
}

// IsRing returns true if the lineal geometry has the ring property.
func (g *GEOAlgorithm) IsRing(geom geoos.Geometry) (bool, error) {
	return geo.IsRing(wkt.MarshalString(geom))

}

// IsSimple returns true if this geoos.Geometry has no anomalous geometric points, such as self intersection or self tangency.
func (g *GEOAlgorithm) IsSimple(geom geoos.Geometry) (bool, error) {
	return geo.IsSimple(wkt.MarshalString(geom))
}

// Length returns the 2D Cartesian length of the geometry if it is a LineString, MultiLineString
func (g *GEOAlgorithm) Length(geom geoos.Geometry) (float64, error) {
	return geo.Length(wkt.MarshalString(geom))
}

// LineMerge returns a (set of) LineString(s) formed by sewing together the constituent line work of a MULTILINESTRING.
func (g *GEOAlgorithm) LineMerge(geom geoos.Geometry) (geoos.Geometry, error) {
	result, err := geo.LineMerge(wkt.MarshalString(geom))
	if err != nil {
		return nil, err
	}
	return wkt.UnmarshalString(result)
}

// NGeometry returns the number of component geometries.
func (g *GEOAlgorithm) NGeometry(geom geoos.Geometry) (int, error) {
	return geo.NGeometry(wkt.MarshalString(geom))
}

// Overlaps returns TRUE if the Geometries "spatially overlap".
// By that we mean they intersect, but one does not completely contain another.
func (g *GEOAlgorithm) Overlaps(geom1, geom2 geoos.Geometry) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.Overlaps(ms1, ms2)
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (g *GEOAlgorithm) PointOnSurface(geom geoos.Geometry) (geoos.Geometry, error) {
	result, err := geo.PointOnSurface(wkt.MarshalString(geom))
	if err != nil {
		return nil, err
	}
	return wkt.UnmarshalString(result)
}

// Relate computes the intersection matrix (Dimensionally Extended
// Nine-Intersection Model (DE-9IM) matrix) for the spatial relationship between
// the two geometries.
func (g *GEOAlgorithm) Relate(s, d geoos.Geometry) (string, error) {
	ms1, ms2 := convertGeomToWKT(s, d)
	return geo.Relate(ms1, ms2)
}

// SharedPaths returns a collection containing paths shared by the two input geometries.
// Those going in the same direction are in the first element of the collection,
// those going in the opposite direction are in the second element.
// The paths themselves are given in the direction of the first geometry.
func (g *GEOAlgorithm) SharedPaths(geom1, geom2 geoos.Geometry) (string, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	result, err := geo.SharedPaths(ms1, ms2)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (g *GEOAlgorithm) Simplify(geom geoos.Geometry, tolerance float64) (geoos.Geometry, error) {
	result, err := geo.Simplify(wkt.MarshalString(geom), tolerance)
	if err != nil {
		return nil, err
	}
	return wkt.UnmarshalString(result)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (g *GEOAlgorithm) SimplifyP(geom geoos.Geometry, tolerance float64) (geoos.Geometry, error) {
	result, err := geo.SimplifyP(wkt.MarshalString(geom), tolerance)
	if err != nil {
		return nil, err
	}
	return wkt.UnmarshalString(result)
}

// Snap the vertices and segments of a geometry to another geoos.Geometry's vertices.
// A snap distance tolerance is used to control where snapping is performed.
// The result geometry is the input geometry with the vertices snapped.
// If no snapping occurs then the input geometry is returned unchanged.
func (g *GEOAlgorithm) Snap(input, reference geoos.Geometry, tolerance float64) (geoos.Geometry, error) {
	inGeom := wkt.MarshalString(input)
	refGeom := wkt.MarshalString(reference)
	result, err := geo.Snap(inGeom, refGeom, tolerance)
	if err != nil {
		return nil, err
	}
	geometry, err := wkt.UnmarshalString(result)
	if err != nil {
		return nil, err
	}
	return geometry, nil
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (g *GEOAlgorithm) SymDifference(geom1, geom2 geoos.Geometry) (geoos.Geometry, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	result, err := geo.SymDifference(ms1, ms2)
	if err != nil {
		return nil, err
	}
	geometry, err := wkt.UnmarshalString(result)
	if err != nil {
		return nil, err
	}
	return geometry, nil
}

// Touches returns TRUE if the only points in common between geom1 and geom2 lie in the union of the boundaries of geom1 and geom2.
// The ouches relation applies to all Area/Area, Line/Line, Line/Area, Point/Area and Point/Line pairs of relationships,
// but not to the Point/Point pair.
func (g *GEOAlgorithm) Touches(geom1, geom2 geoos.Geometry) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.Touches(ms1, ms2)
}

// UnaryUnion does dissolve boundaries between components of a multipolygon (invalid) and does perform union
// between the components of a geometrycollection
func (g *GEOAlgorithm) UnaryUnion(geom geoos.Geometry) (geoos.Geometry, error) {
	result, err := geo.UnaryUnion(wkt.MarshalString(geom))
	if err != nil {
		return nil, err
	}
	return wkt.UnmarshalString(result)
}

// Union returns a new geometry representing all points in this geometry and the other.
func (g *GEOAlgorithm) Union(geom1, geom2 geoos.Geometry) (geoos.Geometry, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	result, err := geo.Union(ms1, ms2)
	if err != nil {
		return nil, err
	}
	geometry, err := wkt.UnmarshalString(result)
	if err != nil {
		return nil, err
	}
	return geometry, nil
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (g *GEOAlgorithm) UniquePoints(geom geoos.Geometry) (geoos.Geometry, error) {
	result, err := geo.UniquePoints(wkt.MarshalString(geom))
	if err != nil {
		return nil, err
	}
	geometry, err := wkt.UnmarshalString(result)
	if err != nil {
		return nil, err
	}
	return geometry, nil

}

// Within returns TRUE if geometry A is completely inside geometry B.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (g *GEOAlgorithm) Within(geom1, geom2 geoos.Geometry) (bool, error) {
	ms1, ms2 := convertGeomToWKT(geom1, geom2)
	return geo.Within(ms1, ms2)
}

// convertGeomToWKT help to convert geoos.Geometry to WKT string
func convertGeomToWKT(geom1, geom2 geoos.Geometry) (string, string) {
	ms1 := wkt.MarshalString(geom1)
	ms2 := wkt.MarshalString(geom2)
	return ms1, ms2
}
