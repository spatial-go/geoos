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

		if start, err := SliceContains(v.Vertexs, pStart); err == nil {
			end, _ := SliceContains(v.Vertexs, pEnd)

			it := start
			distFromStart := measure.PlanarDistance(ip.Matrix, v.Vertexs[it].Matrix)

			// 处理多个交点
			for it != end && it != len(v.Vertexs) {
				if measure.PlanarDistance(v.Vertexs[it].Matrix, v.Vertexs[start].Matrix) >= distFromStart {
					break
				}
				it++
			}

			circ := v.Vertexs[it:]
			v.Vertexs = append([]Vertex{}, v.Vertexs[:it]...)
			v.Vertexs = append(v.Vertexs, *ip)
			v.Vertexs = append(v.Vertexs, circ...)
			break
		}
	}
}
