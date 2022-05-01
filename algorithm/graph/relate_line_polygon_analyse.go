// package graph ...

package graph

func (r *RelationshipByDegrees) lineAndPolygonAnalyse(pointInPolygon, entityInPolygon int) {
	switch r.nLine {
	case 0:
		r.IM.Set(2, 2, 2)
		r.IM.Set(2, 1, 1)
		r.IM.Set(2, 0, 2)
		switch r.nPoint {
		case 1:
			switch pointInPolygon {
			case PartInPolygon:
				r.IM.Set(0, 0, 1)
				r.IM.Set(0, 1, -1)
				r.IM.Set(0, 2, -1)
				r.IM.Set(1, 0, 0)
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 2, -1)
			case PartOutPolygon:
				r.IM.Set(0, 0, -1)
				r.IM.Set(0, 1, -1)
				r.IM.Set(0, 2, 1)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 2, 0)
			case OnlyInPolygon:
				r.IM.Set(0, 0, 1)
				r.IM.Set(0, 1, 0)
				r.IM.Set(0, 2, -1)
				r.IM.Set(1, 0, 0)
				r.IM.Set(1, 1, -1)
				r.IM.Set(1, 2, -1)
			case OnlyOutPolygon:
				r.IM.Set(0, 0, -1)
				r.IM.Set(0, 1, 0)
				r.IM.Set(0, 2, 1)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, -1)
				r.IM.Set(1, 2, 0)
			}
		case 2:
			switch calcSumDegrees(r.degrees...) {
			case 8:
				r.IM.Set(0, 0, 1)
				r.IM.Set(0, 1, 0)
				r.IM.Set(0, 2, 1)
				switch pointInPolygon {
				case OnlyInPolygon:
					r.IM.Set(1, 0, 0)
					r.IM.Set(1, 1, -1)
					r.IM.Set(1, 2, -1)
				case OnlyOutPolygon:
					r.IM.Set(1, 0, -1)
					r.IM.Set(1, 1, -1)
					r.IM.Set(1, 2, 0)
				}
			case 7:
				r.IM.Set(0, 0, 1)
				r.IM.Set(0, 1, 0)
				r.IM.Set(0, 2, 1)
				r.IM.Set(1, 1, 0)
				switch pointInPolygon {
				case PartInPolygon:
					r.IM.Set(1, 0, 0)
					r.IM.Set(1, 2, -1)
				case PartOutPolygon:
					r.IM.Set(1, 0, -1)
					r.IM.Set(1, 2, 0)
				}
			case 6:

				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 2, -1)
				r.IM.Set(0, 1, -1)
				switch entityInPolygon {
				case OnlyInPolygon:
					r.IM.Set(0, 0, 1)
					r.IM.Set(0, 2, -1)
				case OnlyOutPolygon:
					r.IM.Set(0, 0, -1)
					r.IM.Set(0, 2, 1)
				}
			}
		default:
			r.IM.Set(0, 1, 0)
			switch pointInPolygon {
			case OnlyInPolygon:
				r.IM.Set(1, 1, -1)
				r.IM.Set(1, 0, 0)
				r.IM.Set(1, 2, -1)
			case OnlyOutPolygon:
				r.IM.Set(1, 1, -1)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 2, 0)
			case OnlyInLine:
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 2, -1)
			case BothPolygon:
				r.IM.Set(1, 1, -1)
				r.IM.Set(1, 0, 0)
				r.IM.Set(1, 2, 0)
			case PartInPolygon:
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 0, 0)
				r.IM.Set(1, 2, -1)
			case PartOutPolygon:
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 2, 0)
			}

			switch entityInPolygon {
			case OnlyInPolygon:
				r.IM.Set(0, 0, 1)
				r.IM.Set(0, 2, -1)
			case BothPolygon:
				r.IM.Set(0, 0, 1)
				r.IM.Set(0, 2, 1)
			}
		}
		switch pointInPolygon {
		case OnlyInPolygon:
			r.IM.Set(1, 0, 0)
			r.IM.Set(1, 2, -1)
		case OnlyOutPolygon:
			r.IM.Set(1, 0, -1)
			r.IM.Set(1, 2, 0)
		case BothPolygon:
			r.IM.Set(1, 0, 0)
			r.IM.Set(1, 2, 0)
		default:
			r.IM.Set(1, 0, 0)
		}
		switch entityInPolygon {
		case OnlyInPolygon:
			r.IM.Set(0, 0, 1)
			r.IM.Set(0, 1, -1)
			r.IM.Set(0, 2, -1)
			r.IM.Set(2, 0, 2)
		case OnlyOutPolygon:
			r.IM.Set(0, 0, -1)
			r.IM.Set(0, 1, -1)
			r.IM.Set(0, 2, 1)
		case BothPolygon:
			r.IM.Set(0, 0, 1)
			r.IM.Set(0, 1, 0)
			r.IM.Set(0, 2, 1)
		}
	default:
		r.IM.Set(2, 2, 2)
		r.IM.Set(2, 1, 1)
		r.IM.Set(2, 0, 2)
		r.IM.Set(0, 1, 1)
		switch r.nPoint {
		case 2:
			switch pointInPolygon {
			case OnlyInLine:
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 2, -1)
				r.IM.Set(0, 0, -1)
				r.IM.Set(0, 2, -1)
			case OnlyInPolygon:
				r.IM.Set(1, 0, 0)
				r.IM.Set(1, 1, -1)
				r.IM.Set(1, 2, -1)
				r.IM.Set(0, 0, 1)
				r.IM.Set(0, 2, -1)
			case OnlyOutPolygon:
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, -1)
				r.IM.Set(1, 2, 0)
				r.IM.Set(0, 0, -1)
				r.IM.Set(0, 2, 1)
			case PartInPolygon:
				r.IM.Set(1, 0, 0)
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 2, -1)
				r.IM.Set(0, 0, 1)
				r.IM.Set(0, 2, -1)
			case PartOutPolygon:
				r.IM.Set(1, 0, -1)
				r.IM.Set(1, 1, 0)
				r.IM.Set(1, 2, 1)
				r.IM.Set(0, 0, -1)
				r.IM.Set(0, 2, 1)
			}
		default:

			switch _, maxVal := calcMaxDegrees(r.degrees...); maxVal {
			case 4:
				r.IM.Set(0, 0, 1)
				r.IM.Set(0, 2, 1)
				switch pointInPolygon {
				case OnlyInLine:
					r.IM.Set(1, 0, -1)
					r.IM.Set(1, 1, 0)
					r.IM.Set(1, 2, -1)
				case OnlyInPolygon:
					r.IM.Set(1, 0, 0)
					r.IM.Set(1, 1, -1)
					r.IM.Set(1, 2, -1)
				case OnlyOutPolygon:
					r.IM.Set(1, 0, -1)
					r.IM.Set(1, 1, -1)
					r.IM.Set(1, 2, 0)
				case PartInPolygon:
					r.IM.Set(1, 0, 0)
					r.IM.Set(1, 1, 0)
					r.IM.Set(1, 2, -1)
				case PartOutPolygon:
					r.IM.Set(1, 0, -1)
					r.IM.Set(1, 1, 0)
					r.IM.Set(1, 2, 1)
				}
			default:
				switch pointInPolygon {
				case OnlyInLine:
					r.IM.Set(1, 0, -1)
					r.IM.Set(1, 1, 0)
					r.IM.Set(1, 2, -1)
					if entityInPolygon == OnlyInPolygon {
						r.IM.Set(0, 0, 1)
						r.IM.Set(0, 2, -1)
					} else {
						r.IM.Set(0, 0, -1)
						r.IM.Set(0, 2, 1)
					}
				case OnlyInPolygon:
					r.IM.Set(1, 0, 0)
					r.IM.Set(1, 1, -1)
					r.IM.Set(1, 2, -1)
					r.IM.Set(0, 0, 1)
					r.IM.Set(0, 2, -1)
				case OnlyOutPolygon:
					r.IM.Set(1, 0, -1)
					r.IM.Set(1, 1, -1)
					r.IM.Set(1, 2, 0)
					r.IM.Set(0, 0, -1)
					r.IM.Set(0, 2, 1)
				case PartInPolygon:
					r.IM.Set(1, 0, 0)
					r.IM.Set(1, 1, 0)
					r.IM.Set(1, 2, -1)
					r.IM.Set(0, 0, 1)
					r.IM.Set(0, 2, -1)
				case PartOutPolygon:
					r.IM.Set(1, 0, -1)
					r.IM.Set(1, 1, 0)
					r.IM.Set(1, 2, 1)
					r.IM.Set(0, 0, -1)
					r.IM.Set(0, 2, 1)
				}
			}
		}
	}
}
