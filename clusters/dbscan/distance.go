package dbscan

import (
	"math"

	"github.com/spatial-go/geoos/space"
)

// const ...
const (
	// DegreeRad is coefficient to translate from degrees to radians
	DegreeRad = math.Pi / 180.0
	// EarthR is earth radius in km
	EarthR = 6371.0
	// radius := 6371000.0 //6378137.0
)

// DistanceSpherical is a spherical (optimized) distance between two points
//
// Result is distance in kilometers
func DistanceSpherical(p1, p2 space.Point) float64 {
	v1 := (p1[1] - p2[1]) * DegreeRad
	v1 = v1 * v1

	v2 := (p1[0] - p2[0]) * DegreeRad * math.Cos((p1[1]+p2[1])/2.0*DegreeRad)
	v2 = v2 * v2

	return EarthR * math.Sqrt(v1+v2)
}

// FastSine calculates sinus approximated to parabola
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

// FastCos calculates cosines from sinus
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
func DistanceSphericalFast(p1, p2 space.Point) float64 {
	if p1.IsEmpty() {
		p1 = space.Point{0, 0}
	}
	if p2.IsEmpty() {
		p2 = space.Point{0, 0}
	}
	v1 := (p1[1] - p2[1])
	v2 := (p1[0] - p2[0]) * FastCos((p1[1]+p2[1])/2.0*DegreeRad)

	return v1*v1 + v2*v2
}
