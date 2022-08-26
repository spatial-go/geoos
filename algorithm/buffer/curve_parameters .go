package buffer

import "github.com/spatial-go/geoos/algorithm/calc"

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
		calc.QuadrantSegments,
		calc.CapRound,
		calc.JoinRound,
		calc.MitreLimit,
		calc.SimplifyFactor,
		false,
	}
}

// IsEmpty returns test Curve Parameters.
func (c *CurveParameters) IsEmpty() bool {
	return c.MitreLimit == 0.0
}
