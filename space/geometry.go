package space

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// const geomtype
const (
	TypePoint      = "Point"
	TypeMultiPoint = "MultiPoint"

	TypeLineString      = "LineString"
	TypeMultiLineString = "MultiLineString"

	TypePolygon      = "Polygon"
	TypeMultiPolygon = "MultiPolygon"

	TypeCollection = "GeometryCollection"

	TypeBound = "Bound"

	TypeCircle = "Polygon"
)

// Geometry is the interface implemented by other spatial objects
type Geometry interface {

	// CoordinateSystem return Coordinate System.
	CoordinateSystem() int

	// Geom return Geometry without Coordinate System.
	Geom() Geometry

	GeoJSONType() string
	// e.g. 0d, 1d, 2d
	Dimensions() int

	Bound() Bound

	// Num of geometries
	Nums() int

	// IsCollection returns true if the Geometry is  collection.
	IsCollection() bool

	// ToMatrix returns the Steric of a  geometry.
	ToMatrix() matrix.Steric

	// Area returns the area of a polygonal geometry.
	Area() (float64, error)

	// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
	Boundary() (Geometry, error)

	// Buffer Returns a geometry that represents all points whose distance
	// from this space.Geometry is less than or equal to distance.
	Buffer(width float64, quadsegs int) Geometry

	// BufferInMeter Returns a geometry that represents all points whose distance
	// from this space.Geometry is less than or equal to distance.
	BufferInMeter(width float64, quadsegs int) Geometry

	// Centroid Computes the centroid point of a geometry.
	Centroid() Point

	// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
	// that encloses all geometries in the input.
	// In the general case the convex hull is a Polygon.
	// The convex hull of two or more collinear points is a two-point LineString.
	// The convex hull of one or more identical points is a Point.
	ConvexHull() Geometry

	// Distance returns distance Between the two Geometry.
	Distance(g Geometry) (float64, error)

	// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
	// The polygon is defined by the corner points of the bounding box
	// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
	Envelope() Geometry

	// Equals returns true if the Geometry represents the same Geometry or vector.
	Equals(g Geometry) bool

	// EqualsExact Returns true if the two Geometries are exactly equal,
	// up to a specified distance tolerance.
	// Two Geometries are exactly equal within a distance tolerance
	EqualsExact(g Geometry, tolerance float64) bool

	// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
	// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
	IsClosed() bool

	// IsEmpty returns true if the Geometry is empty.
	IsEmpty() bool

	// IsRing returns true if the lineal geometry has the ring property.
	IsRing() bool

	// IsSimple returns true if this space.Geometry has no anomalous geometric points,
	// such as self intersection or self tangency.
	IsSimple() bool

	// IsValid returns true if the  geometry is valid.
	IsValid() bool

	// IsCorrect returns true if the geometry struct is Correct.
	IsCorrect() bool

	// Length Returns the length of this geometry
	Length() float64

	// PointOnSurface Returns a POINT guaranteed to intersect a surface.
	PointOnSurface() Geometry

	// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
	UniquePoints() MultiPoint

	// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
	// May not preserve topology
	Simplify(tolerance float64) Geometry

	// SimplifyP returns a geometry simplified by amount given by tolerance.
	// Unlike Simplify, SimplifyP guarantees it will preserve topology.
	SimplifyP(tolerance float64) Geometry

	// SpheroidDistance returns  spheroid distance Between the two Geometry.
	SpheroidDistance(g Geometry) (float64, error)

	// Filter Performs an operation with the provided .
	Filter(f matrix.Filter) Geometry
}

// compile time checks
var (
	_ Geometry = Point{}
	_ Geometry = MultiPoint{}
	_ Geometry = LineString{}
	_ Geometry = MultiLineString{}
	_ Geometry = Ring{}
	_ Geometry = Polygon{}
	_ Geometry = MultiPolygon{}
	_ Geometry = Bound{}

	_ Geometry = Collection{}
	_ Geometry = &GeometryValid{}

	_ Geometry = &Circle{}
	_ Geometry = &GeometryValid{}
)
