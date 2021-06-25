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
