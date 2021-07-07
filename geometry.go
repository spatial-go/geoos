package geoos

// Geometry is the interface implemented by other spatial objects
type Geometry interface {
	GeoJSONType() string
	// e.g. 0d, 1d, 2d
	Dimensions() int

	Bound() Bound
	// Num of geometries
	Nums() int

	// Area returns the area of a polygonal geometry.
	Area() (float64, error)

	// Equal returns true if the Geometry represents the same Geometry or vector.
	Equal(g Geometry) bool

	// IsEmpty returns true if the Geometry is empty.
	IsEmpty() bool

	// EqualsExact Returns true if the two Geometrys are exactly equal,
	// up to a specified distance tolerance.
	// Two Geometries are exactly equal within a distance tolerance
	EqualsExact(g Geometry, tolerance float64) bool
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
)
