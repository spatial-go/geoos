package geos

type Geometry interface {
	GeoJSONType() string
	Dimensions() int // e.g. 0d, 1d, 2d
	// Num of geometries
	Nums() int
}
