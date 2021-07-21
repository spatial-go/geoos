package algorithm

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
	if p.Edge == nil || p.Edge.NowStatus == calc.CLOSED {
		newCirCuit := Edge{IsClockwise: false}
		p.Edge = &newCirCuit
		p.Edge.Vertexs = append(p.Edge.Vertexs, Vertex{Matrix: matrix.Matrix{point.X(), point.Y()}})
		p.Rings = append(p.Rings, &newCirCuit)
		p.Edge.NowStatus = 0
	} else {
		p.Edge.Vertexs = append(p.Edge.Vertexs, Vertex{Matrix: matrix.Matrix{point.X(), point.Y()}})
		// add line
		if len(p.Edge.Vertexs) > 1 {
			line := Line{IsMain: which}
			line.Start = &p.Edge.Vertexs[len(p.Edge.Vertexs)-2]
			line.End = &p.Edge.Vertexs[len(p.Edge.Vertexs)-1]
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
		p.Edge.NowStatus = calc.CLOSED
		p.Edge.SetClockwise()

		// add line
		line := Line{}
		line.Start = &p.Edge.Vertexs[len(p.Edge.Vertexs)-1]
		line.End = &p.Edge.Vertexs[0]
		p.Lines = append(p.Lines, line)
		// if !p.edge.points[len(p.edge.points)-1].Equal(&p.edge.points[0]) {
		// 	p.edge.points = append(p.edge.points, p.edge.points[0])
		// }
	}
}

// Equal Returns p == pol  .
func (p *Plane) Equal(pol *Plane) bool {
	for i, v2 := range p.Rings {
		for j, v1 := range v2.Vertexs {
			if !v1.Equal(&pol.Rings[i].Vertexs[j]) {
				return false
			}
		}
	}
	return true
}

// ChangeRank change rank .
func (p *Plane) ChangeRank() {
	if p.Rank == calc.MAIN {
		p.Rank = calc.CUT
	} else {
		p.Rank = calc.MAIN
	}
}

// ToString printf polygon to string
func (p *Plane) ToString() string {
	var ss string
	for _, c := range p.Rings {
		ss = ss + "{"
		for _, v := range c.Vertexs {
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
