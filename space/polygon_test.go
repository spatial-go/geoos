package space

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/filter"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestPolygon_Buffer(t *testing.T) {
	type args struct {
		width    float64
		quadsegs int
	}
	tests := []struct {
		name string
		p    Polygon
		args args
		want Geometry
	}{
		{name: "polygon buffer", p: Polygon{{{122.993197, 41.117725}, {122.999399, 41.115696}, {122.99573, 41.109516},
			{122.987146, 41.106994}, {122.984775, 41.107699}, {122.990687, 41.117878}, {122.993197, 41.117725}}},
			args: args{
				width:    0.001,
				quadsegs: 4,
			},
			want: Polygon{{{122.99325784324384, 41.118723147333654}, {122.99350793587271, 41.11867543089336}, {122.99970993587272, 41.11664643089336},
				{123.00002331519273, 41.11647717254184}, {123.0002574970166, 41.11620881855708}, {123.00038277421193, 41.11587541098052}, {123.00038325474164, 41.11551924422623},
				{123.00025887764812, 41.11518549982346}, {122.99658987764812, 41.10900549982346}, {122.99634342934228, 41.10872625039283}, {122.99601188795101, 41.1085565526679},
				{122.98742788795101, 41.1060345526679}, {122.98686098957647, 41.106035475582736}, {122.98448998957647, 41.10674047558273}, {122.98416932802418, 41.10690328556774},
				{122.98392699858715, 41.10716900603418}, {122.98379434195499, 41.10750327110905}, {122.98378851473228, 41.10786284998192}, {122.98391027055912, 41.108201237985504},
				{122.98982227055912, 41.118380237985505}, {122.99006123336798, 41.11865801033471}, {122.99038421204416, 41.11883105794881}, {122.99074784324384, 41.11887614733365},
				{122.99325784324384, 41.118723147333654}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Buffer(tt.args.width, tt.args.quadsegs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Polygon.Buffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolygon_Filter(t *testing.T) {
	var f filter.Filter[matrix.Matrix] = matrix.CreateFilterMatrix()
	tests := []struct {
		name string
		p    Polygon
		want Polygon
	}{
		{"polygon filter", Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}}, {{1.5, 1.5}, {1.5, 2}, {2, 2}, {2, 2}, {2, 1.5}, {1.5, 1.5}}},
			Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}, {{1.5, 1.5}, {1.5, 2}, {2, 2}, {2, 1.5}, {1.5, 1.5}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Filter(f)
			if !got.Equals(tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
