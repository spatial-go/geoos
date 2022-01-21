package subdivision

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestDelaunayTriangulation_Subdivision(t *testing.T) {
	type fields struct {
		sites       []matrix.Matrix
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Subdivision", fields: fields{
				sites: []matrix.Matrix{{10, 10}, {20, 70}, {60, 30}, {80, 70}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DelaunayTriangulation{
				sites:       tt.fields.sites,
			}
			if got := d.Subdivision(); len(got.Edges()) != 15 || got.Edges()[3].Origin()[0] != -690 {
				t.Errorf("Subdivision() = %v", got)
			}
		})
	}
}
