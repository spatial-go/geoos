package algorithm

import "C"
import (
	"github.com/spatial-go/geos/geo"
)

// CapStyle is the style of the cap at the end of a line segment.
type CapStyle int

const (
	_ CapStyle = iota
	// CapRound is a round end cap.
	CapRound
	// CapFlat is a flat end cap.
	CapFlat
	// CapSquare is a square end cap.
	CapSquare
)

// JoinStyle is the style of the joint of two line segments.
type JoinStyle int

const (
	_ JoinStyle = iota
	// JoinRound is a round segment join style.
	JoinRound
	// JoinMitre is a mitred segment join style.
	JoinMitre
	// JoinBevel is a beveled segment join style.
	JoinBevel
)

// BufferOpts are options to the BufferWithOpts method.
type BufferOpts struct {
	// QuadSegs is the number of quadrant segments.
	QuadSegs int
	// CapStyle is the end cap style.
	CapStyle CapStyle
	// JoinStyle is the line segment join style.
	JoinStyle JoinStyle
	// MitreLimit is the limit in the amount of a mitred join.
	MitreLimit float64
}


// buffer计算一个新的几何体作为几何体的膨胀（位置量）或侵蚀量（负量）——分别是几何体的和或差，其半径为缓冲量绝对值的圆
func Buffer(g geo.GEOSGeometry, width float64, quadsegs int32) geo.GEOSGeometry {
	return geo.Buffer(g, width, quadsegs)
}
// 如果两个几何图形相等，则EqualsExact将返回true，因为它们的点在给定公差内。
func EqualsExact(s geo.GEOSGeometry, d geo.GEOSGeometry, tolerance float64) bool{
	return geo.EqualsExact(s, d, tolerance)
}


// BufferWithOpts computes a new geometry as the dilation (position amount) or erosion
// (negative amount) of the geometry -- a sum or difference, respectively, of
// the geometry with a circle of radius of the absolute value of the buffer
// amount.
//
// BufferWithOpts gives the user more control than Buffer over the parameters of
// the buffering, including:
//
//  - # of quadrant segments (defaults to 8 in Buffer)
//  - mitre limit (defaults to 5.0 in Buffer)
//  - end cap style (see CapStyle consts)
//  - join style (see JoinStyle consts)
func  BufferWithOpts(width float64, opts BufferOpts) (geo.GEOSGeometry, error) {
	return nil,nil
}

// OffsetCurve computes a new linestring that is offset from the input
// linestring by the given distance and buffer options.  A negative distance is
// offset on the right side; positive distance offset on the left side.
func  OffsetCurve(distance float64, opts BufferOpts) (geo.GEOSGeometry, error) {
	return nil,nil
}



// NewCollection returns a new geometry that is a collection containing multiple
// geometries given as variadic arguments. The type of the collection (in the
// SFS sense of type -- MultiPoint, MultiLineString, etc.) is determined by the
// first argument. If no geometries are given, the geometry is an empty version
// of the given collection type.
func NewCollection() (geo.GEOSGeometry, error) {
	return nil,nil
}

// Clone performs a deep copy on the geometry.
func  Clone() (geo.GEOSGeometry, error) {
	return nil,nil
}



//=====================一元拓扑函数===============================

// Envelope is the bounding box of a geometry, as a polygon.
func  Envelope() (geo.GEOSGeometry, error) {
	return nil,nil
}

// ConvexHull computes the smallest convex geometry that contains all the points
// of the geometry.
func  ConvexHull() (geo.GEOSGeometry, error) {
	return nil,nil
}

// Boundary is the boundary of the geometry.
func  Boundary() (geo.GEOSGeometry, error) {
	return nil,nil
}

// UnaryUnion computes the union of all the constituent geometries of the
// geometry.
func  UnaryUnion() (geo.GEOSGeometry, error) {
	return nil,nil
}

// PointOnSurface computes a point geometry guaranteed to be on the surface of
// the geometry.
func  PointOnSurface() (geo.GEOSGeometry, error) {
	return nil,nil
}

// Centroid is the center point of the geometry.
func  Centroid() (geo.GEOSGeometry, error) {
	return nil,nil
}

// LineMerge will merge together a collection of LineStrings where they touch
// only at their start and end points. The LineStrings must be fully noded. The
// resulting geometry is a new collection.
func  LineMerge() (geo.GEOSGeometry, error) {
	return nil,nil
}

// Simplify returns a geometry simplified by amount given by tolerance.
// May not preserve topology -- see SimplifyP.
func  Simplify(tolerance float64) (geo.GEOSGeometry, error) {
	return nil,nil
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func  SimplifyP(tolerance float64) (geo.GEOSGeometry, error) {
	return nil,nil
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func  UniquePoints() (geo.GEOSGeometry, error) {
	return nil,nil
}

// SharedPaths finds paths shared between the two given lineal geometries.
// Returns a GeometryCollection having two elements:
//	- first element is a MultiLineString containing shared paths having the _same_ direction on both inputs
//	- second element is a MultiLineString containing shared paths having the _opposite_ direction on the two inputs
func  SharedPaths(other geo.GEOSGeometry) (geo.GEOSGeometry, error) {
	return nil,nil
}

// Snap returns a new geometry where the geometry is snapped to the given
// geometry by given tolerance.
func  Snap(other geo.GEOSGeometry, tolerance float64) (geo.GEOSGeometry, error) {
	return nil,nil
}


// =================================二元拓扑函数=============================================

// Intersection returns a new geometry representing the points shared by this
// geometry and the other.
func  Intersection(other geo.GEOSGeometry) (geo.GEOSGeometry, error) {
	return nil,nil
}

// Difference returns a new geometry representing the points making up this
// geometry that do not make up the other.
func  Difference(other geo.GEOSGeometry) (geo.GEOSGeometry, error) {
	return nil,nil
}

// SymDifference returns a new geometry representing the set combining the
// points in this geometry not in the other, and the points in the other
// geometry and not in this.
func  SymDifference(other geo.GEOSGeometry) (geo.GEOSGeometry, error) {
	return nil,nil
}

// Union returns a new geometry representing all points in this geometry and the
// other.
func  Union(other geo.GEOSGeometry) (geo.GEOSGeometry, error) {
	return nil,nil
}


// ==========================二元谓词函数======================================

// Disjoint returns true if the two geometries have no point in common.
func  Disjoint(other geo.GEOSGeometry) (bool, error) {
	return true,nil
}

// Touches returns true if the two geometries have at least one point in common,
// but their interiors do not intersect.
func  Touches(other geo.GEOSGeometry) (bool, error) {
	return true,nil
}

// Intersects returns true if the two geometries have at least one point in
// common.
func  Intersects(other geo.GEOSGeometry) (bool, error) {
	return true,nil
}

// Crosses returns true if the two geometries have some but not all interior
// points in common.
func  Crosses(other geo.GEOSGeometry) (bool, error) {
	return true,nil
}

// Within returns true if every point of this geometry is a point of the other,
// and the interiors of the two geometries have at least one point in common.
func  Within(other geo.GEOSGeometry) (bool, error) {
	return true,nil
}

// Contains returns true if every point of the other is a point of this geometry,
// and the interiors of the two geometries have at least one point in common.
func  Contains(other geo.GEOSGeometry) (bool, error) {
	return true,nil
}

// Overlaps returns true if the geometries have some but not all points in
// common, they have the same dimension, and the intersection of the interiors
// of the two geometries has the same dimension as the geometries themselves.
func  Overlaps(other geo.GEOSGeometry) (bool, error) {
	return true,nil
}

// Equals returns true if the two geometries have at least one point in common,
// and no point of either geometry lies in the exterior of the other geometry.
func  Equals(other geo.GEOSGeometry) (bool, error) {
	return true,nil
}

// Covers returns true if every point of the other geometry is a point of this
// geometry.
func  Covers(other geo.GEOSGeometry) (bool, error) {
	return true,nil
}

// CoveredBy returns true if every point of this geometry is a point of the
// other geometry.
func  CoveredBy(other geo.GEOSGeometry) (bool, error) {
	return true,nil
}

// ===========================一元谓词函数===================================

// IsEmpty returns true if the set of points of this geometry is empty (i.e.,
// the empty geometry).
func  IsEmpty() (bool, error) {
	return true,nil
}

// IsSimple returns true iff the only self-intersections are at boundary points.
func  IsSimple() (bool, error) {
	return true,nil
}

// IsRing returns true if the lineal geometry has the ring property.
func  IsRing() (bool, error) {
	return true,nil
}

// HasZ returns true if the geometry is 3D.
func  HasZ() (bool, error) {
	return true,nil
}

// IsClosed returns true if the geometry is closed (i.e., start & end points
// equal).
func  IsClosed() (bool, error) {
	return true,nil
}

//  =========================== 几何信息函数===================================

// Type returns the SFS type of the geometry.
func  Type() (geo.GEOSGeometryType, error) {
	
	return nil, nil
}

// SRID returns the geometry's SRID, if set.
func  SRID() (int, error) {
	return 1, nil
}

// SetSRID sets the geometry's SRID.
func  SetSRID(srid int) {
	
}

// NGeometry returns the number of component geometries (eg., for
// a collection).
func  NGeometry() (int, error) {
	return 1,nil
}

// XXX: method to return a slice of geometries

// Geometry returns the nth sub-geometry of the geometry (eg., of a collection).
func  Geometry(n int) (geo.GEOSGeometry, error) {
	// According to GEOS C API, GEOSGetGeometryN returns a pointer to internal
	// storage and must not be destroyed directly, so we bypass the regular
	// constructor to avoid the finalizer.
	return nil,nil
}

// Normalize computes the normal form of the geometry.
// Modifies geometry in-place, clone first if this is not wanted/safe.
func  Normalize() error {
	return nil
}

// NPoint returns the number of points in the geometry.
func  NPoint() (int, error) {
	return 1,nil
}

type float64Getter func(*C.GEOSGeometry, *C.double) C.int

// X returns the x ordinate of the geometry.
// Geometry must be a Point.
func  X() (float64, error) {
	return 1,nil
}

// Y returns the y ordinate of the geometry.
// Geometry must be a Point
func  Y() (float64, error) {
	return 1,nil
}

// Holes returns a slice of geometries (LinearRings) representing the interior
// rings of a polygon (possibly nil).
// Geometry must be a Polygon.
func  Holes() ([]geo.GEOSGeometry, error) {

	return nil, nil
}

// XXX: Holes() returns a [][]Coord?

// Shell returns the exterior ring (a LinearRing) of the geometry.
// Geometry must be a Polygon.
func  Shell() (geo.GEOSGeometry, error) {
	// According to the GEOS C API, GEOSGetExteriorRing returns a pointer
	// to internal storage and must not be destroyed directly, so we bypass
	// the usual constructor to avoid the finalizer.
	return nil,nil
}

// NCoordinate returns the number of coordinates of the geometry.
func  NCoordinate() (int, error) {
	return 1,nil
}

// Coords returns a slice of Coord, a sequence of coordinates underlying the
// point, linestring, or linear ring.
func  Coords() ( error) {
	return nil
}

// Dimension returns the number of dimensions geometry, eg., 1 for point, 2 for
// linestring.
func  Dimension() int {
	return 1
}

// CoordDimension returns the number of dimensions of the coordinates of the
// geometry (2 or 3).
func  CoordDimension() int {
	return 1
}

// Point returns the nth point of the geometry.
// Geometry must be LineString.
func  Point(n int) (geo.GEOSGeometry, error) {
	return nil,nil
}

// StartPoint returns the 0th point of the geometry.
// Geometry must be LineString.
func  StartPoint() (geo.GEOSGeometry, error) {
	return nil,nil
}

// EndPoint returns the (n-1)th point of the geometry.
// Geometry must be LineString.
func  EndPoint() (geo.GEOSGeometry, error) {
	return nil,nil
}





// 其他功能

// Area returns the area of the geometry, which must be a areal geometry like
// a polygon or multipolygon.
func  Area() (float64, error) {
	return 0,nil
}

// Length returns the length of the geometry, which must be a lineal geometry
// like a linestring or linear ring.
func  Length() (float64, error) {
	return 0,nil
}

// Distance returns the Cartesian distance between the two geometries.
func  Distance(other geo.GEOSGeometry) (float64, error) {
	return 0,nil
}

// HausdorffDistance computes the maximum distance of the geometry to the nearest
// point in the other geometry (i.e., considers the whole shape and position of
// the geometries).
func  HausdorffDistance(other geo.GEOSGeometry) (float64, error) {
	return 0,nil
}

// HausdorffDistanceDensify computes the Hausdorff distance (see
// HausdorffDistance) with an additional densification fraction amount.
func  HausdorffDistanceDensify(other geo.GEOSGeometry, densifyFrac float64) (float64, error) {
	return 0,nil
}


// Relate computes the intersection matrix (Dimensionally Extended
// Nine-Intersection Model (DE-9IM) matrix) for the spatial relationship between
// the two geometries.
func  Relate(other geo.GEOSGeometry) (string, error) {

	return "nil", nil
}

// RelatePat returns true if the DE-9IM matrix equals the intersection matrix of
// the two geometries.
func  RelatePat(other geo.GEOSGeometry, pat string) (bool, error) {
	return true,nil
}


