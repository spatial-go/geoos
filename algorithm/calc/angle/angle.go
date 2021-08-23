// Package angle define angel calculation function.
package angle

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

const (
	// PiTimes2 The value of 2*Pi
	PiTimes2 = 2.0 * math.Pi
	// PiOver2 The value of Pi/2
	PiOver2 = math.Pi / 2.0

	// PiOver4 PI_OVER_4The value of Pi/4
	PiOver4 = math.Pi / 4.0

	// None Constant representing no orientation
	None = 0

	// PiDegrees ...
	PiDegrees = 180.0
)

// ToDegrees Converts from radians to degrees.
func ToDegrees(radians float64) float64 {
	return (radians * PiDegrees) / (math.Pi)
}

// ToRadians  Converts from degrees to radians.
func ToRadians(angleDegrees float64) float64 {
	return (angleDegrees * math.Pi) / PiDegrees
}

// Angle Returns the angle of the vector from p0 to p1,
// relative to the positive X-axis.
// The angle is normalized to be in the range [ -Pi, Pi ].
func Angle(p0, p1 matrix.Matrix) float64 {
	dx := p1[0] - p0[0]
	dy := p1[1] - p0[1]
	return math.Atan2(dy, dx)
}

// MatrixAngle Returns the angle of the vector from (0,0) to p,
// relative to the positive X-axis.
// The angle is normalized to be in the range [ -Pi, Pi ].
func MatrixAngle(p matrix.Matrix) float64 {
	return math.Atan2(p[1], p[0])
}

// IsAcute Tests whether the angle between p0-p1-p2 is acute.
// An angle is acute if it is less than 90 degrees.
func IsAcute(p0, p1, p2 matrix.Matrix) bool {
	// relies on fact that A dot B is positive if A ang B is acute
	dx0 := p0[0] - p1[0]
	dy0 := p0[1] - p1[1]
	dx1 := p2[0] - p1[0]
	dy1 := p2[1] - p1[1]
	dotprod := dx0*dx1 + dy0*dy1
	return dotprod > 0
}

// IsObtuse Tests whether the angle between p0-p1-p2 is obtuse.
// An angle is obtuse if it is greater than 90 degrees.
func IsObtuse(p0, p1, p2 matrix.Matrix) bool {
	// relies on fact that A dot B is negative if A ang B is obtuse
	dx0 := p0[0] - p1[0]
	dy0 := p0[1] - p1[1]
	dx1 := p2[0] - p1[0]
	dy1 := p2[1] - p1[1]
	dotprod := dx0*dx1 + dy0*dy1
	return dotprod < 0
}

// Between Returns the smallest angle between two vectors.
// The computed angle will be in the range [0, Pi).
func Between(tip1, tail, tip2 matrix.Matrix) float64 {
	a1 := Angle(tail, tip1)
	a2 := Angle(tail, tip2)

	return Diff(a1, a2)
}

// BetweenOriented Returns the oriented smallest angle between two vectors.
// The computed angle will be in the range (-Pi, Pi].
//  A positive result corresponds to a counterclockwise
// (CCW) rotation
func BetweenOriented(tip1, tail, tip2 matrix.Matrix) float64 {
	a1 := Angle(tail, tip1)
	a2 := Angle(tail, tip2)
	angDel := a2 - a1

	// normalize, maintaining orientation
	if angDel <= -math.Pi {
		return angDel + PiTimes2
	}
	if angDel > math.Pi {
		return angDel - PiTimes2
	}
	return angDel
}

// InteriorAngle Computes the interior angle between two segments of a ring. The ring is
//  assumed to be oriented in a clockwise direction. The computed angle will be
//  in the range [0, 2Pi]
func InteriorAngle(p0, p1, p2 matrix.Matrix) float64 {
	anglePrev := Angle(p1, p0)
	angleNext := Angle(p1, p2)
	return NormalizePositive(angleNext - anglePrev)
}

// Turn Returns whether an angle must turn clockwise or counterclockwise
// to overlap another angle.
func Turn(ang1, ang2 float64) int {
	crossproduct := math.Sin(ang2 - ang1)

	if crossproduct > 0 {
		return calc.CounterClockWise
	}
	if crossproduct < 0 {
		return calc.ClockWise
	}
	return None
}

// Normalize Computes the normalized value of an angle, which is the
// equivalent angle in the range ( -Pi, Pi ].
func Normalize(angle float64) float64 {
	for angle > math.Pi {
		angle -= PiTimes2
	}
	for angle <= -math.Pi {
		angle += PiTimes2
	}
	return angle
}

// NormalizePositive Computes the normalized positive value of an angle, which is the
//  equivalent angle in the range [ 0, 2*Pi ).
func NormalizePositive(angle float64) float64 {
	if angle < 0.0 {
		for angle < 0.0 {
			angle += PiTimes2
		}
		// in case round-off error bumps the value over
		if angle >= PiTimes2 {
			angle = 0.0
		}
	} else {
		for angle >= PiTimes2 {
			angle -= PiTimes2
		}
		// in case round-off error bumps the value under
		if angle < 0.0 {
			angle = 0.0
		}
	}
	return angle
}

// Diff Computes the smallest difference between two angles.
// The angles are assumed to be normalized to the range [-Pi, Pi].
// The result will be in the range [0, Pi].
func Diff(ang1, ang2 float64) float64 {
	delAngle := 0.0

	if ang1 < ang2 {
		delAngle = ang2 - ang1
	} else {
		delAngle = ang1 - ang2
	}

	if delAngle > math.Pi {
		delAngle = (2 * math.Pi) - delAngle
	}

	return delAngle
}
