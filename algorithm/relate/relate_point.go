package relate

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Point  be used during the relate computation.
type Point struct {
	matrix.Matrix
	other matrix.Steric
}

// IntersectionMatrix Gets the IntersectionMatrix for the spatial relationship
// between the input geometries.
func (p *Point) IntersectionMatrix(im *matrix.IntersectionMatrix) *matrix.IntersectionMatrix {
	switch p.other.(type) {
	case matrix.Matrix:
		p.computePoint(im)
		return im
	case matrix.LineMatrix:
		p.computeLine(im)
		return im
	case matrix.PolygonMatrix:
		p.computePolygon(im)
		return im
	}
	return im
}

func (p *Point) computePoint(im *matrix.IntersectionMatrix) {
	im.SetAtLeastString("FF1FF00F2")
}

func (p *Point) computeLine(im *matrix.IntersectionMatrix) {
	if p.other.(matrix.LineMatrix).IsClosed() {
		if InLineMatrix(p.Matrix, p.other.(matrix.LineMatrix)) {
			im.SetAtLeastString("0FFFFF1F2")
			return
		}
		im.SetAtLeastString("0FFFFF1F2")
		return
	}
	if InLineMatrix(p.Matrix, p.other.(matrix.LineMatrix)) {
		im.SetAtLeastString("F0FFFF102")
		return
	}
	im.SetAtLeastString("FF0FFF102")
}

func (p *Point) computePolygon(im *matrix.IntersectionMatrix) {
	inRing := -1
	for i, v := range p.other.(matrix.PolygonMatrix) {
		if i == 0 {
			if InLineMatrix(p.Matrix, matrix.LineMatrix(v)) {
				inRing = 1
				break
			} else if InRing(p.Matrix, matrix.LineMatrix(v)) {
				inRing = 0
			}
			inRing = 2
			break
		}
		if InLineMatrix(p.Matrix, matrix.LineMatrix(v)) {
			inRing = 1
			break
		} else if InRing(p.Matrix, matrix.LineMatrix(v)) {
			if inRing != 2 {
				inRing = 2
			}
			break
		}
	}
	switch inRing {
	case 0:
		im.SetAtLeastString("212101212")
	case 1:
		im.SetAtLeastString("F0FFFF212")
	case 2:
		im.SetAtLeastString("212F11FF2")
	}

}
