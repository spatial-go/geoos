package planar

import (
	"sync"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/encoding/wkt"
)

var algorithmGeos, algorithmMegrez Algorithm
var once sync.Once

type newAlgorithm func() Algorithm

// NormalStrategy returns normal algorithm.
func NormalStrategy() Algorithm {
	return GetStrategy(newMegrezAlgorithm)
}

// GetStrategy returns  algorithm by newAlorithm.
func GetStrategy(f newAlgorithm) Algorithm {
	return f()
}

func newMegrezAlgorithm() Algorithm {
	once.Do(func() {
		algorithmMegrez = &MegrezAlgorithm{}
	})
	return algorithmMegrez
}

func newGEOAlgorithm() Algorithm {
	once.Do(func() {
		algorithmGeos = &GEOAlgorithm{}
	})
	return algorithmGeos
}

// convertGeomToWKT help to convert geoos.Geometry to WKT string
func convertGeomToWKT(geom1, geom2 geoos.Geometry) (string, string) {
	ms1 := wkt.MarshalString(geom1)
	ms2 := wkt.MarshalString(geom2)
	return ms1, ms2
}
