package geoos

// Geometry is the interface implemented by other spatial objects
type Geometry interface {
	GeoJSONType() string
	// e.g. 0d, 1d, 2d
	Dimensions() int
	// Num of geometries
	Nums() int
}
