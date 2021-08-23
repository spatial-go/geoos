package relate

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// PointRelate  be used during the relate computation.
type PointRelate struct {
	matrix.Matrix
	other matrix.Steric
}

// IntersectionMatrix Gets the IntersectionMatrix for the spatial relationship
// between the input geometries.
func (p *PointRelate) IntersectionMatrix(im *matrix.IntersectionMatrix) *matrix.IntersectionMatrix {
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

func (p *PointRelate) computePoint(im *matrix.IntersectionMatrix) {
	im.SetAtLeastString("0FFFFFFF2")
}

func (p *PointRelate) computeLine(im *matrix.IntersectionMatrix) {
	if p.other.(matrix.LineMatrix).IsClosed() {
		if inLine, inVer := InLineVertex(p.Matrix, p.other.(matrix.LineMatrix)); inLine {
			if inVer {
				im.SetAtLeastString("F0FFFF102")
				return
			}
			im.SetAtLeastString("0FFFFF102")
			return
		}
		if InLineMatrix(p.Matrix, p.other.(matrix.LineMatrix)) {
			im.SetAtLeastString("0FFFFF1F2")
			return
		}
		im.SetAtLeastString("FF0FFF1F2")
		return
	}
	if inLine, inVer := InLineVertex(p.Matrix, p.other.(matrix.LineMatrix)); inLine {
		if inVer {
			im.SetAtLeastString("F0FFFF102")
			return
		}
		im.SetAtLeastString("0FFFFF102")
		return
	}
	if InLineMatrix(p.Matrix, p.other.(matrix.LineMatrix)) {
		im.SetAtLeastString("0FFFFF102")
		return
	}
	im.SetAtLeastString("FF0FFF102")
}

func (p *PointRelate) computePolygon(im *matrix.IntersectionMatrix) {
	inRing := -1
	inHoles := 0
	for i, v := range p.other.(matrix.PolygonMatrix) {
		if i == 0 {
			if InLineMatrix(p.Matrix, matrix.LineMatrix(v)) {
				inRing = 1
			} else if InPolygon(p.Matrix, matrix.LineMatrix(v)) {
				inRing = 0
			} else {
				inRing = 2
			}
		} else {
			if InLineMatrix(p.Matrix, matrix.LineMatrix(v)) {
				inRing = 1
				break
			} else if InPolygon(p.Matrix, matrix.LineMatrix(v)) {
				if inRing != 2 {
					inRing = 2
					inHoles = 1
					break
				}
			}
		}
	}
	switch inRing {
	case 0:
		im.SetAtLeastString("0FFFFF212")
	case 1:
		im.SetAtLeastString("F0FFFF212")
	case 2:
		if inHoles == 1 {
			im.SetAtLeastString("FF0FFF212")
		}
		im.SetAtLeastString("FF0FFF212")
	}

}
