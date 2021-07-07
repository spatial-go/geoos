package overlay

import (
	"strconv"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

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

// Polygon is a closed area. The first edge is the outer ring.
// The others are the holes. Each edge is expected to be closed
// ie. the first point matches the last.
type Polygon struct {
	lines []Line
	rings []*Edge
	edge  *Edge
	rank  int
}

// AddPointWhich add a point in polyon with which.
func (p *Polygon) AddPointWhich(point *Point, which bool) {
	if p.edge == nil || p.edge.nowStatus == CLOSED {
		newCirCuit := Edge{isClockwise: false}
		p.edge = &newCirCuit
		p.edge.points = append(p.edge.points, Point{matrix: matrix.Matrix{point.X(), point.Y()}})
		p.rings = append(p.rings, &newCirCuit)
		p.edge.nowStatus = 0
	} else {
		p.edge.points = append(p.edge.points, Point{matrix: matrix.Matrix{point.X(), point.Y()}})
		// add line
		if len(p.edge.points) > 1 {
			line := Line{isMain: which}
			line.start = &p.edge.points[len(p.edge.points)-2]
			line.end = &p.edge.points[len(p.edge.points)-1]
			p.lines = append(p.lines, line)
		}
	}
}

// AddPoint add a point in polyon.
func (p *Polygon) AddPoint(point *Point) {
	p.AddPointWhich(point, false)
}

// CloseRing close edge to ring.
func (p *Polygon) CloseRing() {
	if p.edge != nil {
		p.edge.nowStatus = CLOSED
		p.edge.SetClockwise()

		// add line
		line := Line{}
		line.start = &p.edge.points[len(p.edge.points)-1]
		line.end = &p.edge.points[0]
		p.lines = append(p.lines, line)
		// if !p.edge.points[len(p.edge.points)-1].Equal(&p.edge.points[0]) {
		// 	p.edge.points = append(p.edge.points, p.edge.points[0])
		// }
	}
}

// Equal Returns p == pol  .
func (p *Polygon) Equal(pol *Polygon) bool {
	for i, v2 := range p.rings {
		for j, v1 := range v2.points {
			if !v1.Equal(&pol.rings[i].points[j]) {
				return false
			}
		}
	}
	return true
}

// ChangeRank change rank .
func (p *Polygon) ChangeRank() {
	if p.rank == MAIN {
		p.rank = CUT
	} else {
		p.rank = MAIN
	}
}

// ToString printf polygon to string
func (p *Polygon) ToString() string {
	var ss string
	for _, c := range p.rings {
		ss = ss + "{"
		for _, v := range c.points {
			ss += "{"
			ss += strconv.FormatFloat(v.X(), 'f', 2, 64)
			ss += ","
			ss += strconv.FormatFloat(v.Y(), 'f', 2, 64)
			ss += "},"
		}
		ss += "}"
	}
	return string(ss)
}
