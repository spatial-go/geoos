package de9im

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

func (r *RelationshipByStructure) lineAnalyse(pointInPolygon, entityInPolygon int) {
	switch r.Arg[1].(type) {
	case matrix.LineMatrix:
		switch {
		case r.nPoint == 2 && r.nLine == 1 && r.maxDlLine == 6:
			r.relationshipSymbol = RLL3
		case r.nPoint == 2 && r.nLine == 1 && r.maxDlLine == 5:
			r.relationshipSymbol = RLL8
		case r.nPoint == 2 && r.nLine == 1 && r.maxDlLine == 4:
			r.relationshipSymbol = RLL10
		case r.nPoint == 2 && r.nLine == 1 && r.maxDlLine == 1:
			r.relationshipSymbol = RLL25
		case r.nPoint == 2 && r.nLine == 1 && r.maxDlLine == 3:
			r.relationshipSymbol = RLL26
		case r.nPoint == 2 && r.nLine == 1 && r.maxDlLine == 2:
			r.relationshipSymbol = RLL30

		// TODO two case is the same
		case r.nPoint == 3 && r.nLine == 1 && r.maxDlLine == 5 && r.maxDlPoint == 3:
			r.relationshipSymbol = RLL9
		case r.nPoint == 3 && r.nLine == 1 && r.maxDlLine == 5 && r.maxDlPoint == 3:
			r.relationshipSymbol = RLL18
		case r.nPoint == 3 && r.nLine == 1 && r.maxDlLine == 3 && r.maxDlPoint == 2:
			r.relationshipSymbol = RLL24
		case r.nPoint == 3 && r.nLine == 1 && r.maxDlLine == 3 && r.maxDlPoint == 3:
			r.relationshipSymbol = RLL29

		case r.nPoint == 4 && r.nLine == 1 && r.maxDlLine == 5:
			r.relationshipSymbol = RLL19
		case r.nPoint == 4 && r.nLine == 1 && r.maxDlLine == 3:
			r.relationshipSymbol = RLL33

		case r.nPoint >= 5 && r.nLine == 1 && r.maxDlLine == 5:
			r.relationshipSymbol = RLL17

		case r.nPoint == 1 && r.nLine == 0 && r.maxDlPoint == 4:
			r.relationshipSymbol = RLL2
		case r.nPoint == 1 && r.nLine == 0 && r.maxDlPoint == 3:
			r.relationshipSymbol = RLL4
		case r.nPoint == 1 && r.nLine == 0 && r.maxDlPoint == 2:
			r.relationshipSymbol = RLL21

		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 3 && r.sumDlPoint >= 6:
			if r.nCompositeLine == 1 {
				r.relationshipSymbol = RLL5
			}
			if r.nCompositeLine == 2 {
				r.relationshipSymbol = RLL12
			}
		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 3 && r.sumDlPoint == 5:
			r.relationshipSymbol = RLL27
		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 4:
			r.relationshipSymbol = RLL6
		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 2:
			r.relationshipSymbol = RLL20
		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 4:
			r.relationshipSymbol = RLL23

		case r.nPoint == 3 && r.nLine == 0 && r.maxDlPoint == 4 && r.sumDlPoint == 8:
			r.relationshipSymbol = RLL22
		case r.nPoint == 3 && r.nLine == 0 && r.maxDlPoint == 4 && r.sumDlPoint == 9:
			r.relationshipSymbol = RLL28
		case r.nPoint == 3 && r.nLine == 0 && r.maxDlPoint == 4:
			if true {
				r.relationshipSymbol = RLL7
			} else {
				r.relationshipSymbol = RLL15
			}
		case r.nPoint == 3 && r.nLine == 0 && r.maxDlPoint == 3 && r.sumDlPoint >= 9:
			r.relationshipSymbol = RLL13
		case r.nPoint == 3 && r.nLine == 0 && r.maxDlPoint == 3 && r.sumDlPoint == 8:
			r.relationshipSymbol = RLL31

		case r.nPoint == 4 && r.nLine == 0 && r.maxDlPoint == 3:
			r.relationshipSymbol = RLL11
		case r.nPoint == 4 && r.nLine == 0 && r.maxDlPoint == 4 && r.sumDlPoint == 13:
			r.relationshipSymbol = RLL16
		case r.nPoint == 4 && r.nLine == 0 && r.maxDlPoint == 4 && r.sumDlPoint == 12:
			r.relationshipSymbol = RLL32
		case r.nPoint >= 5 && r.nLine == 0 && r.maxDlPoint == 4:
			r.relationshipSymbol = RLL14
		}
	case matrix.PolygonMatrix:
		switch {
		case r.nPoint == 2 && r.nLine >= 1 && r.maxDlLine == 6 && pointInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RLA4
		case r.nPoint == 2 && r.nLine >= 1 && r.maxDlLine == 6 && pointInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RLA6
		case r.nPoint == 2 && r.nLine >= 1 && r.maxDlLine == 6 && pointInPolygon == BothPolygon:
			r.relationshipSymbol = RLA12
		case r.nPoint == 2 && r.nLine >= 1 && r.maxDlLine == 4:
			r.relationshipSymbol = RLA24
		case r.nPoint == 2 && r.nLine >= 1 && r.maxDlLine == 5 && entityInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RLA26
		case r.nPoint == 2 && r.nLine >= 1 && r.maxDlLine == 5 && pointInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RLA30

		case r.nPoint == 3 && r.nLine >= 1 && r.maxDlLine == 6 && r.maxDlPoint == 4 && pointInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RLA10
		case r.nPoint == 3 && r.nLine >= 1 && r.maxDlLine == 6 && r.maxDlPoint == 4 && pointInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RLA11
		case r.nPoint == 3 && r.nLine >= 1 && r.maxDlLine == 6 && r.maxDlPoint == 3:
			r.relationshipSymbol = RLA18
		case r.nPoint == 3 && r.nLine >= 1 && r.maxDlLine == 5 && r.maxDlPoint == 3:
			r.relationshipSymbol = RLA25
		case r.nPoint == 3 && r.nLine >= 1 && r.maxDlLine == 5 && r.maxDlPoint == 3:
			r.relationshipSymbol = RLA27
		case r.nPoint == 3 && r.nLine >= 1 && r.maxDlLine == 5 && r.maxDlPoint == 4 && pointInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RLA29
		case r.nPoint == 3 && r.nLine >= 1 && r.maxDlLine == 5 && r.maxDlPoint == 4 && pointInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RLA31

		case r.nPoint >= 4 && r.nLine >= 1 && r.maxDlLine == 6:
			r.relationshipSymbol = RLA17
		case r.nPoint >= 4 && r.nLine >= 1 && r.maxDlLine == 5:
			r.relationshipSymbol = RLA28

		case r.nPoint == 1 && r.nLine == 0 && r.maxDlPoint == 3 && pointInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RLA3
		case r.nPoint == 1 && r.nLine == 0 && r.maxDlPoint == 3 && pointInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RLA5
		case r.nPoint == 1 && r.nLine == 0 && r.maxDlPoint == 3 && pointInPolygon == BothPolygon:
			r.relationshipSymbol = RLA8
		case r.nPoint == 1 && r.nLine == 0 && r.maxDlPoint == 2 && entityInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RLA14
		case r.nPoint == 1 && r.nLine == 0 && r.maxDlPoint == 2 && entityInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RLA15

			//TODO

		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 4 && pointInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RLA7
		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 4 && pointInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RLA9
		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 4 && entityInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RLA22
		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 4 && entityInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RLA23
		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 3 && entityInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RLA13
		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 3 && entityInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RLA16

		case r.nPoint >= 3 && r.nLine == 0 && r.maxDlPoint == 4 && entityInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RLA19
		case r.nPoint >= 3 && r.nLine == 0 && r.maxDlPoint == 4 && entityInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RLA20

		}

	}
}

func (r *RelationshipByStructure) polygonAnalyse(pointInPolygon, entityInPolygon int) {
	switch r.Arg[1].(type) {
	case matrix.PolygonMatrix:
		switch {
		case r.nPoint == 0 && r.nLine >= 1:
			r.relationshipSymbol = RAA9
		case r.nPoint >= 2 && r.nLine >= 1 && entityInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RAA3
		case r.nPoint >= 2 && r.nLine >= 1 && entityInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RAA10
		case r.nPoint >= 2 && r.nLine >= 1:
			r.relationshipSymbol = RAA11

		case r.nPoint == 1 && r.nLine == 0 && entityInPolygon == OnlyOutPolygon:
			r.relationshipSymbol = RAA2
		case r.nPoint == 1 && r.nLine == 0 && entityInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RAA8
		case r.nPoint == 1 && r.nLine == 0:
			r.relationshipSymbol = RAA7

		case r.nPoint == 2 && r.nLine == 0 && r.maxDlPoint == 4:
			r.relationshipSymbol = RAA4

		case r.nPoint == 0 && r.nLine == 0 && entityInPolygon == OnlyInPolygon:
			r.relationshipSymbol = RAA6
		case r.nPoint == 0 && r.nLine == 0:
			r.relationshipSymbol = RAA5

		}

	}
}
