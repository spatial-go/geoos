package overlay

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestPointOverlay_Intersection(t *testing.T) {
	type fields struct {
		subject  matrix.Steric
		clipping matrix.Steric
	}
	tests := []struct {
		name    string
		fields  fields
		want    matrix.Steric
		wantErr bool
	}{
		{"point point0", fields{matrix.Matrix{100, 100}, matrix.Matrix{100, 100}}, matrix.Matrix{100, 100}, false},
		{"point point1", fields{matrix.Matrix{100, 100}, matrix.Matrix{100, 101}}, nil, false},
		{"point line0", fields{matrix.Matrix{100, 100}, matrix.LineMatrix{{100, 100}, {100, 101}}}, matrix.Matrix{100, 100}, false},
		{"point line1", fields{matrix.Matrix{100, 100}, matrix.LineMatrix{{100, 105}, {100, 101}}}, nil, false},
		{"point poly1", fields{matrix.Matrix{100, 100}, matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}}, matrix.Matrix{100, 100}, false},
		{"point poly2", fields{matrix.Matrix{100, 100}, matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}}}, matrix.Matrix{100, 100}, false},
		{"point poly3", fields{matrix.Matrix{100, 100}, matrix.PolygonMatrix{{{105, 105}, {105, 101}, {101, 101}, {101, 105}, {105, 105}}}}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PointOverlay{
				Subject:  tt.fields.subject,
				Clipping: tt.fields.clipping,
			}
			got, err := p.Intersection()
			if (err != nil) != tt.wantErr {
				t.Errorf("PointOverlay.Intersection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointOverlay.Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
