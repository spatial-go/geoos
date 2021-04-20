package common

import "math"

const (
	// DegreeRad is coefficient to translate from degrees to radians
	DegreeRad = math.Pi / 180.0
	// EarthR is earth radius in km
	EarthR = 6371.0
	// radius := 6371000.0 //6378137.0

	// Angle of sin 60 = 0.866025403785
	Sin60 = 0.866025403785
	Cos60 = 0.5
)
