package buffer

import (
	"sort"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// Interior  interface Computes a location of an interior point in a Geometry.
type Interior interface {
	InteriorPoint() matrix.Matrix
	Add(geom matrix.Steric)
}

// InteriorPoint Computes a location of an interior point in a Geometry.
func InteriorPoint(geom matrix.Steric) (interiorPt matrix.Matrix) {
	if geom.IsEmpty() {
		return nil
	}

	dim := effectiveDimension(geom)
	// this should not happen, but just in case...
	if dim < 0 {
		return nil
	}
	var interior Interior
	if dim == 0 {
		interior = &InteriorPointPoint{}
	} else if dim == 1 {
		interior = &InteriorPointLine{}
	} else {
		interior = &InteriorPointArea{}
	}
	interior.Add(geom)
	return interior.InteriorPoint()
}

func effectiveDimension(geom matrix.Steric) int {
	dim := -1
	if _, ok := geom.(matrix.Collection); ok {
		return dim
	}
	if !geom.IsEmpty() {
		elemDim := geom.Dimensions()
		if elemDim > dim {
			dim = elemDim
		}
	}
	return dim
}

// InteriorPointPoint Computes a point in the interior of an point geometry.
// Find a point which is closest to the centroid of the geometry.
type InteriorPointPoint struct {
	centroid, interiorPoint matrix.Matrix
	minDistance             float64
	geom                    matrix.Steric
}

// Add Tests the point(s) defined by a Geometry for the best inside point.
// If a Geometry is not of dimension 0 it is not tested.
func (ip *InteriorPointPoint) Add(point matrix.Steric) {
	ip.geom = point
	ip.centroid = Centroid(point)
	switch p := point.(type) {
	case matrix.Matrix:
		ip.addPoint(p)
	case matrix.Collection:
		for _, v := range p {
			ip.addPoint(v.(matrix.Matrix))
		}
	}
}

func (ip *InteriorPointPoint) addPoint(point matrix.Matrix) {
	dist := measure.PlanarDistance(point, ip.centroid)
	if ip.minDistance == 0.0 || dist < ip.minDistance {
		ip.interiorPoint = point
		ip.minDistance = dist
	}
}

// InteriorPoint returns InteriorPoint.
func (ip *InteriorPointPoint) InteriorPoint() matrix.Matrix {
	return ip.interiorPoint
}

// InteriorPointLine Computes a point in the interior of an linear geometry.
// Find an interior vertex which is closest to
// the centroid of the linestring.
// If there is no interior vertex, find the endpoint which is
//  closest to the centroid.
type InteriorPointLine struct {
	InteriorPointPoint
}

// Add Tests the interior vertices (if any)
// defined by a linear Geometry for the best inside point.
func (ip *InteriorPointLine) Add(line matrix.Steric) {
	ip.geom = line
	ip.centroid = Centroid(line)
	switch p := line.(type) {
	case matrix.LineMatrix:
		ip.addInterior(p)
	case matrix.Collection:
		for _, v := range p {
			ip.addInterior(v.(matrix.LineMatrix))
		}
	}
	if ip.interiorPoint == nil {
		ip.addEndpoint(line)
	}
}

func (ip *InteriorPointLine) addInterior(pts matrix.LineMatrix) {
	for _, v := range pts {
		ip.addPoint(v)
	}
}

// addEndpoint Tests the endpoint vertices
// defined by a linear Geometry for the best inside point.
func (ip *InteriorPointLine) addEndpoint(line matrix.Steric) {
	switch p := line.(type) {
	case matrix.LineMatrix:
		ip.addEndpoints(p)
	case matrix.Collection:
		for _, v := range p {
			ip.addEndpoints(v.(matrix.LineMatrix))
		}
	}
}
func (ip *InteriorPointLine) addEndpoints(pts matrix.LineMatrix) {
	ip.addPoint(pts[0])
	ip.addPoint(pts[len(pts)-1])
}

// InteriorPointArea Computes a point in the interior of an areal geometry.
// The point will lie in the geometry interior
// in all except certain pathological cases.
type InteriorPointArea struct {
	InteriorPointPoint
	maxWidth, interiorSectionWidth float64
	interiorPointY, centreY        float64
}

func avg(a, b float64) float64 {
	return (a + b) / 2.0
}

// Add  Processes a geometry to determine the best interior point for
// all component polygons.
func (ip *InteriorPointArea) Add(poly matrix.Steric) {
	ip.geom = poly
	ip.centroid = Centroid(poly)

	switch p := poly.(type) {
	case matrix.PolygonMatrix:
		ip.processPolygon(p)
	case matrix.Collection:
		for _, v := range p {
			ip.processPolygon(v.(matrix.PolygonMatrix))
		}
	}
}

// ScanLineY ...
func (ip *InteriorPointArea) ScanLineY(polygon matrix.PolygonMatrix) float64 {
	b := polygon.Bound()
	loY := b[0][1]
	hiY := b[1][1]
	ip.centreY = (loY + hiY) / 2.0
	for _, v := range polygon {
		loY, hiY = ip.processY(v, loY, hiY)
	}
	scanLineY := avg(hiY, loY)
	return scanLineY
}
func (ip *InteriorPointArea) processY(ring matrix.LineMatrix, loY, hiY float64) (float64, float64) {
	for _, v := range ring {
		y := v[1]
		loY, hiY = ip.updateInterval(loY, hiY, y)
	}
	return loY, hiY
}

func (ip *InteriorPointArea) updateInterval(loY, hiY, y float64) (float64, float64) {
	if y <= ip.centreY {
		if y > loY {
			loY = y
		}
	} else if y > ip.centreY {
		if y < hiY {
			hiY = y
		}
	}
	return loY, hiY
}

// processPolygon Computes an interior point of a component Polygon
// and updates current best interior point if appropriate.
func (ip *InteriorPointArea) processPolygon(polygon matrix.PolygonMatrix) {
	ip.interiorPointY = ip.ScanLineY(polygon)
	ip.process(polygon)
	width := ip.Width()
	if width > ip.maxWidth {
		ip.maxWidth = width
		ip.interiorPoint = ip.InteriorPoint()
	}
}

// Width Gets the width of the scanline section containing the interior point.
// Used to determine the best point to use.
func (ip *InteriorPointArea) Width() float64 {
	return ip.interiorSectionWidth
}

// process Compute the interior point.
func (ip *InteriorPointArea) process(polygon matrix.PolygonMatrix) {
	// This results in returning a nil Coordinate
	if polygon.IsEmpty() {
		return
	}
	// set default interior point in case polygon has zero area
	crossings := []float64{}
	for _, v := range polygon {
		crossings = ip.scanRing(v, crossings)
	}
	ip.findBestMidpoint(crossings)
}

func (ip *InteriorPointArea) scanRing(ring matrix.LineMatrix, crossings []float64) []float64 {
	// skip rings which don't cross scan line
	if !ip.intersectsHorizontalLine(ring.Bound()[0], ring.Bound()[1], ip.interiorPointY) {
		return crossings
	}

	for i, v := range ring {
		if i == 0 {
			continue
		}
		ptPrev := ring[i-1]
		pt := v
		crossings = ip.addEdgeCrossing(ptPrev, pt, ip.interiorPointY, crossings)
	}
	return crossings
}

func (ip *InteriorPointArea) addEdgeCrossing(p0, p1 matrix.Matrix, scanY float64, crossings []float64) []float64 {
	// skip non-crossing segments
	if !ip.intersectsHorizontalLine(p0, p1, scanY) {
		return crossings
	}
	if !ip.isEdgeCrossingCounted(p0, p1, scanY) {
		return crossings
	}

	// edge intersects scan line, so add a crossing
	xInt := ip.intersection(p0, p1, scanY)
	crossings = append(crossings, xInt)
	return crossings
	//checkIntersectionDD(p0, p1, scanY, xInt);
}

// findBestMidpoint Finds the midpoint of the widest interior section.
func (ip *InteriorPointArea) findBestMidpoint(crossings []float64) {
	// zero-area polygons will have no crossings
	if len(crossings) == 0 {
		return
	}

	sort.Float64sAreSorted(sort.Float64Slice(crossings))

	// Entries in crossings list are expected to occur in pairs representing a
	// section of the scan line interior to the polygon (which may be zero-length)
	for i := 0; i < len(crossings); i += 2 {
		x1 := crossings[i]
		// crossings count must be even so this should be safe
		x2 := crossings[i+1]

		width := x2 - x1
		if width > ip.interiorSectionWidth {
			ip.interiorSectionWidth = width
			interiorPointX := avg(x1, x2)
			ip.interiorPoint = matrix.Matrix{interiorPointX, ip.interiorPointY}
		}
	}
}

// isEdgeCrossingCounted Tests if an edge intersection contributes to the crossing count.
// Some crossing situations are not counted,
//  to ensure that the list of crossings
//  captures strict inside/outside topology.
func (ip *InteriorPointArea) isEdgeCrossingCounted(p0, p1 matrix.Matrix, scanY float64) bool {
	y0 := p0[1]
	y1 := p1[1]
	// skip horizontal lines
	if y0 == y1 {
		return false
	}
	// handle cases where vertices lie on scan-line
	// downward segment does not include start point
	if y0 == scanY && y1 < scanY {
		return false
	}
	// upward segment does not include endpoint
	if y1 == scanY && y0 < scanY {
		return false
	}
	return true
}

// intersection Computes the intersection of a segment with a horizontal line.
// The segment is expected to cross the horizontal line
// - this condition is not checked.
func (ip *InteriorPointArea) intersection(p0, p1 matrix.Matrix, Y float64) float64 {
	x0 := p0[0]
	x1 := p1[0]

	if x0 == x1 {
		return x0
	}
	// Assert: segDX is non-zero, due to previous equality test
	segDX := x1 - x0
	segDY := p1[1] - p0[1]
	m := segDY / segDX
	x := x0 + ((Y - p0[1]) / m)
	return x
}

// Tests if a line segment intersects a horizontal line.
func (ip *InteriorPointArea) intersectsHorizontalLine(p0, p1 matrix.Matrix, y float64) bool {
	// both ends above?
	if p0[1] > y && p1[1] > y {
		return false
	}
	// both ends below?
	if p0[1] < y && p1[1] < y {
		return false
	}
	// segment must intersect line
	return true
}

// ScanLineYOrdinateFinder Finds a safe scan line Y ordinate by projecting
// the polygon segments
type ScanLineYOrdinateFinder struct {
	poly matrix.PolygonMatrix

	centreY, hiY, loY float64
}

// ScanLineY Finds a safe scan line Y ordinate by projecting
// the polygon segments
func ScanLineY(poly matrix.PolygonMatrix) float64 {
	finder := &ScanLineYOrdinateFinder{poly: poly}
	finder.hiY = poly.Bound()[1][1]
	finder.loY = poly.Bound()[0][1]
	finder.centreY = avg(finder.loY, finder.hiY)
	return finder.ScanLineY()
}

// ScanLineY Finds a safe scan line Y ordinate by projecting
// the polygon segments
func (s *ScanLineYOrdinateFinder) ScanLineY() float64 {
	s.process(matrix.LineMatrix(s.poly[0]))
	for _, v := range s.poly {
		s.process(v)
	}
	scanLineY := avg(s.hiY, s.loY)
	return scanLineY
}

func (s *ScanLineYOrdinateFinder) process(line matrix.LineMatrix) {

	for _, v := range line {
		s.updateInterval(v[1])
	}
}

func (s *ScanLineYOrdinateFinder) updateInterval(y float64) {
	if y <= s.centreY {
		if y > s.loY {
			s.loY = y
		}
	} else if y > s.centreY {
		if y < s.hiY {
			s.hiY = y
		}
	}
}
