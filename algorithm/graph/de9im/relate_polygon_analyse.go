package de9im

import "github.com/spatial-go/geoos/algorithm/operation"

// PolygonRelationshipByDegrees ...
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
	if l.entityInPolygon == operation.IncludePolygon {
		l.IM.Set(2, 1, -1)
		l.IM.Set(2, 0, -1)
	}

}

func (l *PolygonRelationshipByDegrees) setBoundaryIM() {
	if l.entityInPolygon == operation.OnlyInPolygon || l.entityInPolygon == operation.BothPolygon || l.entityInPolygon == operation.PartInPolygon {
		l.IM.Set(1, 0, 1)
	} else {
		l.IM.Set(1, 0, -1)
	}
	if l.nLine == 0 {
		l.IM.Set(1, 1, 0)
	} else {
		l.IM.Set(1, 1, 1)
	}
	if l.entityInPolygon == operation.OnlyInPolygon || l.entityInPolygon == operation.PartInPolygon {
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
	if inPolygon == operation.OnlyInPolygon || inPolygon == operation.BothPolygon || inPolygon == operation.PartInPolygon || inPolygon == operation.IncludePolygon {
		isInterior = 2
	} else {
		isInterior = -1
	}

	if inPolygon == operation.OnlyOutPolygon || inPolygon == operation.BothPolygon || inPolygon == operation.PartOutPolygon || inPolygon == operation.IncludePolygon {
		isExterior = 2
	} else {
		isExterior = -1
	}

	if inPolygon == operation.OnlyInPolygon || inPolygon == operation.OnlyOutPolygon || inPolygon == operation.PartInPolygon || inPolygon == operation.PartOutPolygon {
		isBoundary = -1
	} else {
		isBoundary = 1
	}
	return
}
