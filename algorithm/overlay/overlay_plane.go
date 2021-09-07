package overlay

import (
	"strconv"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Plane is a closed area. The first edge is the outer ring.
// The others are the holes. Each edge is expected to be closed
// ie. the first point matches the last.
type Plane struct {
	Lines []Line
	Rings []*Edge
	Edge  *Edge
	Rank  int
}

// AddPointWhich add a point in polyon with which.
func (p *Plane) AddPointWhich(point *Vertex, which bool) {
	if p.Edge == nil || p.Edge.NowStatus == calc.OverlayClosed {
		newCirCuit := Edge{IsClockwise: false}
		p.Edge = &newCirCuit
		p.Edge.Vertexes = append(p.Edge.Vertexes, Vertex{Matrix: matrix.Matrix{point.X(), point.Y()}})
		p.Rings = append(p.Rings, &newCirCuit)
		p.Edge.NowStatus = 0
	} else {
		if len(p.Edge.Vertexes) >= 1 && p.Edge.Vertexes[len(p.Edge.Vertexes)-1].Matrix.Equals(point.Matrix) {
			return
		}
		p.Edge.Vertexes = append(p.Edge.Vertexes, Vertex{Matrix: matrix.Matrix{point.X(), point.Y()}})
		// add line
		if len(p.Edge.Vertexes) > 1 {
			line := Line{IsMain: which}
			line.Start = &p.Edge.Vertexes[len(p.Edge.Vertexes)-2]
			line.End = &p.Edge.Vertexes[len(p.Edge.Vertexes)-1]
			p.Lines = append(p.Lines, line)
		}
	}
}

// AddPoint add a point in polyon.
func (p *Plane) AddPoint(point *Vertex) {
	p.AddPointWhich(point, false)
}

// CloseRing close edge to ring.
func (p *Plane) CloseRing() {
	if p.Edge != nil {
		p.Edge.NowStatus = calc.OverlayClosed
		p.Edge.SetClockwise()

		// add line
		line := Line{}
		line.Start = &p.Edge.Vertexes[len(p.Edge.Vertexes)-1]
		line.End = &p.Edge.Vertexes[0]
		p.Lines = append(p.Lines, line)
		// if !p.edge.points[len(p.edge.points)-1].Equal(&p.edge.points[0]) {
		// 	p.edge.points = append(p.edge.points, p.edge.points[0])
		// }
	}
}

// Equal Returns p == pol  .
func (p *Plane) Equal(pol *Plane) bool {
	for i, v2 := range p.Rings {
		for j, v1 := range v2.Vertexes {
			if !v1.Equals(&pol.Rings[i].Vertexes[j]) {
				return false
			}
		}
	}
	return true
}

// ChangeRank change rank .
func (p *Plane) ChangeRank() {
	if p.Rank == calc.OverlayMain {
		p.Rank = calc.OverlayCut
	} else {
		p.Rank = calc.OverlayMain
	}
}

// ToString printf polygon to string
func (p *Plane) ToString() string {
	var ss string
	for _, c := range p.Rings {
		ss = ss + "{"
		for _, v := range c.Vertexes {
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
