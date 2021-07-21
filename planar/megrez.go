package planar

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay"
	"github.com/spatial-go/geoos/space"
)

// MegrezAlgorithm algorithm implement
type MegrezAlgorithm struct{}

// Area returns the area of a polygonal geometry.
func (g *MegrezAlgorithm) Area(geom space.Geometry) (float64, error) {
	switch geom.GeoJSONType() {
	case space.TypePolygon:
		return geom.(space.Polygon).Area()
	case space.TypeMultiPolygon:
		return geom.(space.MultiPolygon).Area()
	default:
		return 0.0, nil
	}
}

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
func (g *MegrezAlgorithm) Boundary(geom space.Geometry) (space.Geometry, error) {
	return geom.Boundary()
}

// Buffer sReturns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (g *MegrezAlgorithm) Buffer(geom space.Geometry, width float64, quadsegs int32) (geometry space.Geometry) {
	buff := buffer.Buffer(geom.ToMatrix(), width)
	switch b := buff.(type) {
	case matrix.LineMatrix:
		return space.LineString(b)
	case matrix.PolygonMatrix:
		return space.Polygon(b)
	}
	return nil
}

// Centroid  computes the geometric center of a geometry, or equivalently, the center of mass of the geometry as a POINT.
// For [MULTI]POINTs, this is computed as the arithmetic mean of the input coordinates.
// For [MULTI]LINESTRINGs, this is computed as the weighted length of each line segment.
// For [MULTI]POLYGONs, "weight" is thought in terms of area.
// If an empty geometry is supplied, an empty GEOMETRYCOLLECTION is returned.
// If NULL is supplied, NULL is returned.
// If CIRCULARSTRING or COMPOUNDCURVE are supplied, they are converted to linestring wtih CurveToLine first,
// then same than for LINESTRING
func (g *MegrezAlgorithm) Centroid(geom space.Geometry) (space.Geometry, error) {
	if geom == nil || geom.IsEmpty() {
		return nil, nil
	}
	return space.Centroid(geom), nil
}

// Contains space.Geometry A contains space.Geometry B if and only if no points of B lie in the exterior of A,
// and at least one point of the interior of B lies in the interior of A.
// An important subtlety of this definition is that A does not contain its boundary, but A does contain itself.
// Returns TRUE if geometry B is completely inside geometry A.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (g *MegrezAlgorithm) Contains(geom1, geom2 space.Geometry) (bool, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Contains(geom1, geom2)
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (g *MegrezAlgorithm) ConvexHull(geom space.Geometry) (space.Geometry, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).ConvexHull(geom)
}

// CoveredBy returns TRUE if no point in space.Geometry A is outside space.Geometry B
func (g *MegrezAlgorithm) CoveredBy(geom1, geom2 space.Geometry) (bool, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).CoveredBy(geom1, geom2)
}

// Covers returns TRUE if no point in space.Geometry B is outside space.Geometry A
func (g *MegrezAlgorithm) Covers(geom1, geom2 space.Geometry) (bool, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Covers(geom1, geom2)
}

// Crosses takes two geometry objects and returns TRUE if their intersection "spatially cross",
// that is, the geometries have some, but not all interior points in common.
// The intersection of the interiors of the geometries must not be the empty set
// and must have a dimensionality less than the maximum dimension of the two input geometries.
// Additionally, the intersection of the two geometries must not equal either of the source geometries.
// Otherwise, it returns FALSE.
func (g *MegrezAlgorithm) Crosses(geom1, geom2 space.Geometry) (bool, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Crosses(geom1, geom2)
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (g *MegrezAlgorithm) Difference(geom1, geom2 space.Geometry) (space.Geometry, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Difference(geom1, geom2)
}

// Disjoint Overlaps, Touches, Within all imply geometries are not spatially disjoint.
// If any of the aforementioned returns true, then the geometries are not spatially disjoint.
// Disjoint implies false for spatial intersection.
func (g *MegrezAlgorithm) Disjoint(geom1, geom2 space.Geometry) (bool, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Disjoint(geom1, geom2)
}

// Distance returns the minimum 2D Cartesian (planar) distance between two geometries, in projected units (spatial ref units).
func (g *MegrezAlgorithm) Distance(geom1, geom2 space.Geometry) (float64, error) {
	return geom1.Distance(geom2)
}

// SphericalDistance calculates spherical distance
//
// To get real distance in m
func (g *MegrezAlgorithm) SphericalDistance(geom1, geom2 space.Geometry) (float64, error) {
	return geom1.SpheroidDistance(geom2)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (g *MegrezAlgorithm) Envelope(geom space.Geometry) (space.Geometry, error) {
	switch geom.GeoJSONType() {
	case space.TypePoint:
		return geom, nil
	default:
		return geom.Bound().ToPolygon(), nil
	}
}

// Equals returns TRUE if the given Geometries are "spatially equal".
func (g *MegrezAlgorithm) Equals(geom1, geom2 space.Geometry) (bool, error) {
	return geom1.Equals(geom2), nil
}

// EqualsExact returns true if both geometries are Equal, as evaluated by their
// points being within the given tolerance.
func (g *MegrezAlgorithm) EqualsExact(geom1, geom2 space.Geometry, tolerance float64) (bool, error) {
	return geom1.EqualsExact(geom2, tolerance), nil
}

// HausdorffDistance returns the Hausdorff distance between two geometries, a measure of how similar
// or dissimilar 2 geometries are. Implements algorithm for computing a distance metric which can be
// thought of as the "Discrete Hausdorff Distance". This is the Hausdorff distance restricted
// to discrete points for one of the geometries
func (g *MegrezAlgorithm) HausdorffDistance(geom1, geom2 space.Geometry) (float64, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).HausdorffDistance(geom1, geom2)
}

// HausdorffDistanceDensify computes the Hausdorff distance with an additional densification fraction amount
func (g *MegrezAlgorithm) HausdorffDistanceDensify(s, d space.Geometry, densifyFrac float64) (float64, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).HausdorffDistanceDensify(s, d, densifyFrac)
}

// Intersection returns a geometry that represents the point set intersection of the Geometries.
func (g *MegrezAlgorithm) Intersection(geom1, geom2 space.Geometry) (space.Geometry, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Intersection(geom1, geom2)
}

// Intersects If a geometry  shares any portion of space then they intersect
func (g *MegrezAlgorithm) Intersects(geom1, geom2 space.Geometry) (bool, error) {
	//todo
	return GetStrategy(newGEOAlgorithm).Intersects(geom1, geom2)
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (g *MegrezAlgorithm) IsClosed(geom space.Geometry) (bool, error) {
	elem := space.ElementValid{Geometry: geom}
	return elem.IsClosed(), nil
}

// IsEmpty returns true if this space.Geometry is an empty geometry.
// If true, then this space.Geometry represents an empty geometry collection, polygon, point etc.
func (g *MegrezAlgorithm) IsEmpty(geom space.Geometry) (bool, error) {
	return geom.IsEmpty(), nil
}

// IsRing returns true if the lineal geometry has the ring property.
func (g *MegrezAlgorithm) IsRing(geom space.Geometry) (bool, error) {
	elem := space.ElementValid{Geometry: geom}
	return elem.IsClosed() && elem.IsSimple(), nil
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points, such as self intersection or self tangency.
func (g *MegrezAlgorithm) IsSimple(geom space.Geometry) (bool, error) {
	return geom.IsSimple(), nil
}

// Length returns the 2D Cartesian length of the geometry if it is a LineString, MultiLineString
func (g *MegrezAlgorithm) Length(geom space.Geometry) (float64, error) {
	return geom.Length(), nil
}

// LineMerge returns a (set of) LineString(s) formed by sewing together the constituent line work of a MULTILINESTRING.
func (g *MegrezAlgorithm) LineMerge(geom space.Geometry) (space.Geometry, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).LineMerge(geom)
}

// NGeometry returns the number of component geometries.
func (g *MegrezAlgorithm) NGeometry(geom space.Geometry) (int, error) {
	return geom.Nums(), nil
}

// Overlaps returns TRUE if the Geometries "spatially overlap".
// By that we mean they intersect, but one does not completely contain another.
func (g *MegrezAlgorithm) Overlaps(geom1, geom2 space.Geometry) (bool, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Overlaps(geom1, geom2)
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (g *MegrezAlgorithm) PointOnSurface(geom space.Geometry) (space.Geometry, error) {
	m := buffer.InteriorPoint(geom.ToMatrix())
	return space.Point(m), nil
}

// Relate computes the intersection matrix (Dimensionally Extended
// Nine-Intersection Model (DE-9IM) matrix) for the spatial relationship between
// the two geometries.
func (g *MegrezAlgorithm) Relate(s, d space.Geometry) (string, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Relate(s, d)
}

// SharedPaths returns a collection containing paths shared by the two input geometries.
// Those going in the same direction are in the first element of the collection,
// those going in the opposite direction are in the second element.
// The paths themselves are given in the direction of the first geometry.
func (g *MegrezAlgorithm) SharedPaths(geom1, geom2 space.Geometry) (string, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).SharedPaths(geom1, geom2)
}

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (g *MegrezAlgorithm) Simplify(geom space.Geometry, tolerance float64) (space.Geometry, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Simplify(geom, tolerance)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (g *MegrezAlgorithm) SimplifyP(geom space.Geometry, tolerance float64) (space.Geometry, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).SimplifyP(geom, tolerance)
}

// Snap the vertices and segments of a geometry to another space.Geometry's vertices.
// A snap distance tolerance is used to control where snapping is performed.
// The result geometry is the input geometry with the vertices snapped.
// If no snapping occurs then the input geometry is returned unchanged.
func (g *MegrezAlgorithm) Snap(input, reference space.Geometry, tolerance float64) (space.Geometry, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Snap(input, reference, tolerance)
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (g *MegrezAlgorithm) SymDifference(geom1, geom2 space.Geometry) (space.Geometry, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).SymDifference(geom1, geom2)
}

// Touches returns TRUE if the only points in common between geom1 and geom2 lie in the union of the boundaries of geom1 and geom2.
// The ouches relation applies to all Area/Area, Line/Line, Line/Area, Point/Area and Point/Line pairs of relationships,
// but not to the Point/Point pair.
func (g *MegrezAlgorithm) Touches(geom1, geom2 space.Geometry) (bool, error) {
	//TODO
	return GetStrategy(newGEOAlgorithm).Touches(geom1, geom2)
}

// UnaryUnion does dissolve boundaries between components of a multipolygon (invalid) and does perform union
// between the components of a geometrycollection
func (g *MegrezAlgorithm) UnaryUnion(geom space.Geometry) (space.Geometry, error) {
	if geom.GeoJSONType() == space.TypeMultiPolygon {
		var matrix4 matrix.MultiPolygonMatrix
		for _, v := range geom.(space.MultiPolygon) {
			matrix4 = append(matrix4, v)
		}
		result := overlay.UnaryUnion(matrix4)
		return space.Polygon(result), nil
	}
	return nil, ErrNotPolygon
}

// Union returns a new geometry representing all points in this geometry and the other.
func (g *MegrezAlgorithm) Union(geom1, geom2 space.Geometry) (space.Geometry, error) {
	if geom1.GeoJSONType() == space.TypePolygon && geom2.GeoJSONType() == space.TypePolygon {
		result := overlay.Union(matrix.PolygonMatrix(geom1.(space.Polygon)), matrix.PolygonMatrix(geom2.(space.Polygon)))
		return space.Polygon(result), nil
	} else if geom1.GeoJSONType() == space.TypePoint && geom2.GeoJSONType() == space.TypePoint {
		return space.MultiPoint{geom1.(space.Point), geom2.(space.Point)}, nil
	} else if geom1.GeoJSONType() == space.TypeLineString && geom2.GeoJSONType() == space.TypeLineString {
		return space.MultiLineString{geom1.(space.LineString), geom2.(space.LineString)}, nil
	}
	return space.Collection{geom1, geom2}, nil
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (g *MegrezAlgorithm) UniquePoints(geom space.Geometry) (space.Geometry, error) {
	return geom.UniquePoints(), nil
}

// Within returns TRUE if geometry A is completely inside geometry B.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func (g *MegrezAlgorithm) Within(geom1, geom2 space.Geometry) (bool, error) {
	// optimization - lower dimension cannot contain areas
	if geom2.Dimensions() == 2 && geom1.Dimensions() < 2 {
		return false, nil
	}
	// optimization - P cannot contain a non-zero-length L
	// Note that a point can contain a zero-length lineal geometry,
	// since the line has no boundary due to Mod-2 Boundary Rule
	if geom2.Dimensions() == 1 && geom1.Dimensions() < 1 && geom2.Length() > 0.0 {
		return false, nil
	}
	// optimization - envelope test
	if !geom2.Bound().ContainsBound(geom1.Bound()) {
		return false, nil
	}
	// optimization for rectangle arguments
	if geom2.GeoJSONType() == space.TypePolygon && geom2.(space.Polygon).IsRectangle() {
		return geom2.Bound().ContainsBound(geom1.Bound()), nil
	}
	// general case
	//return relate(g).isContains()

	//todo
	return GetStrategy(newGEOAlgorithm).Within(geom1, geom2)
}
