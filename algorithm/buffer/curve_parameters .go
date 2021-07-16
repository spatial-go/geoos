package buffer

// const
const (
	// CLOCKWISE ...
	CLOCKWISE = -1
	// ANGLE ...
	ANGLE = 2.0

	LEFT     = 1
	RIGHT    = 2
	INTERIOR = 0
	EXTERIOR = 2

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

// CurveParameters  A value containing the parameters which
// specify how a buffer should be constructed..
type CurveParameters struct {
	QuadrantSegments, EndCapStyle, JoinStyle int
	MitreLimit, SimplifyFactor               float64
	IsSingleSided                            bool
}

// DefaultCurveParameters Creates a default set of parameters.
func DefaultCurveParameters() *CurveParameters {
	return &CurveParameters{
		QUADRANTSEGMENTS,
		CAPROUND,
		JOINROUND,
		MITRELIMIT,
		SIMPLIFYFACTOR,
		false,
	}
}

// IsEmpty returns test Curve Parameters.
func (c *CurveParameters) IsEmpty() bool {
	return c.MitreLimit == 0.0
}
