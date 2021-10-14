package overlay

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos"
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
			matrix.Collection{matrix.LineMatrix{{100, 100}, {101, 101}}}, false},
		{"line poly2", fields{&PointOverlay{matrix.LineMatrix{{100, 100}, {100, 101}},
			matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		}},
			matrix.Collection{matrix.LineMatrix{{100, 100}, {100, 101}}}, false},
		{"line poly3", fields{&PointOverlay{matrix.LineMatrix{{100, 100}, {100, 101}},
			matrix.PolygonMatrix{{{105, 105}, {105, 101}, {101, 101}, {101, 105}, {105, 105}}},
		}},
			matrix.Collection{}, false},

		{"line poly5", fields{&PointOverlay{matrix.LineMatrix{{111.98638916015625, 38.50357937743225}, {111.96372985839844, 38.42723559654225},
			{111.91085815429688, 38.344887442462806}, {111.87309265136717, 38.24680876017446}, {111.84906005859375, 38.15075747130226}},
			matrix.PolygonMatrix{{{112.01248168945312, 38.4519755295767}, {111.7096710205078, 38.39441521865825}, {111.90948486328125, 38.176671418717746},
				{112.01248168945312, 38.4519755295767}}},
		}},
			matrix.Collection{matrix.LineMatrix{{111.96859688870948, 38.44363360825604}, {111.96372985839844, 38.42723559654225}, {111.91085815429688, 38.344887442462806},
				{111.87309265136717, 38.24680876017446}, {111.8671003424276, 38.22285924285058},
			}},
			false,
		},
	}
	for _, tt := range tests {
		if !geoos.GeoosTestTag && tt.name != "line poly5" {
			continue
		}
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
		{"line line0", fields{&PointOverlay{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}}, matrix.LineMatrix{{50, 50}, {50, 150}}}},
			matrix.Collection{matrix.LineMatrix{{50, 150}, {50, 200}, {60, 200}}}, false},
		{"line line1", fields{&PointOverlay{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}}, matrix.LineMatrix{{50, 120}, {50, 150}}}},
			matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 120}}, matrix.LineMatrix{{50, 150}, {50, 200}, {60, 200}}}, false},
		{"line line2", fields{&PointOverlay{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}}, matrix.LineMatrix{{50, 150}, {50, 250}}}},
			matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 150}}, matrix.LineMatrix{{50, 200}, {60, 200}}}, false},
		{"line line3", fields{&PointOverlay{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}}, matrix.LineMatrix{{50, 100}, {50, 150}}}},
			matrix.Collection{matrix.LineMatrix{{50, 150}, {50, 200}, {60, 200}}}, false},
		{"line line4", fields{&PointOverlay{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}}, matrix.LineMatrix{{50, 150}, {50, 200}}}},
			matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 150}}, matrix.LineMatrix{{50, 200}, {60, 200}}}, false},
		{"line line5", fields{&PointOverlay{matrix.LineMatrix{{50, 100}, {50, 200}}, matrix.LineMatrix{{50, 50}, {50, 250}}}},
			matrix.Collection{}, false},
		{"line line6", fields{&PointOverlay{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}}, matrix.LineMatrix{{50, 50}, {50, 250}}}},
			matrix.Collection{matrix.LineMatrix{{50, 200}, {60, 200}}}, false},

		{"line line7", fields{&PointOverlay{matrix.LineMatrix{{50, 100}, {50, 200}}, matrix.LineMatrix{{30, 30}, {30, 150}}}},
			matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 200}}}, false},
		{"line line8", fields{&PointOverlay{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}}, matrix.LineMatrix{{30, 150}, {60, 150}}}},
			matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 150}}, matrix.LineMatrix{{50, 150}, {50, 200}, {60, 200}}}, false},

		{"line poly1", fields{&PointOverlay{matrix.LineMatrix{{111.98638916015625, 38.50357937743225}, {111.96372985839844, 38.42723559654225},
			{111.91085815429688, 38.344887442462806}, {111.87309265136717, 38.24680876017446}, {111.84906005859375, 38.15075747130226}},
			matrix.PolygonMatrix{{{112.01248168945312, 38.4519755295767}, {111.7096710205078, 38.39441521865825}, {111.90948486328125, 38.176671418717746},
				{112.01248168945312, 38.4519755295767}}},
		}},
			matrix.Collection{matrix.LineMatrix{{111.98638916015625, 38.50357937743225}, {111.96859688870948, 38.44363360825604}},
				matrix.LineMatrix{{111.96859688870948, 38.44363360825604}, {111.96372985839844, 38.42723559654225}, {111.91085815429688, 38.344887442462806},
					{111.87309265136717, 38.24680876017446}, {111.8671003424276, 38.22285924285058}},
				matrix.LineMatrix{{111.8671003424276, 38.22285924285058}, {111.84906005859375, 38.15075747130226}},
			}, false},
		{name: "line poly6", fields: fields{&PointOverlay{matrix.LineMatrix{{200, 300}, {500, 300}, {500, 600}, {800, 900}},
			matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}},
		}},
			want: matrix.Collection{matrix.LineMatrix{{500, 500}, {500, 600}, {800, 900}}, matrix.LineMatrix{{200, 300}, {300, 300}}},
		},
	}
	for _, tt := range tests {
		if !geoos.GeoosTestTag && tt.name != "line poly1" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			p := &LineOverlay{
				PointOverlay: tt.fields.PointOverlay,
			}
			got, err := p.Difference()
			if (err != nil) != tt.wantErr {
				t.Errorf("LineOverlay.Difference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("LineOverlay.Difference()%v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
