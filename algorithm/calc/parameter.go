package calc

import "math"

const (

	// OffsetSegmentSeparationFactor  Factor which controls how close offset segments can be to
	//   skip adding a filler or mitre.
	OffsetSegmentSeparationFactor = 1.0e-3

	// InsideTurnVertexSnapDistanceFactor Factor which controls how close curve vertices on inside turns can be to be snapped
	InsideTurnVertexSnapDistanceFactor = 1.0e-3

	// CurveVertexSnapDistanceFactor Factor which controls how close curve vertices can be to be snapped
	CurveVertexSnapDistanceFactor = 1.0e-6

	// MaxClosingSegLenFactor Factor which determines how short closing segs can be for round buffers
	MaxClosingSegLenFactor = 80
)

// const ...
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

// const overlay parameters.
const (
	// POINTS ...
	POINTS = iota
	// CLOSED egde is closed.
	CLOSED
	// MAIN overlay main polygon.
	MAIN
	// CUT overlay cut polygon.
	CUT
	// CLIP overlay clip.
	CLIP
	// MERGE overlay merge.
	MERGE
)

// const default parameters.
const (
	// CLOCKWISE ...
	CLOCKWISE        = -1
	COUNTERCLOCKWISE = 1

	// ANGLE ...
	ANGLE = 2.0

	LEFT  = 1
	RIGHT = 2

	// MinRingSize ...
	MinRingSize = 3
	// MaxRingSize ...
	MaxRingSize = 9
	// NearnessFactor ...
	NearnessFactor = 0.99

	// CAPROUND Specifies a round line buffer end cap style.
	CAPROUND = 1
	// CAPFLAT Specifies a flat line buffer end cap style.
	CAPFLAT = 2
	// CAPSQUARE Specifies a square line buffer end cap style.
	CAPSQUARE = 3

	//JOINROUND Specifies a round join style.
	JOINROUND = 1
	// JOINMITRE Specifies a mitre join style.
	JOINMITRE = 2
	// JOINBEVEL Specifies a bevel join style.
	JOINBEVEL = 3

	// QUADRANTSEGMENTS The default number of facets into which to divide a fillet of 90 degrees.
	// A value of 8 gives less than 2% max error in the buffer distance.
	// For a max error of &lt; 1%, use QS = 12.
	// For a max error of &lt; 0.1%, use QS = 18.
	QUADRANTSEGMENTS = 8

	// MITRELIMIT
	// The default mitre limit
	// Allows fairly pointy mitres.
	MITRELIMIT = 5.0

	// SIMPLIFYFACTOR
	// The default simplify factor
	// Provides an accuracy of about 1%, which matches the accuracy of the default Quadrant Segments parameter.
	SIMPLIFYFACTOR = 0.01
)

// const default DE-9IM  and Dimension parameters.
const (
	// The location value for the exterior of a geometry.
	// Also, DE-9IM row index of the exterior of the first geometry and column index
	INTERIOR = 0
	BOUNDARY = 1
	EXTERIOR = 2

	// FALSE Dimension value of the empty geometry (-1).
	// TRUE Dimension value of non-empty geometries (= {P, L, A}).
	// DONTCARE Dimension value for any dimension (= {FALSE, TRUE}).
	FALSE    = -1
	TRUE     = -2
	DONTCARE = -3

	P           = 0
	L           = 1
	A           = 2
	SYMFALSE    = 'F'
	SYMTRUE     = 'T'
	SYMDONTCARE = '*'
	SYMP        = '0'
	SYML        = '1'
	SYMA        = '2'
)

// const calc parameter
const (
	// The smallest representable relative difference between two  values.
	EPS   = 1.23259516440783e-32 /* = 2^-106 */
	SPLIT = 134217729.0          // 2^27+1, for IEEE

	MAXPRINTDIGITS = 32

	SCINOTEXPONENTCHAR = "E"
	SCINOTZERO         = "0.0E0"
)

var (
	// PI The value nearest to the constant Pi.
	PI = &PairFloat{3.141592653589793116e+00,
		1.224646799147353207e-16}
	// TWOPI The value nearest to the constant 2 * Pi.
	TWOPI = &PairFloat{
		6.283185307179586232e+00,
		2.449293598294706414e-16}
	// PI2 The value nearest to the constant Pi / 2.
	PI2 = &PairFloat{
		1.570796326794896558e+00,
		6.123233995736766036e-17}
	//E  The value nearest to the constant e (the natural logarithm base).
	E = &PairFloat{
		2.718281828459045091e+00,
		1.445646891729250158e-16}
)
