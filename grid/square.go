// Package grid is used to generate grid data.
// include square and hexagon.
package grid

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/space"
)

// Grid ...
type Grid struct {
	Geometry space.Geometry
}

// SquareGrid ,Draw a grid according to the distance, including the given area.
func SquareGrid(bound space.Bound, cellSize float64) (gridGeoms [][]Grid) {
	var (
		minPoint = bound.Min
		maxPoint = bound.Max

		west  = minPoint[0]
		south = minPoint[1]
		east  = maxPoint[0]
		north = maxPoint[1]
	)
	boundWidth := east - west
	boundHeight := north - south

	// Calculate the latitude and longitude corresponding to the length cellSize
	cellWidth := cellSize * (boundWidth / measure.SpheroidDistance(matrix.Matrix{west, south}, matrix.Matrix{east, south}))
	cellHeight := cellSize * (boundHeight / measure.SpheroidDistance(matrix.Matrix{west, north}, matrix.Matrix{west, south}))

	// Round up (including all points)
	columns := math.Ceil(boundWidth / cellWidth)
	rows := math.Ceil(boundHeight / cellHeight)
	deltaX := (columns*cellWidth - boundWidth) / 2
	deltaY := (rows*cellHeight - boundHeight) / 2

	// Draw grid
	currentX := west - deltaX
	for column := int64(0); column < int64(columns); column++ {
		currentY := south - deltaY
		geomRows := []Grid{}
		for row := int64(0); row < int64(rows); row++ {
			point0 := space.Point{currentX, currentY}
			point1 := space.Point{currentX, currentY + cellHeight}
			point2 := space.Point{currentX + cellWidth, currentY + cellHeight}
			point3 := space.Point{currentX + cellWidth, currentY}
			ring := space.Ring{point0, point1, point2, point3, point0}
			polygon := space.Polygon{ring}
			geomRows = append(geomRows, Grid{Geometry: polygon})
			currentY += cellHeight
		}
		gridGeoms = append(gridGeoms, geomRows)
		currentX += cellWidth
	}
	return gridGeoms
}
