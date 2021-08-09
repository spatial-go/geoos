package buffer

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// Centroid Computes the centroid point of a geometry.
func Centroid(geom matrix.Steric) matrix.Matrix {
	cent := &CentroidComputer{}

	if geom == nil || geom.IsEmpty() {
		return nil
	}
	cent.Add(geom)
	centroid := cent.GetCentroid()
	return centroid
}

// CentroidComputer Computes the centroid of a  matrix.Steric of any dimension.
// For collections the centroid is computed for the collection of
// non-empty elements of highest dimension.
// The centroid of an empty matrix.Steric is nil.
type CentroidComputer struct {
	AreaBasePt    matrix.Matrix // the point all triangles are based at
	TriangleCent3 matrix.Matrix // temporary variable to hold centroid of triangle
	Areasum2      float64       // Partial area sum
	Cg3           matrix.Matrix // partial centroid sum

	// data for linear centroid computation, if needed
	LineCentSum matrix.Matrix
	TotalLength float64

	PtCount   int
	PtCentSum matrix.Matrix
}

// GetCentroid Gets the computed centroid.
// returns he computed centroid, or nil if the input is empty
func (c *CentroidComputer) GetCentroid() matrix.Matrix {

	//  The centroid is computed from the highest dimension components present in the input.
	//  I.e. areas dominate lineal matrix.Steric, which dominates points.
	//  Degenerate matrix.Steric are computed using their effective dimension
	//  (e.g. areas may degenerate to lines or points)
	cent := matrix.Matrix{}
	if math.Abs(c.Areasum2) > 0.0 {
		// Input contains areal matrix.Steric
		cent = append(cent, c.Cg3[0]/3/c.Areasum2, c.Cg3[1]/3/c.Areasum2)
	} else if c.TotalLength > 0.0 {
		// Input contains lineal matrix.Steric
		cent = append(cent, c.LineCentSum[0]/c.TotalLength, c.LineCentSum[1]/c.TotalLength)
	} else if c.PtCount > 0 {
		//Input contains matrix.Steric only
		cent = append(cent, c.PtCentSum[0]/float64(c.PtCount), c.PtCentSum[1]/float64(c.PtCount))
	} else {
		return nil
	}
	return cent
}

// Add Adds a Steric to the centroid accumulator.
func (c *CentroidComputer) Add(pt matrix.Steric) {
	switch st := pt.(type) {
	case matrix.Matrix:
		c.AddPoint(st)
	case matrix.LineMatrix:
		c.AddLineSegments(st)
	case matrix.PolygonMatrix:
		c.AddPolygon(st)
	case matrix.Collection:
		for _, v := range st {
			c.Add(v)
		}
	}
}

// AddPoint Adds a point to the point centroid accumulator.
func (c *CentroidComputer) AddPoint(pt matrix.Matrix) {
	if pt.IsEmpty() {
		return
	}
	c.PtCount++
	if c.PtCentSum == nil {
		c.PtCentSum = make(matrix.Matrix, 2)
	}
	c.PtCentSum[0] += pt[0]
	c.PtCentSum[1] += pt[1]
}

// AddLineSegments Adds the line segments  to the linear centroid accumulators.
func (c *CentroidComputer) AddLineSegments(lines matrix.LineMatrix) {
	linelen := 0.0
	if c.LineCentSum == nil {
		c.LineCentSum = matrix.Matrix{0, 0}
	}
	for i := 0; i < len(lines)-1; i++ {
		segmentLen := measure.PlanarDistance(matrix.Matrix(lines[i]), matrix.Matrix(lines[i+1]))
		if segmentLen == 0.0 {
			continue
		}
		linelen += segmentLen
		midX := (lines[i][0] + lines[i+1][0]) / 2
		midY := (lines[i][1] + lines[i+1][1]) / 2
		c.LineCentSum[0] += segmentLen * midX
		c.LineCentSum[1] += segmentLen * midY
	}
	c.TotalLength += linelen
	if linelen == 0.0 && len(lines) > 0 {
		c.AddPoint(lines[0])
	}
}

// AddPolygon Adds the polygon  to the polygon centroid accumulators.
func (c *CentroidComputer) AddPolygon(poly matrix.PolygonMatrix) {
	for i, v := range poly {
		isPositiveArea := false
		if i == 0 {
			if len(v) > 0 {
				c.AreaBasePt = v[0]
			}
			isPositiveArea = !measure.IsCCW(v)
		} else {
			isPositiveArea = measure.IsCCW(v)
		}
		for i := 0; i < len(v)-1; i++ {
			c.addTriangle(matrix.Matrix(c.AreaBasePt), v[i], v[i+1], isPositiveArea)
		}
		c.AddLineSegments(v)
	}
}

// addTriangle Adds the Triangle  to the Triangle centroid accumulators.
func (c *CentroidComputer) addTriangle(p0, p1, p2 matrix.Matrix, isPositiveArea bool) {
	sign := 1.0
	if !isPositiveArea {
		sign = -1.0
	}
	c.TriangleCent3 = matrix.Matrix{p0[0] + p1[0] + p2[0], p0[1] + p1[1] + p2[1]}
	area2 := (p1[0]-p0[0])*(p2[1]-p0[1]) -
		(p2[0]-p0[0])*(p1[1]-p0[1])
	if c.Cg3 == nil {
		c.Cg3 = matrix.Matrix{0, 0}
	}
	c.Cg3[0] += sign * area2 * c.TriangleCent3[0]
	c.Cg3[1] += sign * area2 * c.TriangleCent3[1]
	c.Areasum2 += sign * area2
}
