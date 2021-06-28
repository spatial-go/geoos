package dbscan

import (
	"math"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/common"
)

// DistanceSpherical is a spherical (optimized) distance between two points
//
// Result is distance in kilometers
func DistanceSpherical(p1, p2 geoos.Point) float64 {
	v1 := (p1[1] - p2[1]) * common.DegreeRad
	v1 = v1 * v1

	v2 := (p1[0] - p2[0]) * common.DegreeRad * math.Cos((p1[1]+p2[1])/2.0*common.DegreeRad)
	v2 = v2 * v2

	return common.EarthR * math.Sqrt(v1+v2)
}

// FastSine caclulates sinus approximated to parabola
//
// Taken from: http://forum.devmaster.net/t/fast-and-accurate-sine-cosine/9648
func FastSine(x float64) float64 {
	const (
		B = 4 / math.Pi
		C = -4 / (math.Pi * math.Pi)
		P = 0.225
	)

	if x > math.Pi || x < -math.Pi {
		panic("out of range")
	}

	y := B*x + C*x*math.Abs(x)
	return P*(y*math.Abs(y)-y) + y
}

// FastCos calculates cosinus from sinus
func FastCos(x float64) float64 {
	x += math.Pi / 2.0
	for x > math.Pi {
		x -= 2 * math.Pi
	}

	return FastSine(x)
}

// DistanceSphericalFast calculates spherical distance with fast cosine
// without sqrt and normalization to Earth radius/radians
//
// To get real distance in km, take sqrt and multiply result by EarthR*DegreeRad
//
// In this library eps (distance) is adjusted so that we don't need
// to do sqrt and multiplication
func DistanceSphericalFast(p1, p2 geoos.Point) float64 {
	v1 := (p1[1] - p2[1])
	v2 := (p1[0] - p2[0]) * FastCos((p1[1]+p2[1])/2.0*common.DegreeRad)

	return v1*v1 + v2*v2
}
