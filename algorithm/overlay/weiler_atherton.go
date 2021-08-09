package overlay

import (
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// SliceContains Returns index of slice.
func SliceContains(list []Vertex, p *Vertex) (int, error) {
	for i, v := range list {
		if v.X() == p.X() && v.Y() == p.Y() {
			return i, nil
		}
	}
	return len(list), algoerr.ErrNotInSlice
}

// AddPointToVertexSlice add point to vertex slice
func AddPointToVertexSlice(edges []*Edge, pStart, pEnd, ip *Vertex) {
	for _, v := range edges {

		if start, err := SliceContains(v.Vertexes, pStart); err == nil {
			end, _ := SliceContains(v.Vertexes, pEnd)

			it := start
			distFromStart := measure.PlanarDistance(ip.Matrix, v.Vertexes[it].Matrix)

			// 处理多个交点
			for it != end && it != len(v.Vertexes) {
				if measure.PlanarDistance(v.Vertexes[it].Matrix, v.Vertexes[start].Matrix) >= distFromStart {
					break
				}
				it++
			}

			circ := v.Vertexes[it:]
			v.Vertexes = append([]Vertex{}, v.Vertexes[:it]...)
			v.Vertexes = append(v.Vertexes, *ip)
			v.Vertexes = append(v.Vertexes, circ...)
			break
		}
	}
}
