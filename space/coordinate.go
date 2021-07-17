package space

import (
	"fmt"
)

// Coordinate coord
type Coordinate struct {
	X float64
	Y float64
	Z float64
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%f %f", c.X, c.Y)
}
