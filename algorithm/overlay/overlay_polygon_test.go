package overlay

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestPolygonOverlay_Intersection(t *testing.T) {
	type fields struct {
		PointOverlay  *PointOverlay
		subjectPlane  *Plane
		clippingPlane *Plane
	}
	tests := []struct {
		name    string
		fields  fields
		want    matrix.Steric
		wantErr bool
	}{
		{"poly point0", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}, matrix.Matrix{100, 100}}}, matrix.Matrix{100, 100}, false},
		{"poly line0", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
			matrix.LineMatrix{{100, 100}, {101, 101}}}},
			matrix.Collection{matrix.Matrix{101, 101}}, false},
		{"poly line1", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{105, 105}, {105, 101}, {101, 101}, {101, 105}, {105, 105}}},
			matrix.LineMatrix{{100, 100}, {90, 101}}}},
			matrix.Collection{}, false},

		{"poly poly2", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{105, 105}, {105, 103}, {103, 103}, {103, 105}, {105, 105}}},
			matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		}},
			matrix.PolygonMatrix{}, false},
		{"poly poly3", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		}},
			matrix.PolygonMatrix{{{5, 10}, {10, 10}, {10, 5}, {5, 5}, {5, 10}}}, false},

		{"poly poly1", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		}},
			matrix.PolygonMatrix{{{101, 100}, {100, 100}, {100, 101}, {101, 101}, {101, 100}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PolygonOverlay{
				PointOverlay:  tt.fields.PointOverlay,
				subjectPlane:  tt.fields.subjectPlane,
				clippingPlane: tt.fields.clippingPlane,
			}
			got, err := p.Intersection()
			if (err != nil) != tt.wantErr {
				t.Errorf("PolygonOverlay.Intersection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("PolygonOverlay.Intersection()%v = %v, \nwant %v type %T, want %T", tt.name, got, tt.want, got, tt.want)
			}
		})
	}
}

func TestPolygonOverlay_Union(t *testing.T) {
	type fields struct {
		PointOverlay  *PointOverlay
		subjectPlane  *Plane
		clippingPlane *Plane
	}
	tests := []struct {
		name    string
		fields  fields
		want    matrix.Steric
		wantErr bool
	}{
		{"poly poly", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		}},
			matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}}}, false},

		{"poly poly1", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		}},
			matrix.PolygonMatrix{{{101, 100}, {101, 101}, {100, 101}, {90, 101}, {90, 90}, {101, 90}, {101, 100}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PolygonOverlay{
				PointOverlay:  tt.fields.PointOverlay,
				subjectPlane:  tt.fields.subjectPlane,
				clippingPlane: tt.fields.clippingPlane,
			}
			got, err := p.Union()
			if (err != nil) != tt.wantErr {
				t.Errorf("PolygonOverlay.Union() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PolygonOverlay.Union() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}

func TestPolygonOverlay_Difference(t *testing.T) {
	type fields struct {
		PointOverlay  *PointOverlay
		subjectPlane  *Plane
		clippingPlane *Plane
	}
	tests := []struct {
		name    string
		fields  fields
		want    matrix.Steric
		wantErr bool
	}{
		{"poly poly", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		}},
			matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {5, 5}, {5, 10}}}, false},

		{"poly poly1", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		}},
			matrix.PolygonMatrix{}, false},
		{"poly poly1", fields{PointOverlay: &PointOverlay{
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
			matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		}},
			matrix.PolygonMatrix{{{100, 101}, {90, 101}, {90, 90}, {101, 90}, {101, 100}, {100, 100}, {100, 101}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PolygonOverlay{
				PointOverlay:  tt.fields.PointOverlay,
				subjectPlane:  tt.fields.subjectPlane,
				clippingPlane: tt.fields.clippingPlane,
			}
			got, err := p.Difference()
			if (err != nil) != tt.wantErr {
				t.Errorf("PolygonOverlay.Difference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PolygonOverlay.Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolygonOverlay_SymDifference(t *testing.T) {
	type fields struct {
		PointOverlay  *PointOverlay
		subjectPlane  *Plane
		clippingPlane *Plane
	}
	tests := []struct {
		name    string
		fields  fields
		want    matrix.Steric
		wantErr bool
	}{
		{"poly poly", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		}},
			matrix.Collection{matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {5, 5}, {5, 10}}},
				matrix.PolygonMatrix{{{10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}, {10, 10}, {10, 5}}}}, false},

		{"poly poly1", fields{PointOverlay: &PointOverlay{matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		}},
			matrix.Collection{
				matrix.PolygonMatrix{{{100, 101}, {90, 101}, {90, 90}, {101, 90}, {101, 100}, {100, 100}, {100, 101}}}},
			false},
		{"poly poly1", fields{PointOverlay: &PointOverlay{
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
			matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		}},
			matrix.Collection{matrix.PolygonMatrix{{{100, 101}, {90, 101}, {90, 90}, {101, 90}, {101, 100}, {100, 100}, {100, 101}}}},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PolygonOverlay{
				PointOverlay:  tt.fields.PointOverlay,
				subjectPlane:  tt.fields.subjectPlane,
				clippingPlane: tt.fields.clippingPlane,
			}
			got, err := p.SymDifference()
			if (err != nil) != tt.wantErr {
				t.Errorf("PolygonOverlay.SymDifference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PolygonOverlay.SymDifference() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}
