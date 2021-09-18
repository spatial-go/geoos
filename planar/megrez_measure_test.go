package planar

import (
	"testing"

	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/space"
)

func TestAlgorithm_Area(t *testing.T) {
	const polygon = `POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))`
	geometry, _ := wkt.UnmarshalString(polygon)
	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "area", args: args{g: geometry}, want: 4.0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalStrategy().Area(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Length(t *testing.T) {
	const line = `LINESTRING(0 0, 1 1)`
	geometry, _ := wkt.UnmarshalString(line)
	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "lengh", args: args{g: geometry}, want: 1.4142135623730951, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Length(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Length() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Length() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Distance(t *testing.T) {
	point01, _ := wkt.UnmarshalString(`POINT(1 3)`)
	point02, _ := wkt.UnmarshalString(`POINT(4 7)`)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "distance", args: args{g1: point01, g2: point02}, want: 5, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Distance(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Distance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_SphericalDistance(t *testing.T) {
	point01 := space.Point{116.397439, 39.909177}
	point02 := space.Point{116.397725, 39.903079}
	point03 := space.Point{118.1487, 39.586671}

	type args struct {
		p1 space.Point
		p2 space.Point
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "SphericalDistance", args: args{p1: point01, p2: point02}, want: 678.5053586786567, wantErr: false},
		{name: "SphericalDistance", args: args{p1: point01, p2: point03}, want: 153953.98145619757, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, _ := G.SphericalDistance(tt.args.p1, tt.args.p2)
			if got != tt.want {
				t.Errorf("SphericalDistance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_NGeometry(t *testing.T) {
	multiPoint, _ := wkt.UnmarshalString(`MULTIPOINT ( -1 0, -1 2, -1 3, -1 4, -1 7, 0 1, 0 3, 1 1, 2 0, 6 0, 7 8, 9 8, 10 6 )`)
	multiLineString, _ := wkt.UnmarshalString(`MULTILINESTRING((10 130,50 190,110 190,140 150,150 80,100 10,20 40,10 130),
	(70 40,100 50,120 80,80 110,50 90,70 40))`)
	multiPolygon, _ := wkt.UnmarshalString(`MULTIPOLYGON (((40 40, 20 45, 45 30, 40 40)),
	((20 35, 10 30, 10 10, 30 5, 45 20, 20 35)),
	((30 20, 20 15, 20 25, 30 20)))`)
	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "ngeometry multiPoint", args: args{g: multiPoint}, want: 13, wantErr: false},
		{name: "ngeometry multiLineString", args: args{g: multiLineString}, want: 2, wantErr: false},
		{name: "ngeometry multiPolygon", args: args{g: multiPolygon}, want: 3, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.NGeometry(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.NGeometry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GEOAlgorithm.NGeometry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_HausdorffDistance(t *testing.T) {
	g3 := "LINESTRING (130 0, 0 0, 0 150)"
	g4 := "LINESTRING (10 10, 10 150, 130 10)"

	const g1 = `LINESTRING (0 0, 2 0)`
	const g2 = `MULTIPOINT (0 1, 1 0, 2 1)`
	geom1, _ := wkt.UnmarshalString(g1)
	geom2, _ := wkt.UnmarshalString(g2)
	geom3, _ := wkt.UnmarshalString(g3)
	geom4, _ := wkt.UnmarshalString(g4)
	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "HausdorffDistance", args: args{
			g1: geom1,
			g2: geom2,
		}, want: 1, wantErr: false},
		{name: "HausdorffDistance", args: args{
			g1: geom3,
			g2: geom4,
		}, want: 14.142135623730951, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.HausdorffDistance(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("HausdorffDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HausdorffDistance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMegrezAlgorithm_HausdorffDistanceDensify(t *testing.T) {

	g3 := "LINESTRING (130 0, 0 0, 0 150)"
	g4 := "LINESTRING (10 10, 10 150, 130 10)"

	const g1 = `LINESTRING (0 0, 2 0)`
	const g2 = `MULTIPOINT (0 1, 1 0, 2 1)`
	geom1, _ := wkt.UnmarshalString(g1)
	geom2, _ := wkt.UnmarshalString(g2)
	geom3, _ := wkt.UnmarshalString(g3)
	geom4, _ := wkt.UnmarshalString(g4)
	type args struct {
		geom1       space.Geometry
		geom2       space.Geometry
		densifyFrac float64
	}
	tests := []struct {
		name    string
		g       *megrezAlgorithm
		args    args
		want    float64
		wantErr bool
	}{
		{name: "HausdorffDistanceDensify", args: args{
			geom1:       geom1,
			geom2:       geom2,
			densifyFrac: 0.001,
		}, want: 1, wantErr: false},
		{name: "HausdorffDistanceDensify", args: args{
			geom1:       geom3,
			geom2:       geom4,
			densifyFrac: 1,
		}, want: 14.142135623730951, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NormalStrategy()
			got, err := g.HausdorffDistanceDensify(tt.args.geom1, tt.args.geom2, tt.args.densifyFrac)
			if (err != nil) != tt.wantErr {
				t.Errorf("MegrezAlgorithm.HausdorffDistanceDensify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MegrezAlgorithm.HausdorffDistanceDensify() = %v, want %v", got, tt.want)
			}
		})
	}
}
