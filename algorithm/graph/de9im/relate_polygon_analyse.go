package de9im

func (r *RelationshipByDegrees) polygonTwoAnalyse(pointInPolygon, entityInPolygon int) {
	switch r.nLine {
	case 0:
		switch r.nPoint {
		case 1:
			switch entityInPolygon {
			case OnlyInPolygon:
				r.IM.Set(0, 0, 2)
				r.IM.Set(0, 1, -1)
				r.IM.Set(0, 2, -1)
				r.IM.Set(1, 0, 1)
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 2, -1)
				r.IM.Set(2, 2, 2)
				r.IM.Set(2, 1, 1)
				r.IM.Set(2, 0, 2)
			case OnlyOutPolygon:
				r.IM.Set(0, 0, -1)
				r.IM.Set(0, 1, -1)
				r.IM.Set(0, 2, 2)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 2, 1)
				r.IM.Set(2, 2, 2)
				r.IM.Set(2, 1, 1)
				r.IM.Set(2, 0, 2)
			case IncludePolygon:
				r.IM.Set(0, 0, 2)
				r.IM.Set(0, 1, 1)
				r.IM.Set(0, 2, 2)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 2, 1)
				r.IM.Set(2, 2, 2)
				r.IM.Set(2, 1, -1)
				r.IM.Set(2, 0, -1)
			}
		default:
			r.IM.Set(0, 0, 2)
			r.IM.Set(0, 1, 1)
			r.IM.Set(0, 2, 2)
			r.IM.Set(1, 0, 1)
			r.IM.Set(1, 1, 0)
			r.IM.Set(1, 2, 1)
			r.IM.Set(2, 2, 2)
			r.IM.Set(2, 1, 1)
			r.IM.Set(2, 0, 2)
		}
	default:
		switch r.nPoint {
		case 2:
			switch entityInPolygon {
			case OnlyInPolygon, PartInPolygon:
				r.IM.Set(0, 0, 2)
				r.IM.Set(0, 1, -1)
				r.IM.Set(0, 2, -1)
				r.IM.Set(1, 0, 1)
				r.IM.Set(1, 1, 1)
				r.IM.Set(1, 2, -1)
				r.IM.Set(2, 2, 2)
				r.IM.Set(2, 1, 1)
				r.IM.Set(2, 0, 2)
			case OnlyOutPolygon, PartOutPolygon:
				r.IM.Set(0, 0, -1)
				r.IM.Set(0, 1, -1)
				r.IM.Set(0, 2, 2)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, 1)
				r.IM.Set(1, 2, 1)
				r.IM.Set(2, 2, 2)
				r.IM.Set(2, 1, 1)
				r.IM.Set(2, 0, 2)
			case IncludePolygon:
				r.IM.Set(0, 0, 2)
				r.IM.Set(0, 1, 1)
				r.IM.Set(0, 2, 2)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, 1)
				r.IM.Set(1, 2, 1)
				r.IM.Set(2, 2, 2)
				r.IM.Set(2, 1, -1)
				r.IM.Set(2, 0, -1)
			}
		default:
			switch entityInPolygon {
			case OnlyInPolygon, PartInPolygon:
				r.IM.Set(0, 0, 2)
				r.IM.Set(0, 1, -1)
				r.IM.Set(0, 2, -1)
				r.IM.Set(1, 0, 1)
				r.IM.Set(1, 1, 1)
				r.IM.Set(1, 2, -1)
				r.IM.Set(2, 2, 2)
				r.IM.Set(2, 1, 1)
				r.IM.Set(2, 0, 2)
			case OnlyOutPolygon, PartOutPolygon:
				r.IM.Set(0, 0, -1)
				r.IM.Set(0, 1, -1)
				r.IM.Set(0, 2, 2)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, 1)
				r.IM.Set(1, 2, 1)
				r.IM.Set(2, 2, 2)
				r.IM.Set(2, 1, 1)
				r.IM.Set(2, 0, 2)
			case IncludePolygon:
				r.IM.Set(0, 0, 2)
				r.IM.Set(0, 1, 1)
				r.IM.Set(0, 2, 2)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, 1)
				r.IM.Set(1, 2, 1)
				r.IM.Set(2, 2, 2)
				r.IM.Set(2, 1, -1)
				r.IM.Set(2, 0, -1)
			default:
				r.IM.Set(0, 0, 2)
				r.IM.Set(0, 1, 1)
				r.IM.Set(0, 2, 2)
				r.IM.Set(1, 0, 1)
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 2, 1)
				r.IM.Set(2, 2, 2)
				r.IM.Set(2, 1, 1)
				r.IM.Set(2, 0, 2)
			}
		}
	}
}
