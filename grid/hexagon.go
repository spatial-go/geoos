package grid

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/space"
)

// const ...
const (
	// Angle of sin 60 = 0.866025403785
	Sin60 = 0.866025403785
	Cos60 = 0.5
)

// HexagonGrid Draw a grid according to the distance, including the given area
func HexagonGrid(bound space.Bound, cellSize float64) (gridGeoms [][]Grid) {
	var (
		minPoint = bound.Min
		maxPoint = bound.Max
		west     = minPoint[0]
		south    = minPoint[1]
		east     = maxPoint[0]
		north    = maxPoint[1]

		Sin60 = Sin60
		Cos60 = Cos60
	)
	boundHeight := north - south
	boundWidth := east - west

	// Calculate the latitude and longitude corresponding to the length cellSize.
	cellHeight := cellSize * (boundHeight / measure.SpheroidDistance(matrix.Matrix{west, north}, matrix.Matrix{west, south}))
	cellWidth := cellSize * (boundWidth / measure.SpheroidDistance(matrix.Matrix{west, south}, matrix.Matrix{east, south}))

	// Get the number of rows and columns of the grid to be drawn in the bound range
	columns := math.Ceil(boundWidth/(cellHeight+cellHeight*Cos60) + 1)
	rows := math.Ceil(boundHeight/(2*cellWidth*Sin60) + 1)

	// CurrentX,currentY are the center point coordinates of the hexagon.
	// Draw the hexagon with six points calculated from the triangle relationship and the center point..
	// The order of drawing hexagons is from bottom to top, from left to right
	currentX := west
	for column := int64(0); column < int64(columns); column++ {
		currentY := south
		if column%2 != 0 {
			currentY = south + Sin60*cellHeight
		}
		geomRows := []Grid{}
		for row := int64(0); row < int64(rows); row++ {
			// The directions of the point 0、1、2、3、4、5 of the hexagon are 1、3、5、7、9、1 o'clock direction by turn.
			point0 := space.Point{currentX + Cos60*cellWidth, currentY + Sin60*cellHeight}
			point1 := space.Point{currentX + cellWidth, currentY}
			point2 := space.Point{currentX + Cos60*cellWidth, currentY - Sin60*cellHeight}
			point3 := space.Point{currentX - Cos60*cellWidth, currentY - Sin60*cellHeight}
			point4 := space.Point{currentX - cellWidth, currentY}
			point5 := space.Point{currentX - Cos60*cellWidth, currentY + Sin60*cellHeight}
			ring := space.Ring{point0, point1, point2, point3, point4, point5, point0}
			polygon := space.Polygon{ring}
			geomRows = append(geomRows, Grid{Geometry: polygon})
			currentY = currentY + 2*cellHeight*Sin60
		}
		gridGeoms = append(gridGeoms, geomRows)
		currentX = currentX + cellWidth + cellWidth*Cos60
	}
	return
}
