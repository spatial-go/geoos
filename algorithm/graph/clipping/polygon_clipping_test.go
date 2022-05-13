package clipping

import (
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/overlay"
)

func TestPolygonClipping_Intersection(t *testing.T) {
	type fields struct {
		PointClipping *PointClipping
		subjectPlane  *overlay.Plane
		clippingPlane *overlay.Plane
	}
	tests := []struct {
		name    string
		fields  fields
		want    matrix.Steric
		wantErr bool
	}{
		{"poly point0", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}, matrix.Matrix{100, 100}}}, matrix.Matrix{100, 100}, false},
		{"poly line0", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
			matrix.LineMatrix{{100, 100}, {101, 101}}}},
			matrix.LineMatrix{{100, 100}, {101, 101}}, false},
		{"poly line1", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{105, 105}, {105, 101}, {101, 101}, {101, 105}, {105, 105}}},
			matrix.LineMatrix{{100, 100}, {90, 101}}}},
			matrix.Collection{}, false},

		{"poly poly1", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		}},
			matrix.PolygonMatrix{{{101, 100}, {100, 100}, {100, 101}, {101, 101}, {101, 100}}}, false},

		{"poly poly2", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{105, 105}, {105, 103}, {103, 103}, {103, 105}, {105, 105}}},
			matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		}},
			matrix.PolygonMatrix{}, false},
		{"poly poly3", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		}},
			matrix.PolygonMatrix{{{5, 10}, {10, 10}, {10, 5}, {5, 5}, {5, 10}}}, false},
		{"poly poly3-1", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 5}, {5, 15}, {15, 15}, {15, 5}, {5, 5}}},
		}},
			matrix.PolygonMatrix{{{5, 10}, {10, 10}, {10, 5}, {5, 5}, {5, 10}}}, false},
		{"poly poly4", fields{PointClipping: &PointClipping{
			matrix.PolygonMatrix{{{111.30523681640625, 38.117271658305}, {112.34344482421875, 38.11727165830543}, {112.34344482421875, 38.89103282648846},
				{111.30523681640625, 38.89103282648846}, {111.30523681640625, 38.117271658305}}},
			matrix.PolygonMatrix{{{111.50848388671875, 37.6359849542696}, {112.64007568359375, 37.6359849542696}, {112.64007568359375, 38.35027253825765},
				{111.50848388671875, 38.35027253825765}, {111.50848388671875, 37.6359849542696}}},
		}},
			matrix.PolygonMatrix{{{112.34344482421875, 38.35027253825765}, {112.34344482421875, 38.11727165830543}, {111.50848388671875, 38.11727165830543},
				{111.50848388671875, 38.35027253825765}, {112.34344482421875, 38.35027253825765}}}, false},

		{"poly poly5", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}, {{1, 1}, {9, 1}, {9, 9}, {1, 9}, {1, 1}}},
			matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		}},
			matrix.PolygonMatrix{{{9, 9}, {9, 5}, {10, 5}, {10, 10}, {5, 10}, {5, 9}, {9, 9}}}, false},
		{"poly poly5-1", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
			matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}, {{1, 1}, {9, 1}, {9, 9}, {1, 9}, {1, 1}}},
		}},
			matrix.PolygonMatrix{{{9, 9}, {9, 5}, {10, 5}, {10, 10}, {5, 10}, {5, 9}, {9, 9}}}, false},
	}
	for _, tt := range tests {
		if !geoos.GeoosTestTag &&
			tt.name != "poly line0" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			p := &PolygonClipping{
				PointClipping: tt.fields.PointClipping,
				subjectPlane:  tt.fields.subjectPlane,
				clippingPlane: tt.fields.clippingPlane,
			}
			got, err := p.Intersection()
			if (err != nil) != tt.wantErr {
				t.Errorf("PolygonClipping.Intersection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Proximity(tt.want) {
				if gotPoly, ok := got.(matrix.PolygonMatrix); ok {
					if wantPoly, ok := tt.want.(matrix.PolygonMatrix); ok {
						if measure.AreaOfPolygon(gotPoly) == measure.AreaOfPolygon(wantPoly) {
							return
						}
					}
				}
				t.Errorf("PolygonClipping.Intersection()%v = %v, \nwant %v type %T, want %T", tt.name, got, tt.want, got, tt.want)
			}
		})
	}
}

func TestPolygonClipping_Union(t *testing.T) {
	type fields struct {
		PointClipping *PointClipping
		subjectPlane  *overlay.Plane
		clippingPlane *overlay.Plane
	}
	tests := []struct {
		name    string
		fields  fields
		want    []matrix.Steric
		wantErr bool
	}{
		{"poly poly", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		}},
			[]matrix.Steric{matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}}}}, false},

		{"poly poly01", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 5}, {5, 15}, {15, 15}, {15, 5}, {5, 5}}},
		}},
			[]matrix.Steric{
				matrix.PolygonMatrix{{{10, 5}, {10, 0}, {0, 0}, {0, 10}, {5, 10}, {5, 15}, {15, 15}, {15, 5}, {10, 5}}},
				matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}}},
			}, false},

		{"poly poly02", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}, {{1, 1}, {9, 1}, {9, 9}, {1, 9}, {1, 1}}},
			matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		}},
			[]matrix.Steric{matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}}, {{5, 9}, {1, 9}, {1, 1}, {9, 1}, {9, 5}, {5, 5}, {5, 9}}}}, false},

		{"poly poly03", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
			matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}, {{1, 1}, {9, 1}, {9, 9}, {1, 9}, {1, 1}}},
		}},
			[]matrix.Steric{matrix.PolygonMatrix{{{10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}}, {{9, 5}, {5, 5}, {5, 9}, {1, 9}, {1, 1}, {9, 1}, {9, 5}}}}, false},

		{"poly poly1", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		}},
			[]matrix.Steric{matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}}, false},

		{name: "poly 1",
			fields: fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				matrix.PolygonMatrix{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}},
			},
			}, want: []matrix.Steric{matrix.Collection{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				matrix.PolygonMatrix{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}}},
				matrix.Collection{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
					matrix.PolygonMatrix{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}}},
			},
			wantErr: false},

		{name: "poly 2",
			fields: fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				matrix.PolygonMatrix{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}},
			},
			}, want: []matrix.Steric{matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
				matrix.PolygonMatrix{{{2, 2}, {1, 2}, {1, 1}, {2, 1}, {5, 1}, {5, 2}, {2, 2}}},
			},
			wantErr: false},

		{name: "poly 3",
			fields: fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 3}, {2, 3}, {2, 2}}},
			},
			}, want: []matrix.Steric{matrix.Collection{matrix.PolygonMatrix{{{2, 2}, {2, 1}, {1, 1}, {1, 2}, {2, 2}}},
				matrix.PolygonMatrix{{{2, 2}, {2, 3}, {5, 3}, {5, 2}, {2, 2}}}},
				matrix.Collection{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
					matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 3}, {2, 3}, {2, 2}}}},
			},
			wantErr: false},

		{name: "poly 4",
			fields: fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{1, 2}, {3, 2}, {3, 3}, {1, 3}, {1, 2}}},
				matrix.PolygonMatrix{{{2, 1}, {5, 1}, {5, 5}, {2, 5}, {2, 1}}},
			},
			}, want: []matrix.Steric{matrix.PolygonMatrix{{{2, 2}, {1, 2}, {1, 3}, {2, 3}, {2, 5}, {5, 5}, {5, 1}, {2, 1}, {2, 2}}},
				matrix.PolygonMatrix{{{2, 3}, {1, 3}, {1, 2}, {2, 2}, {2, 1}, {5, 1}, {5, 5}, {2, 5}, {2, 3}}},
			},
			wantErr: false},

		{name: "poly 5",
			fields: fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
				matrix.PolygonMatrix{{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}},
			},
			}, want: []matrix.Steric{matrix.PolygonMatrix{{{1, 1}, {1, 5}, {5, 5}, {5, 1}, {1, 1}}},
				matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
			},
			wantErr: false},

		{name: "poly 6",
			fields: fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 3}, {1, 3}, {1, 1}}},
			},
			}, want: []matrix.Steric{matrix.PolygonMatrix{{{2, 1}, {1, 1}, {1, 2}, {1, 3}, {5, 3}, {5, 1}, {2, 1}}},
				matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 3}, {1, 3}, {1, 1}}},
			},
			wantErr: false},
	}
	for _, tt := range tests {
		if !geoos.GeoosTestTag &&
			tt.name != "poly 2" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			p := &PolygonClipping{
				PointClipping: tt.fields.PointClipping,
				subjectPlane:  tt.fields.subjectPlane,
				clippingPlane: tt.fields.clippingPlane,
			}
			got, err := p.Union()
			if (err != nil) != tt.wantErr {
				t.Errorf("PolygonClipping.Union() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			isEqual := got.Proximity(tt.want[0])
			if len(tt.want) > 1 {
				isEqual1 := got.Proximity(tt.want[1])
				isEqual = isEqual || isEqual1
			}

			if !isEqual {
				t.Errorf("PolygonClipping.Union()%v = %v, \nwant %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPolygonClipping_Difference(t *testing.T) {
	type fields struct {
		PointClipping *PointClipping
		subjectPlane  *overlay.Plane
		clippingPlane *overlay.Plane
	}
	tests := []struct {
		name    string
		fields  fields
		want    matrix.Steric
		wantErr bool
	}{
		{"poly poly", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		}},
			matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {5, 5}, {5, 10}}}, false},

		{"poly poly1", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		}},
			matrix.PolygonMatrix{}, false},

		{"poly poly2", fields{PointClipping: &PointClipping{
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
			matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		}},
			matrix.PolygonMatrix{{{100, 101}, {90, 101}, {90, 90}, {101, 90}, {101, 100}, {100, 100}, {100, 101}}}, false},
		{"poly poly3", fields{PointClipping: &PointClipping{
			matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}},
		}},
			matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}, {{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}}, false},
	}
	for _, tt := range tests {
		if !geoos.GeoosTestTag &&
			tt.name != "poly poly2" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			p := &PolygonClipping{
				PointClipping: tt.fields.PointClipping,
				subjectPlane:  tt.fields.subjectPlane,
				clippingPlane: tt.fields.clippingPlane,
			}
			got, err := p.Difference()
			if (err != nil) != tt.wantErr {
				t.Errorf("PolygonClipping.Difference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Proximity(tt.want) {
				t.Errorf("PolygonClipping.Difference()%v = %v, \nwant %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPolygonClipping_SymDifference(t *testing.T) {
	type fields struct {
		PointClipping *PointClipping
		subjectPlane  *overlay.Plane
		clippingPlane *overlay.Plane
	}
	tests := []struct {
		name    string
		fields  fields
		want    matrix.Steric
		wantErr bool
	}{
		{"poly poly", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		}},
			matrix.Collection{matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {5, 5}, {5, 10}}},
				matrix.PolygonMatrix{{{10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}, {10, 10}, {10, 5}}}}, false},
		{"poly poly0", fields{PointClipping: &PointClipping{matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		}},
			matrix.Collection{
				matrix.PolygonMatrix{{{100, 101}, {90, 101}, {90, 90}, {101, 90}, {101, 100}, {100, 100}, {100, 101}}}},
			false},
		{"poly poly1", fields{PointClipping: &PointClipping{
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
			matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		}},
			matrix.Collection{matrix.PolygonMatrix{{{100, 101}, {90, 101}, {90, 90}, {101, 90}, {101, 100}, {100, 100}, {100, 101}}}},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PolygonClipping{
				PointClipping: tt.fields.PointClipping,
				subjectPlane:  tt.fields.subjectPlane,
				clippingPlane: tt.fields.clippingPlane,
			}
			got, err := p.SymDifference()
			if (err != nil) != tt.wantErr {
				t.Errorf("PolygonClipping.SymDifference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Proximity(tt.want) {
				t.Errorf("PolygonClipping.SymDifference() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}
