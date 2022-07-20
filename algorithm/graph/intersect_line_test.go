package graph

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay/chain"
)

func TestIntersectLine(t *testing.T) {
	type args struct {
		m  matrix.LineMatrix
		m1 matrix.LineMatrix
	}
	tests := []struct {
		name string
		args args
		want chain.CorrelationNodeResult
	}{
		{"line line0", args{matrix.LineMatrix{{100, 100}, {100, 101}}, matrix.LineMatrix{{100, 100}, {100, 101}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{
					InterNode:       matrix.Matrix{100, 100},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {100, 101}},
				},
				{
					InterNode:       matrix.Matrix{100, 101},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {100, 101}},
				},
			}, []*chain.IntersectionCorrelationNode{
				{
					InterNode:       matrix.Matrix{100, 100},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {100, 101}},
				},
				{
					InterNode:       matrix.Matrix{100, 101},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {100, 101}},
				},
			},
			},
		},
		{"line line1", args{matrix.LineMatrix{{100, 100}, {100, 101}}, matrix.LineMatrix{{100, 100}, {90, 102}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{
					InterNode:       matrix.Matrix{100, 100},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {100, 101}},
				},
			}, []*chain.IntersectionCorrelationNode{
				{
					InterNode:       matrix.Matrix{100, 100},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {90, 102}},
				},
			},
			},
		},
		{"line line2", args{matrix.LineMatrix{{90, 90}, {100, 100}, {100, 101}, {102, 102}}, matrix.LineMatrix{{95, 95}, {100, 100}, {100, 101}, {90, 102}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{
					InterNode:       matrix.Matrix{95, 95},
					CorrelationNode: matrix.LineMatrix{{90, 90}, {95, 95}},
				},
				{
					InterNode:       matrix.Matrix{95, 95},
					CorrelationNode: matrix.LineMatrix{{95, 95}, {100, 100}, {100, 101}},
				},
				{
					InterNode:       matrix.Matrix{100, 101},
					CorrelationNode: matrix.LineMatrix{{95, 95}, {100, 100}, {100, 101}},
				},
				{
					InterNode:       matrix.Matrix{100, 101},
					CorrelationNode: matrix.LineMatrix{{100, 101}, {102, 102}},
				},
			}, []*chain.IntersectionCorrelationNode{
				{
					InterNode:       matrix.Matrix{95, 95},
					CorrelationNode: matrix.LineMatrix{{95, 95}, {100, 100}, {100, 101}},
				},
				{
					InterNode:       matrix.Matrix{100, 101},
					CorrelationNode: matrix.LineMatrix{{95, 95}, {100, 100}, {100, 101}},
				},
				{
					InterNode:       matrix.Matrix{100, 101},
					CorrelationNode: matrix.LineMatrix{{100, 101}, {90, 102}},
				},
			},
			},
		},
		{"line line21", args{matrix.LineMatrix{{90, 90}, {100, 100}, {100, 101}, {102, 102}},
			matrix.LineMatrix{{95, 98}, {100, 100}, {100, 101}, {90, 102}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{
					InterNode:       matrix.Matrix{100, 100},
					CorrelationNode: matrix.LineMatrix{{90, 90}, {100, 100}},
				},
				{
					InterNode:       matrix.Matrix{100, 100},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {100, 101}},
				},
				{
					InterNode:       matrix.Matrix{100, 101},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {100, 101}},
				},
				{
					InterNode:       matrix.Matrix{100, 101},
					CorrelationNode: matrix.LineMatrix{{100, 101}, {102, 102}},
				},
			}, []*chain.IntersectionCorrelationNode{
				{
					InterNode:       matrix.Matrix{100, 100},
					CorrelationNode: matrix.LineMatrix{{95, 98}, {100, 100}},
				},
				{
					InterNode:       matrix.Matrix{100, 100},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {100, 101}},
				},
				{
					InterNode:       matrix.Matrix{100, 101},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {100, 101}},
				},
				{
					InterNode:       matrix.Matrix{100, 101},
					CorrelationNode: matrix.LineMatrix{{100, 101}, {90, 102}},
				},
			},
			},
		},
		{"line line3", args{matrix.LineMatrix{{100, 100}, {200, 100}, {200, 0}}, matrix.LineMatrix{{100, 150}, {250, 0}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{
					InterNode:       matrix.Matrix{150, 100},
					CorrelationNode: matrix.LineMatrix{{100, 100}, {150, 100}},
				},
				{
					InterNode:       matrix.Matrix{150, 100},
					CorrelationNode: matrix.LineMatrix{{150, 100}, {200, 100}, {200, 50}},
				},
				{
					InterNode:       matrix.Matrix{200, 50},
					CorrelationNode: matrix.LineMatrix{{150, 100}, {200, 100}, {200, 50}},
				},
				{
					InterNode:       matrix.Matrix{200, 50},
					CorrelationNode: matrix.LineMatrix{{200, 50}, {200, 0}},
				},
			}, []*chain.IntersectionCorrelationNode{
				{
					InterNode:       matrix.Matrix{150, 100},
					CorrelationNode: matrix.LineMatrix{{100, 150}, {150, 100}},
				},
				{
					InterNode:       matrix.Matrix{150, 100},
					CorrelationNode: matrix.LineMatrix{{150, 100}, {200, 50}},
				},
				{
					InterNode:       matrix.Matrix{200, 50},
					CorrelationNode: matrix.LineMatrix{{150, 100}, {200, 50}},
				},
				{
					InterNode:       matrix.Matrix{200, 50},
					CorrelationNode: matrix.LineMatrix{{200, 50}, {250, 0}},
				},
			},
			},
		},
		{"line line4", args{matrix.LineMatrix{{100, 100}, {200, 100}, {200, 0}}, matrix.LineMatrix{{150, 150}, {150, 50}, {200, 50}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{150, 100}, CorrelationNode: matrix.LineMatrix{{100, 100}, {150, 100}}},
				{InterNode: matrix.Matrix{150, 100}, CorrelationNode: matrix.LineMatrix{{150, 100}, {200, 100}, {200, 50}}},
				{InterNode: matrix.Matrix{200, 50}, CorrelationNode: matrix.LineMatrix{{150, 100}, {200, 100}, {200, 50}}},
				{InterNode: matrix.Matrix{200, 50}, CorrelationNode: matrix.LineMatrix{{200, 50}, {200, 0}}},
			}, []*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{150, 100}, CorrelationNode: matrix.LineMatrix{{150, 150}, {150, 100}}},
				{InterNode: matrix.Matrix{150, 100}, CorrelationNode: matrix.LineMatrix{{150, 100}, {150, 50}, {200, 50}}},
				{InterNode: matrix.Matrix{200, 50}, CorrelationNode: matrix.LineMatrix{{150, 100}, {150, 50}, {200, 50}}},
			},
			},
		},
		{"line line5", args{matrix.LineMatrix{{100, 100}, {200, 100}, {200, 0}}, matrix.LineMatrix{{100, 200}, {300, 0}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{100, 100}, {200, 100}}},
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{200, 100}, {200, 0}}},
			}, []*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{100, 200}, {200, 100}}},
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{200, 100}, {300, 0}}},
			},
			},
		},
		{"line line6", args{matrix.LineMatrix{{100, 100}, {200, 100}, {200, 0}}, matrix.LineMatrix{{100, 200}, {200, 100}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{100, 100}, {200, 100}}},
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{200, 100}, {200, 0}}},
			}, []*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{100, 200}, {200, 100}}},
			},
			},
		},
		{"line line7", args{matrix.LineMatrix{{100, 100}, {200, 100}, {200, 0}}, matrix.LineMatrix{{100, 200}, {200, 100}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{100, 100}, {200, 100}}},
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{200, 100}, {200, 0}}},
			}, []*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{100, 200}, {200, 100}}},
			},
			},
		},
		{"line poly1", args{matrix.LineMatrix{{100, 100}, {200, 100}, {200, 200}, {100, 200}, {100, 100}}, matrix.LineMatrix{{100, 200}, {300, 0}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{100, 100}, {200, 100}}},
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{200, 100}, {200, 200}, {100, 200}}},
				{InterNode: matrix.Matrix{100, 200}, CorrelationNode: matrix.LineMatrix{{200, 100}, {200, 200}, {100, 200}}},
				{InterNode: matrix.Matrix{100, 200}, CorrelationNode: matrix.LineMatrix{{100, 200}, {100, 100}}},
			}, []*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{100, 200}, CorrelationNode: matrix.LineMatrix{{100, 200}, {200, 100}}},
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{100, 200}, {200, 100}}},
				{InterNode: matrix.Matrix{200, 100}, CorrelationNode: matrix.LineMatrix{{200, 100}, {300, 0}}},
			},
			},
		},
		{"line poly2", args{matrix.LineMatrix{{100, 100}, {200, 100}, {200, 200}, {100, 200}, {100, 100}}, matrix.LineMatrix{{100, 250}, {350, 0}}},
			chain.CorrelationNodeResult{[]*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{200, 150}, CorrelationNode: matrix.LineMatrix{{100, 100}, {200, 100}, {200, 150}}},
				{InterNode: matrix.Matrix{200, 150}, CorrelationNode: matrix.LineMatrix{{200, 150}, {200, 200}, {150, 200}}},
				{InterNode: matrix.Matrix{150, 200}, CorrelationNode: matrix.LineMatrix{{200, 150}, {200, 200}, {150, 200}}},
				{InterNode: matrix.Matrix{150, 200}, CorrelationNode: matrix.LineMatrix{{150, 200}, {100, 200}, {100, 100}}},
			}, []*chain.IntersectionCorrelationNode{
				{InterNode: matrix.Matrix{150, 200}, CorrelationNode: matrix.LineMatrix{{100, 250}, {150, 200}}},
				{InterNode: matrix.Matrix{150, 200}, CorrelationNode: matrix.LineMatrix{{150, 200}, {200, 150}}},
				{InterNode: matrix.Matrix{200, 150}, CorrelationNode: matrix.LineMatrix{{150, 200}, {200, 150}}},
				{InterNode: matrix.Matrix{200, 150}, CorrelationNode: matrix.LineMatrix{{200, 150}, {350, 0}}},
			},
			},
		},
	}
	for _, tt := range tests {
		if !geoos.GeoosTestTag && tt.name != "line line21" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectLine(tt.args.m, tt.args.m1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v IntersectLine() = %v ：%T,\n want %v ：%T", tt.name, got, got, tt.want, tt.want)
			}
		})
	}
}
