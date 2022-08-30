package calc

// Defined constant variable
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

	// DecimalPlaces ...
	DecimalPlaces = "%.10f"

	// DefaultTolerance6 ...
	DefaultTolerance6 = 1.0e-6

	// DefaultTolerance7 ...
	DefaultTolerance7 = 1.0e-7

	// DefaultTolerance8 ...
	DefaultTolerance8 = 1.0e-8

	// DefaultTolerance9 ...
	DefaultTolerance9 = 1.0e-9

	// DefaultTolerance10 ...
	DefaultTolerance10 = 1.0e-10

	// DefaultTolerance12 ...
	DefaultTolerance12 = 1.0e-12

	// DefaultTolerance13 ...
	DefaultTolerance13 = 1.0e-13

	// DefaultTolerance15 ...
	DefaultTolerance15 = 1.0e-15

	AccuracyFloat = DefaultTolerance12

	DefaultTolerance = DefaultTolerance10
)

// const  Defined constant variable overlay parameters.
const (
	// OverlayPoints ...
	OverlayPoints = iota
	// OverlayClosed edge is closed.
	OverlayClosed
	// OverlayMain overlay main polygon.
	OverlayMain
	// OverlayCut overlay cut polygon.
	OverlayCut
	// OverlayClip overlay clip.
	OverlayClip
	// OverlayMerge overlay merge.
	OverlayMerge
)

// const Defined constant variable  parameters.
const (
	// ClockWise ...
	ClockWise        = -1
	CounterClockWise = 1

	SideLeft  = 1
	SideRight = 2

	// MinRingSize ...
	MinRingSize = 3
	// MaxRingSize ...
	MaxRingSize = 9
	// NearnessFactor ...
	NearnessFactor = 0.99

	// CapRound Specifies a round line buffer end cap style.
	CapRound = 1
	// CapFlat Specifies a flat line buffer end cap style.
	CapFlat = 2
	// CapSquare Specifies a square line buffer end cap style.
	CapSquare = 3

	//JoinRound Specifies a round join style.
	JoinRound = 1
	// JoinMitre Specifies a mitre join style.
	JoinMitre = 2
	// JoinBevel Specifies a bevel join style.
	JoinBevel = 3

	// QuadrantSegments The default number of facets into which to divide a fillet of 90 degrees.
	// A value of 8 gives less than 2% max error in the buffer distance.
	// For a max error of &lt; 1%, use QS = 12.
	// For a max error of &lt; 0.1%, use QS = 18.
	QuadrantSegments = 8

	// MitreLimit
	// The default mitre limit
	// Allows fairly pointy mitres.
	MitreLimit = 5.0

	// SimplifyFactor
	// The default simplify factor
	// Provides an accuracy of about 1%, which matches the accuracy of the default Quadrant Segments parameter.
	SimplifyFactor = 0.01
)

// const default DE-9IM  and Dimension parameters.
const (
	// The location value for the exterior of a geometry.
	// Also, DE-9IM row index of the exterior of the first geometry and column index
	ImInterior = 0
	ImBoundary = 1
	ImExterior = 2

	// ImFalse Dimension value of the empty geometry (-1).
	// ImTrue Dimension value of non-empty geometries (= {P, L, A}).
	// ImNotCare Dimension value for any dimension (= {ImFalse, TRUE}).
	ImFalse   = -1
	ImTrue    = -2
	ImNotCare = -3

	ImP          = 0
	ImL          = 1
	ImA          = 2
	ImSymFalse   = 'F'
	ImSymTrue    = 'T'
	ImSymNotCare = '*'
	ImSymP       = '0'
	ImSymL       = '1'
	ImSymA       = '2'
)
