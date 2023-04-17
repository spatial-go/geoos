package de9im

type PolygonRelationshipByDegrees struct {
	*RelationshipByDegrees
	entityInPolygon int
}

func (l *PolygonRelationshipByDegrees) setInteriorIM() {
	isInterior, isBoundary, isExterior := l.relatePolygon(l.entityInPolygon)
	l.IM.Set(0, 0, isInterior)
	l.IM.Set(0, 1, isBoundary)
	l.IM.Set(0, 2, isExterior)
}

func (l *PolygonRelationshipByDegrees) setExteriorIM() {
	l.IM.Set(2, 2, 2)
	l.IM.Set(2, 1, 1)
	l.IM.Set(2, 0, 2)
	if l.entityInPolygon == IncludePolygon {
		l.IM.Set(2, 1, -1)
		l.IM.Set(2, 0, -1)
	}

}

func (l *PolygonRelationshipByDegrees) setBoundaryIM() {
	if l.entityInPolygon == OnlyInPolygon || l.entityInPolygon == BothPolygon || l.entityInPolygon == PartInPolygon {
		l.IM.Set(1, 0, 1)
	} else {
		l.IM.Set(1, 0, -1)
	}
	if l.nLine == 0 {
		l.IM.Set(1, 1, 0)
	} else {
		l.IM.Set(1, 1, 1)
	}
	if l.entityInPolygon == OnlyInPolygon || l.entityInPolygon == PartInPolygon {
		l.IM.Set(1, 2, -1)
	} else {
		l.IM.Set(1, 2, 1)
	}
}

func (l *PolygonRelationshipByDegrees) produce() {
	l.setExteriorIM()
	l.setBoundaryIM()
	l.setInteriorIM()
}

func (l *PolygonRelationshipByDegrees) relatePolygon(inPolygon int) (
	isInterior, isBoundary, isExterior int) {
	if inPolygon == OnlyInPolygon || inPolygon == BothPolygon || inPolygon == PartInPolygon || inPolygon == IncludePolygon {
		isInterior = 2
	} else {
		isInterior = -1
	}

	if inPolygon == OnlyOutPolygon || inPolygon == BothPolygon || inPolygon == PartOutPolygon || inPolygon == IncludePolygon {
		isExterior = 2
	} else {
		isExterior = -1
	}

	if inPolygon == OnlyInPolygon || inPolygon == OnlyOutPolygon || inPolygon == PartInPolygon || inPolygon == PartOutPolygon {
		isBoundary = -1
	} else {
		isBoundary = 1
	}
	return
}
