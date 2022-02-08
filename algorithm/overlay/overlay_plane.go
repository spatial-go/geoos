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
	}
}

// addLines
func (p *Plane) addLines() {
	if len(p.Edge.Vertexes) <= 1 {
		return
	}
	for i := range p.Edge.Vertexes {
		line := Line{}
		start := i
		end := i + 1
		if end >= len(p.Edge.Vertexes) {
			end = 0
		}
		line.Start = &p.Edge.Vertexes[start]
		line.End = &p.Edge.Vertexes[end]
		p.Lines = append(p.Lines, line)
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
		if p.Edge.IsClockwise {
			p.Reverse()
			p.Edge.SetClockwise()
		}
		p.addLines()
	}
}

// Reverse reverse vertexes
func (p *Plane) Reverse() {
	vertexesLength := len(p.Edge.Vertexes) - 1
	for i := 0; i < vertexesLength/2; i++ {
		temp := p.Edge.Vertexes[i]
		p.Edge.Vertexes[i] = p.Edge.Vertexes[vertexesLength-1-i]
		p.Edge.Vertexes[vertexesLength-1-i] = temp
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

// IsVertexInHole ...
func (p *Plane) IsVertexInHole(v *Vertex) (inHole bool) {
	for i, w := range p.Rings {
		if i == 0 {
			continue
		}
		if _, err := SliceContains(w.Vertexes, v); err == nil {
			inHole = true
			break
		}
	}
	return
}
