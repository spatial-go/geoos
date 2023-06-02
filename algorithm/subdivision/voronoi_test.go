package subdivision

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/space"
)

func TestVoronoi_GetResult(t *testing.T) {
	type fields struct {
		sites []matrix.Matrix
		env   *envelope.Envelope
	}
	tests := []struct {
		name   string
		fields fields
		want   []matrix.PolygonMatrix
	}{
		{
			name: "voronoi test",
			fields: fields{
				sites: []matrix.Matrix{{4, 3}, {15, 0}, {0, 4}, {15, 11}, {10, 3}},
				env:   envelope.Bound([]matrix.Matrix{{-10, -6}, {10, 6}}),
			},
			want: []matrix.PolygonMatrix{
				{{{10, -6},
					{10, 6},
					{7.000000000000001, 6},
					{7.000000000000001, -6},
					{10, -6},
				},
				},
				{{{10, -2.666666666666667},
					{10, -6},
					{8.000000000000002, -6.000000000000001},
					{10, -2.666666666666667},
				},
				},
				{{{8.000000000000002, -6.000000000000001},
					{10, -2.666666666666667},
					{10, 6},
					{7.000000000000001, 6},
					{7.000000000000001, -6},
					{8.000000000000002, -6.000000000000001},
				},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewVoronoi()
			v.AddSites(tt.fields.sites)
			if tt.fields.env != nil {
				v.SetEnvelope(*tt.fields.env)
			}
			got := v.GetResult()
			gotEnv := envelope.PolygonMatrixList(got)
			wantEnv := envelope.PolygonMatrixList(tt.want)
			if !wantEnv.Proximity(gotEnv) {
				t.Errorf("Get Voronoi Result Error got=%v ,want=%v", gotEnv, wantEnv)
			}
			gotCollection := space.Collection{}
			for _, pm := range got {
				if len(pm) == 0 {
					continue
				}
				gotCollection = append(gotCollection, space.Polygon{pm[0]})
			}
			// t.Log(wkt.MarshalString(gotCollection))
		})
	}
}
