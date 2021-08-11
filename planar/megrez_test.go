package planar

import (
	"testing"

	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/space"
)

func TestAlgorithm_IsClosed(t *testing.T) {
	const linestring = `LINESTRING(1 1,2 3,3 2,1 2)`
	const closedLinestring = `LINESTRING(1 1,2 3,3 2,1 2,1 1)`
	line, _ := wkt.UnmarshalString(linestring)
	closedLine, _ := wkt.UnmarshalString(closedLinestring)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "line", args: args{g: line}, want: false, wantErr: false},
		{name: "closedLine", args: args{g: closedLine}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.IsClosed(tt.args.g)
			t.Log(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("geometry: %s IsClose() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("geometry: %s IsClose() got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestAlgorithm_EqualsExact(t *testing.T) {
	geometry1, _ := wkt.UnmarshalString("POINT(116.309878625564 40.0427783817455)")
	geometry2, _ := wkt.UnmarshalString("POINT(116.309878725564 40.0427783827455)")
	geometry3, _ := wkt.UnmarshalString("POINT(116.309877625564 40.0427783827455)")
	type args struct {
		g1        space.Geometry
		g2        space.Geometry
		tolerance float64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "equals exact", args: args{g1: geometry1, g2: geometry2, tolerance: 0.000001}, want: true, wantErr: false},
		{name: "not equals exact", args: args{g1: geometry1, g2: geometry3, tolerance: 0.000001}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.EqualsExact(tt.args.g1, tt.args.g2, tt.args.tolerance)
			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.EqualsExact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GEOAlgorithm.EqualsExact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Equals(t *testing.T) {
	geometry1, _ := wkt.UnmarshalString("POINT(116.309878625564 40.0427783817455)")
	geometry2, _ := wkt.UnmarshalString("POINT(116.309878725564 40.0427783827455)")
	geometry3, _ := wkt.UnmarshalString("POINT(116.309878625564 40.0427783817455)")
	type args struct {
		g1        space.Geometry
		g2        space.Geometry
		tolerance float64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "equals exact", args: args{g1: geometry1, g2: geometry3, tolerance: 0.001}, want: true, wantErr: false},
		{name: "not equals exact", args: args{g1: geometry1, g2: geometry2, tolerance: 0.001}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Equals(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.Equals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GEOAlgorithm.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_IsSimple(t *testing.T) {
	const polygon = `POLYGON((1 2, 3 4, 5 6, 5 3, 1 2))`
	const linestring = `LINESTRING(1 1,2 2,2 3.5,1 3,1 2,2 1)`
	poly, _ := wkt.UnmarshalString(polygon)
	line, _ := wkt.UnmarshalString(linestring)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "polygon", args: args{g: poly}, want: true, wantErr: false},
		{name: "line", args: args{g: line}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			//G := GetStrategy(newGEOAlgorithm)
			got, err := G.IsSimple(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSimple() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsSimple() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_IsRing(t *testing.T) {
	const linestring1 = `LINESTRING(1 2, 3 4, 5 6, 5 3, 1 2)`
	const linestring2 = `LINESTRING(1 1,2 2,2 3.5,1 3,1 2,2 1)`
	line1, _ := wkt.UnmarshalString(linestring1)
	line2, _ := wkt.UnmarshalString(linestring2)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "line1", args: args{g: line1}, want: true, wantErr: false},
		{name: "line2", args: args{g: line2}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.IsRing(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsRing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsRing() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_IsEmpty(t *testing.T) {
	const polygon = `POLYGON((1 2, 3 4, 5 6, 1 2))`
	geometry, _ := wkt.UnmarshalString(polygon)
	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "empty", args: args{g: geometry}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.IsEmpty(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsEmpty() got = %v, want %v", got, tt.want)
			}
		})
	}
}
