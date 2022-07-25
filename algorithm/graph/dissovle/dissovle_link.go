package dissovle

import (
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/graph"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// dissovleLink returns edge by link nodes
func dissovleLink(gu graph.Graph, interLline matrix.LineMatrix) (results []matrix.LineMatrix, err error) {
	results = []matrix.LineMatrix{}
	result := matrix.LineMatrix{}

	guNodes := gu.Nodes()
	beUsed := map[int]int{}

	result = append(result, interLline...)
	lenResult := len(result)
	for {
		for j, v := range guNodes {
			if beUsed[j] == 1 {
				continue
			}
			if v.NodeType == graph.CNode || v.NodeType == graph.LNode {
				line := v.Value.(matrix.LineMatrix)
				if line.Equals(interLline) {
					beUsed[j] = 1
					continue
				}
				startPoint := matrix.Matrix(line[0])
				lastPoint := matrix.Matrix(line[len(line)-1])

				if matrix.Matrix(result[len(result)-1]).EqualsExact(startPoint, calc.DefaultTolerance*4) && beUsed[j] < 1 {
					for i, point := range line {
						if i == 0 {
							continue
						}
						result = append(result, point)
					}
					beUsed[j] = 1
					break
				}
				if matrix.Matrix(result[len(result)-1]).EqualsExact(lastPoint, calc.DefaultTolerance*4) && beUsed[j] < 1 {
					for i, point := range line.Reverse() {
						if i == 0 {
							continue
						}
						result = append(result, point)
					}
					beUsed[j] = 1
					break
				}

			} else {
				beUsed[j] = 1
			}
		}
		if result.IsClosed() {
			results = append(results, result)
			result = append(matrix.LineMatrix{}, interLline...)
			lenResult = len(result)
			continue
		}
		if len(beUsed) == len(guNodes) {
			break
		}
		if lenResult == len(result) {
			interLline = interLline.Reverse()
			result = append(matrix.LineMatrix{}, interLline...)
			lenResult = len(result)
			continue
		}
		lenResult = len(result)

	}

	return results, nil
}
