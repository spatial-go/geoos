package de9im

type LineRelationshipByDegrees struct {
	*RelationshipByDegrees
	pointInPolygon, entityInPolygon int
}

func (l *LineRelationshipByDegrees) setInteriorIM() {
	isInterior, isBoundary, isExterior := l.relatePolygon(l.entityInPolygon)
	if isInterior == 1 {
		l.IM.Set(0, 0, 1)
	} else {
		l.IM.Set(0, 0, -1)
	}
	if isBoundary == 1 {
		l.IM.Set(0, 1, 1)
	} else {
		l.IM.Set(0, 1, -1)
	}
	if isExterior == 1 {
		l.IM.Set(0, 2, 1)
	} else {
		l.IM.Set(0, 2, -1)
	}
}

func (l *LineRelationshipByDegrees) setExteriorIM() {
	l.IM.Set(2, 2, 2)
	l.IM.Set(2, 1, 1)
	l.IM.Set(2, 0, 2)
}

func (l *LineRelationshipByDegrees) setBoundaryIM() {
	isInterior, isBoundary, isExterior := l.relatePolygon(l.pointInPolygon)
	if isInterior == 1 {
		l.IM.Set(1, 0, 0)
	} else {
		l.IM.Set(1, 0, -1)
	}
	if isBoundary == 1 {
		l.IM.Set(1, 1, 0)
	} else {
		l.IM.Set(1, 1, -1)
	}
	if isExterior == 1 {
		l.IM.Set(1, 2, 0)
	} else {
		l.IM.Set(0, 2, -1)
	}
}

func (l *LineRelationshipByDegrees) produce() {
	l.setExteriorIM()
	l.setBoundaryIM()
	l.setInteriorIM()
}

func (l *LineRelationshipByDegrees) relatePolygon(inPolygon int) (
	isInterior, isBoundary, isExterior int) {
	if inPolygon == OnlyInPolygon || inPolygon == BothPolygon || inPolygon == PartInPolygon || inPolygon == IncludePolygon {
		isInterior = 1
	}

	if inPolygon == OnlyOutPolygon || inPolygon == BothPolygon || inPolygon == PartOutPolygon {
		isExterior = 1
	}

	if inPolygon == OnlyInLine || inPolygon == PartInPolygon || inPolygon == PartOutPolygon {
		isBoundary = 1
	}
	return
}
