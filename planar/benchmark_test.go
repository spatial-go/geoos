package planar

import (
	"testing"

	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/space"
)

var geom space.Geometry = space.Polygon{{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}, {-1, -1}}}

// Benchmark_MegrezArea test megrez area
func Benchmark_MegrezArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NormalStrategy().Area(geom)
	}
}

// Benchmark_GeosArea test geos area
func Benchmark_GeosArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GetStrategy(newGEOAlgorithm).Area(geom)
	}
}

// Benchmark_MegrezUnaryUnion test megrez UnaryUnion
func Benchmark_MegrezUnaryUnion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiPolygon, _ := wkt.UnmarshalString(`MULTIPOLYGON(((0 0, 10 0, 10 10, 0 10, 0 0)), ((5 5, 15 5, 15 15, 5 15, 5 5)))`)
		G := NormalStrategy()
		_, _ = G.UnaryUnion(multiPolygon)

	}
}

// Benchmark_GeosUnaryUnion test geos UnaryUnion
func Benchmark_GeosUnaryUnion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiPolygon, _ := wkt.UnmarshalString(`MULTIPOLYGON(((0 0, 10 0, 10 10, 0 10, 0 0)), ((5 5, 15 5, 15 15, 5 15, 5 5)))`)
		G := GetStrategy(newGEOAlgorithm)
		_, _ = G.UnaryUnion(multiPolygon)
	}
}
