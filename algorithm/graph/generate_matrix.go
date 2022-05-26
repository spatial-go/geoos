package graph

import "github.com/spatial-go/geoos/algorithm/matrix"

// GenerateSteric ...
func GenerateSteric(g Graph) (matrix.Steric, error) {
	result := matrix.Collection{}
	for _, v := range g.Nodes() {
		if !v.Stat {
			continue
		}
		if v.NodeType == PNode && g.Degree(v.Index) >= 1 {
			continue
		}
		result = append(result, v.Value)
	}
	if len(result) == 1 {
		return result[0], nil
	}
	return result, nil
}
