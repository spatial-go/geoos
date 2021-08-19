package overlay

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestLineOverlay_Intersection(t *testing.T) {
	type fields struct {
		PointOverlay *PointOverlay
	}
	tests := []struct {
		name    string
		fields  fields
		want    matrix.Steric
		wantErr bool
	}{
		{"line point0", fields{&PointOverlay{matrix.LineMatrix{{100, 100}, {100, 101}}, matrix.Matrix{100, 100}}}, matrix.Matrix{100, 100}, false},
		{"line line0", fields{&PointOverlay{matrix.LineMatrix{{100, 100}, {100, 101}}, matrix.LineMatrix{{100, 100}, {100, 101}}}},
			matrix.Collection{matrix.LineMatrix{{100, 100}, {100, 101}}}, false},
		{"line line1", fields{&PointOverlay{matrix.LineMatrix{{100, 100}, {100, 101}}, matrix.LineMatrix{{100, 100}, {90, 102}}}},
			matrix.Collection{matrix.Matrix{100, 100}}, false},
		{"line poly1", fields{&PointOverlay{matrix.LineMatrix{{100, 100}, {101, 101}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		}},
			matrix.Collection{matrix.Matrix{101, 101}}, false},
		{"line poly2", fields{&PointOverlay{matrix.LineMatrix{{100, 100}, {100, 101}},
			matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		}},
			matrix.Collection{matrix.LineMatrix{{100, 100}, {100, 101}}}, false},
		{"line poly3", fields{&PointOverlay{matrix.LineMatrix{{100, 100}, {100, 101}},
			matrix.PolygonMatrix{{{105, 105}, {105, 101}, {101, 101}, {101, 105}, {105, 105}}},
		}},
			matrix.Collection{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &LineOverlay{
				PointOverlay: tt.fields.PointOverlay,
			}
			got, err := p.Intersection()
			if (err != nil) != tt.wantErr {
				t.Errorf("LineOverlay.Intersection() %v error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineOverlay.Intersection() %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestLineOverlay_Difference(t *testing.T) {
	type fields struct {
		PointOverlay *PointOverlay
	}
	tests := []struct {
		name    string
		fields  fields
		want    matrix.Steric
		wantErr bool
	}{
		{"line line0", fields{&PointOverlay{matrix.LineMatrix{{50, 100}, {50, 200}}, matrix.LineMatrix{{50, 50}, {50, 150}}}},
			matrix.Collection{matrix.LineMatrix{{50, 150}, {50, 200}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &LineOverlay{
				PointOverlay: tt.fields.PointOverlay,
			}
			got, err := p.Difference()
			if (err != nil) != tt.wantErr {
				t.Errorf("LineOverlay.Difference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineOverlay.Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
