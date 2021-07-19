package space

import "sort"

// Interior  interface Computes a location of an interior point in a Geometry.
type Interior interface {
	InteriorPoint() Point
	Add(geom Geometry)
}

// InteriorPoint Computes a location of an interior point in a Geometry.
func InteriorPoint(geom Geometry) (interiorPt Point) {
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

func effectiveDimension(geom Geometry) int {
	dim := -1
	if geom.GeoJSONType() == TypeCollection {
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
	centroid, interiorPoint Point
	minDistance             float64
	geom                    Geometry
}

// Add Tests the point(s) defined by a Geometry for the best inside point.
// If a Geometry is not of dimension 0 it is not tested.
func (ip *InteriorPointPoint) Add(geom Geometry) {
	ip.geom = geom
	ip.centroid = ip.geom.Centroid()
	if geom.GeoJSONType() == TypePoint {
		ip.add(geom.(Point))
	} else if geom.GeoJSONType() == TypeCollection {
		for _, v := range geom.(Collection) {
			ip.add(v.(Point))
		}
	}
}

func (ip *InteriorPointPoint) add(point Point) {
	dist, _ := point.Distance(ip.centroid)
	if ip.minDistance == 0.0 || dist < ip.minDistance {
		ip.interiorPoint = point
		ip.minDistance = dist
	}
}

// InteriorPoint returns InteriorPoint.
func (ip *InteriorPointPoint) InteriorPoint() Point {
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
func (ip *InteriorPointLine) Add(geom Geometry) {
	ip.geom = geom
	ip.centroid = ip.geom.Centroid()

	if geom.GeoJSONType() == TypeLineString {
		ip.addInterior(geom.(LineString))
	} else if geom.GeoJSONType() == TypeCollection {
		for _, v := range geom.(Collection) {
			ip.addInterior(v.(LineString))
		}
	}
	if ip.interiorPoint == nil {
		ip.addEndpoint(geom)
	}
}

func (ip *InteriorPointLine) addInterior(pts LineString) {
	for _, v := range pts {
		ip.add(Point(v))
	}
}

// addEndpoint Tests the endpoint vertices
// defined by a linear Geometry for the best inside point.
func (ip *InteriorPointLine) addEndpoint(geom Geometry) {
	if geom.GeoJSONType() == TypeLineString {
		ip.addEndpoints(geom.(LineString))
	} else if geom.GeoJSONType() == TypeCollection {
		for _, v := range geom.(Collection) {
			ip.addEndpoints(v.(LineString))
		}
	}
}
func (ip *InteriorPointLine) addEndpoints(pts LineString) {
	ip.add(pts[0])
	ip.add(pts[len(pts)-1])
}

// InteriorPointArea Computes a point in the interior of an areal geometry.
// The point will lie in the geometry interior
// in all except certain pathological cases.
type InteriorPointArea struct {
	InteriorPointPoint
	maxWidth, interiorSectionWidth float64
	interiorPointY                 float64
}

func avg(a, b float64) float64 {
	return (a + b) / 2.0
}

// Add  Processes a geometry to determine the best interior point for
// all component polygons.
func (ip *InteriorPointArea) Add(geom Geometry) {
	ip.geom = geom
	ip.centroid = ip.geom.Centroid()

	ip.interiorPointY = ScanLineY(geom.(Polygon))

	if geom.GeoJSONType() == TypePolygon {
		ip.processPolygon(geom.(Polygon))
	} else if geom.GeoJSONType() == TypeCollection {
		for _, v := range geom.(Collection) {
			ip.processPolygon(v.(Polygon))
		}
	}
}

// processPolygon Computes an interior point of a component Polygon
// and updates current best interior point if appropriate.
func (ip *InteriorPointArea) processPolygon(polygon Polygon) {
	ip.process()
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
func (ip *InteriorPointArea) process() {
	/**
	 * This results in returning a null Coordinate
	 */
	if ip.geom.IsEmpty() {
		return
	}
	/**
	 * set default interior point in case polygon has zero area
	 */
	crossings := []float64{}
	ip.scanRing(ip.geom.(Polygon).Shell(), crossings)
	for _, v := range ip.geom.(Polygon).Holes() {
		ip.scanRing(v, crossings)
	}
	ip.findBestMidpoint(crossings)
}

func (ip *InteriorPointArea) scanRing(ring Ring, crossings []float64) {
	// skip rings which don't cross scan line
	if !ip.intersectsHorizontalLine(ring.Bound().Min, ring.Bound().Max, ip.interiorPointY) {
		return
	}

	for i, v := range ring {
		if i == 0 {
			continue
		}
		ptPrev := ring[i-1]
		pt := v
		ip.addEdgeCrossing(ptPrev, pt, ip.interiorPointY, crossings)
	}
}

func (ip *InteriorPointArea) addEdgeCrossing(p0, p1 Point, scanY float64, crossings []float64) {
	// skip non-crossing segments
	if !ip.intersectsHorizontalLine(p0, p1, scanY) {
		return
	}
	if !ip.isEdgeCrossingCounted(p0, p1, scanY) {
		return
	}

	// edge intersects scan line, so add a crossing
	xInt := ip.intersection(p0, p1, scanY)
	crossings = append(crossings, xInt)
	//checkIntersectionDD(p0, p1, scanY, xInt);
}

// findBestMidpoint Finds the midpoint of the widest interior section.
// Sets the {@link #interiorPoint} location
// and the {@link #interiorSectionWidth}
func (ip *InteriorPointArea) findBestMidpoint(crossings []float64) {
	// zero-area polygons will have no crossings
	if len(crossings) == 0 {
		return
	}

	sort.Sort(sort.Float64Slice(crossings))
	/*
	 * Entries in crossings list are expected to occur in pairs representing a
	 * section of the scan line interior to the polygon (which may be zero-length)
	 */
	for i := 0; i < len(crossings); i += 2 {
		x1 := crossings[i]
		// crossings count must be even so this should be safe
		x2 := crossings[i+1]

		width := x2 - x1
		if width > ip.interiorSectionWidth {
			ip.interiorSectionWidth = width
			interiorPointX := avg(x1, x2)
			ip.interiorPoint = Point{interiorPointX, ip.interiorPointY}
		}
	}
}

// isEdgeCrossingCounted Tests if an edge intersection contributes to the crossing count.
// Some crossing situations are not counted,
//  to ensure that the list of crossings
//  captures strict inside/outside topology.
func (ip *InteriorPointArea) isEdgeCrossingCounted(p0, p1 Point, scanY float64) bool {
	y0 := p0.Y()
	y1 := p1.Y()
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
func (ip *InteriorPointArea) intersection(p0, p1 Point, Y float64) float64 {
	x0 := p0.X()
	x1 := p1.X()

	if x0 == x1 {
		return x0
	}
	// Assert: segDX is non-zero, due to previous equality test
	segDX := x1 - x0
	segDY := p1.Y() - p0.Y()
	m := segDY / segDX
	x := x0 + ((Y - p0.Y()) / m)
	return x
}

// Tests if a line segment intersects a horizontal line.
func (ip *InteriorPointArea) intersectsHorizontalLine(p0, p1 Point, y float64) bool {
	// both ends above?
	if p0.Y() > y && p1.Y() > y {
		return false
	}
	// both ends below?
	if p0.Y() < y && p1.Y() < y {
		return false
	}
	// segment must intersect line
	return true
}

// ScanLineYOrdinateFinder Finds a safe scan line Y ordinate by projecting
// the polygon segments
type ScanLineYOrdinateFinder struct {
	poly Polygon

	centreY, hiY, loY float64
}

// ScanLineY Finds a safe scan line Y ordinate by projecting
// the polygon segments
func ScanLineY(poly Polygon) float64 {
	finder := &ScanLineYOrdinateFinder{poly: poly}
	finder.hiY = poly.Bound().Top()
	finder.loY = poly.Bound().Bottom()
	finder.centreY = avg(finder.loY, finder.hiY)
	return finder.ScanLineY()
}

// ScanLineY Finds a safe scan line Y ordinate by projecting
// the polygon segments
func (s *ScanLineYOrdinateFinder) ScanLineY() float64 {
	s.process(LineString(s.poly.Shell()))
	for _, v := range s.poly {
		s.process(v)
	}
	scanLineY := avg(s.hiY, s.loY)
	return scanLineY
}

func (s *ScanLineYOrdinateFinder) process(line LineString) {

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
