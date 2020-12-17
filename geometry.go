package geos

<<<<<<< HEAD
// Geometry is the interface implemented by an object that can ...
=======
// Geometry ...
>>>>>>> a1310ea0c8d9c74dab601390b4f91981dc57f2c9
type Geometry interface {
	GeoJSONType() string
	Dimensions() int // e.g. 0d, 1d, 2d
	// Num of geometries
	Nums() int
}
