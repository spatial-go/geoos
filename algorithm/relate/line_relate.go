package relate

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// LineRelate  be used during the relate computation.
type LineRelate struct {
	matrix.LineMatrix
	other matrix.Steric
}

// IntersectionMatrix Gets the IntersectionMatrix for the spatial relationship
// between the input geometries.
func (p *LineRelate) IntersectionMatrix(im *matrix.IntersectionMatrix) *matrix.IntersectionMatrix {
	switch p.other.(type) {
	case matrix.Matrix:
		pr := &PointRelate{p.other.(matrix.Matrix), p.LineMatrix}
		return pr.IntersectionMatrix(im).Transpose()
	case matrix.LineMatrix:
		p.computeLine(im)
		return im
	case matrix.PolygonMatrix:
		p.computePolygon(im)
		return im
	}
	return im
}

func (p *LineRelate) computeLine(im *matrix.IntersectionMatrix) {

	if p.other.(matrix.LineMatrix).IsClosed() {
		mark, ips := IntersectionEdge(p.LineMatrix, p.other.(matrix.LineMatrix))
		if mark {
			if ips.IsOriginal() {
				im.SetAtLeastString("1FF00F102")
				return
			}
			im.SetAtLeastString("0F1FF01F2")
			return
		}
		for _, v := range p.LineMatrix {
			if InPolygon(v, p.other.(matrix.LineMatrix)) {
				im.SetAtLeastString("FF1FF01F2")
				return
			}
		}
		im.SetAtLeastString("FF1FF01F2")
		return
	}
	mark, ips := IntersectionEdge(p.LineMatrix, p.other.(matrix.LineMatrix))
	if mark {
		if ips.IsOriginal() {
			im.SetAtLeastString("1FF00F102")
			return
		}
		im.SetAtLeastString("0F1FF0102")
		return
	}
	im.SetAtLeastString("1FF00F102")
}

func (p *LineRelate) computePolygon(im *matrix.IntersectionMatrix) {
	inRing := -1
	interNum := 0

	for i, v := range p.other.(matrix.PolygonMatrix) {
		mark, ips := IntersectionEdge(p.LineMatrix, matrix.LineMatrix(v))
		if ips.IsOriginal() {
			interNum = 1
		}
		if mark {
			inRing = 1
			break
		}
		if i == 0 {
			if InPolygon(p.LineMatrix[0], v) {
				inRing = 0
			} else {
				inRing = 2
			}
		} else {
			if InPolygon(p.LineMatrix[0], matrix.LineMatrix(v)) {
				if inRing != 2 {
					inRing = 2
					break
				}
			}
		}
	}
	switch inRing {
	case 0:
		im.SetAtLeastString("1FF0FF212")
	case 1:
		if interNum != 0 {
			im.SetAtLeastString("F1FF0F212")
		} else {
			im.SetAtLeastString("1010F0212")
		}
	case 2:
		im.SetAtLeastString("FF1FF0212")
	}

}
