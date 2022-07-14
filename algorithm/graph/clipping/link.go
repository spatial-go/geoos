package clipping

import (
	"log"
	"os"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/graph"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

// link returns edge by link nodes
func link(gu, gi graph.Graph) (results []matrix.LineMatrix, err error) {
	results = []matrix.LineMatrix{}
	result := matrix.LineMatrix{}

	attempts := 0
	guNodes := gu.Nodes()
	giNodes := []*graph.Node{}
	if gi != nil {
		giNodes = gi.Nodes()
	}
	beUsed := map[int]int{}
	lenBeUsed := len(beUsed)
	beUsedGi := map[int]int{}
	for {
		for j, v := range guNodes {
			if v.NodeType == graph.CNode || v.NodeType == graph.LNode {
				line := v.Value.(matrix.LineMatrix)
				startPoint := matrix.Matrix(line[0])
				lastPoint := matrix.Matrix(line[len(line)-1])
				if len(result) == 0 {
					if beUsed[j] < 1 {
						result = append(result, line...)
						beUsed[j] = 1
						break
					}
				} else {
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
				}
			}
		}
		if result.IsClosed() || len(result) == 0 {
			if attempts < 100 {
				if len(result) > 0 {
					results = append(results, result)
				}
				result = matrix.LineMatrix{}
				if len(beUsed) >= len(guNodes)-1 {
					break
				}
				continue
			} else {
				results = append(results, result)
				break
			}
		}
		if len(beUsed) <= lenBeUsed {
			giUsed := 0
			for j, v := range giNodes {
				if v.NodeType == graph.CNode || v.NodeType == graph.LNode {
					line := v.Value.(matrix.LineMatrix)
					startPoint := matrix.Matrix(line[0])
					lastPoint := matrix.Matrix(line[len(line)-1])

					if matrix.Matrix(result[len(result)-1]).EqualsExact(startPoint, calc.DefaultTolerance*4) && beUsedGi[j] < 1 {
						for i, point := range line {
							if i == 0 {
								continue
							}
							result = append(result, point)
						}
						beUsedGi[j] = 1
						giUsed++
						break
					}
					if matrix.Matrix(result[len(result)-1]).EqualsExact(lastPoint, calc.DefaultTolerance*4) && beUsedGi[j] < 1 {
						for i, point := range line.Reverse() {
							if i == 0 {
								continue
							}
							result = append(result, point)
						}
						beUsedGi[j] = 1
						giUsed++
						break
					}
				}
			}

			if giUsed > 0 {
				continue
			}
			if attempts < 100 {
				result = matrix.LineMatrix{}
				continue
			}
			if !geoos.GeoosTestTag {
				geom := space.Collection{}
				for _, v := range guNodes {
					geom = append(geom, space.TransGeometry(v.Value))
				}
				for _, v := range giNodes {
					geom = append(geom, space.TransGeometry(v.Value))
				}
				writeGeom(dir+"data_link.geojson", geom)
			}
			return nil, algorithm.ErrWrongLink
		}
		lenBeUsed = len(beUsed)
	}
	return results, nil
}

func linkmerge(gu, gi graph.Graph) (results []matrix.LineMatrix, err error) {
	results = []matrix.LineMatrix{}
	result := matrix.LineMatrix{}

	guNodes := gu.Nodes()
	// giNodes := []*graph.Node{}
	// if gi != nil {
	// 	giNodes = gi.Nodes()
	// }
	beUsed := map[int]int{}
	currentNode := 0
	for {
		for j, v := range guNodes {
			if v.NodeType == graph.CNode || v.NodeType == graph.LNode {
				line := v.Value.(matrix.LineMatrix)
				startPoint := matrix.Matrix(line[0])
				lastPoint := matrix.Matrix(line[len(line)-1])
				if len(result) == 0 {
					if beUsed[j] < 1 {
						result = append(result, line...)
						beUsed[j] = 1
						currentNode = j
						break
					}
				} else {
					if matrix.Matrix(result[len(result)-1]).EqualsExact(startPoint, calc.DefaultTolerance*4) && beUsed[j] < 1 {
						for i, point := range line {
							if i == 0 {
								continue
							}
							result = append(result, point)
						}
						beUsed[j] = 1
						currentNode = j
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
						currentNode = j
						break
					}
				}
			}
			currentNode = j
		}

		if len(result) > 0 && currentNode < len(guNodes)-1 {
			// giUsed := 0
			// for _, v := range giNodes {
			// 	if v.NodeType == graph.CNode || v.NodeType == graph.LNode {
			// 		line := v.Value.(matrix.LineMatrix)
			// 		startPoint := matrix.Matrix(line[0])
			// 		lastPoint := matrix.Matrix(line[len(line)-1])

			// 		if matrix.Matrix(result[len(result)-1]).EqualsExact(startPoint, calc.DefaultTolerance*4) {
			// 			for i, point := range line {
			// 				if i == 0 {
			// 					continue
			// 				}
			// 				result = append(result, point)
			// 			}
			// 			giUsed++
			// 			break
			// 		}
			// 		if matrix.Matrix(result[len(result)-1]).EqualsExact(lastPoint, calc.DefaultTolerance*4) {
			// 			for i, point := range line.Reverse() {
			// 				if i == 0 {
			// 					continue
			// 				}
			// 				result = append(result, point)
			// 			}
			// 			giUsed++
			// 			break
			// 		}
			// 	}
			// }
			continue
		} else if len(result) > 0 && currentNode >= len(guNodes)-1 {
			results = append(results, result)
			result = matrix.LineMatrix{}
			continue
		}

		if currentNode >= len(guNodes)-1 {
			for j, v := range guNodes {
				if v.NodeType == graph.CNode || v.NodeType == graph.LNode {
					if beUsed[j] != 1 {
						line := v.Value.(matrix.LineMatrix)
						results = append(results, line)
					}
				}
			}
			break
		}

	}
	return results, nil
}

func writeGeom(filename string, geom space.Geometry) {
	data, _ := geojson.NewGeometry(geom).MarshalJSON()
	if file, err := os.Create(filename); err != nil {
		log.Println(err)
	} else {
		if _, err := file.WriteString(
			`{
		"type": "FeatureCollection",
		"crs": {
			"type": "name",
			"properties": {
				"name": "bd09"
			}
		},
		"features": [
			{
				"_id": "qVxKbM",
				"type": "Feature",
				"geometry": `,
		); err != nil {
			log.Println(err)
		}
		_, _ = file.Write(data)
		_, _ = file.WriteString(
			` }
		]
		}`,
		)
	}
}
