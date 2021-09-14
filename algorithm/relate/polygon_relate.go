package relate

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// PolygonRelate  be used during the relate computation.
type PolygonRelate struct {
	matrix.PolygonMatrix
	other matrix.Steric
}

// IntersectionMatrix Gets the IntersectionMatrix for the spatial relationship
// between the input geometries.
func (p *PolygonRelate) IntersectionMatrix(im *matrix.IntersectionMatrix) *matrix.IntersectionMatrix {
	switch p.other.(type) {
	case matrix.Matrix:
		pr := &PointRelate{p.other.(matrix.Matrix), p.PolygonMatrix}
		return pr.IntersectionMatrix(im).Transpose()
	case matrix.LineMatrix:
		lr := &LineRelate{p.other.(matrix.LineMatrix), p.PolygonMatrix}
		return lr.IntersectionMatrix(im).Transpose()
	case matrix.PolygonMatrix:
		p.computePolygon(im)
		return im
	}
	return im
}

func (p *PolygonRelate) computePolygon(im *matrix.IntersectionMatrix) {
	inRing := -1
	ipNum := 0
	for i, v := range p.other.(matrix.PolygonMatrix) {
		l := p.PolygonMatrix[0]
		if mark, ips := IntersectionEdge(l, v); mark {
			inRing = 1
			ipNum = len(ips)
		}
		if i == 0 {
			if inRing != 1 {
				if InPolygon(l[0], v) {
					inRing = 0
				} else {
					if InPolygon(v[0], l) {
						inRing = 20
					} else {
						inRing = 2
					}
				}
			} else {
				lInRing := 0
				lInRingOther := 0
				lOutRing := 0
				for _, m := range l {
					if InLineMatrix(m, v) {
						continue
					}
					if InPolygon(m, v) {
						lInRing++
					} else {
						lOutRing++
					}
				}
				for _, m := range v {
					if InLineMatrix(m, l) {
						continue
					}
					if InPolygon(m, l) {
						lInRingOther++
					}
				}
				if lInRing > 0 && lOutRing == 0 {
					if ipNum > 1 {
						inRing = 10
					} else {
						inRing = 11
					}
				} else if lInRing > 0 && lOutRing > 0 {
					inRing = 1
				} else if lInRing == 0 && lOutRing > 0 {
					if lInRingOther > 0 {
						if ipNum > 1 {
							inRing = 12
						} else {
							inRing = 13
						}
					} else {
						if ipNum > 1 {
							inRing = 22
						} else {
							inRing = 23
						}
					}
				}
			}
		} else {
			if inRing != 1 {
				if InPolygon(l[0], v) {
					if inRing != 2 {
						inRing = 2
						break
					}
				}
			} else {
				lInRing := 0
				lOutRing := 0
				for _, m := range l {
					if InLineMatrix(m, v) {
						continue
					}
					if InPolygon(m, v) {
						lInRing++
					} else {
						lOutRing++
					}
				}
				if lInRing > 0 && lOutRing == 0 {
					if inRing != 2 {
						inRing = 2
						break
					}
				} else if lInRing > 0 && lOutRing > 0 {
					inRing = 1
					break
				}
			}
		}

	}
	switch inRing {
	case 0:
		im.SetAtLeastString("2FF1FF212")
	case 1:
		im.SetAtLeastString("212101212")
	case 2:
		im.SetAtLeastString("FF2FF1212")
	case 10:
		im.SetAtLeastString("2FF11F212")
	case 12:
		im.SetAtLeastString("212F11FF2")
	case 13:
		im.SetAtLeastString("212F01FF2")
	case 11:
		im.SetAtLeastString("2FF01F212")
	case 20:
		im.SetAtLeastString("212FF1FF2")
	case 22:
		im.SetAtLeastString("FF2F11212")
	case 23:
		im.SetAtLeastString("FF2F01212")
	}
}
