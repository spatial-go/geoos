package topograph

import (
	"testing"

	"github.com/spatial-go/geoos/space"
)

type args struct {
	A space.Geometry
	B space.Geometry
}

var tr = GetRelationship(NewTopograph)

func TestTopological_Within(t *testing.T) {

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"within 1", args{space.Point{2, 2}, space.LineString{{1, 1}, {3, 3}}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.Within(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topological.Within() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Topological.Within() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopological_Contains(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Contains 0", args{space.Point{2, 2}, space.LineString{{1, 1}, {3, 3}}}, false, false},
		{"Contains 1",
			args{space.Polygon{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.Contains(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topological.Contains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Topological.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopological_Covers(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Covers 0", args{space.Point{2, 2}, space.LineString{{1, 1}, {3, 3}}}, false, false},
		{"Covers 1",
			args{space.Polygon{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.Covers(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topological.Covers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Topological.Covers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopological_CoveredBy(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"CoveredBy 0", args{space.Point{2, 2}, space.LineString{{1, 1}, {3, 3}}}, true, false},
		{"CoveredBy 1",
			args{space.Polygon{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.CoveredBy(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topological.CoveredBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Topological.CoveredBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopological_Crosses(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Crosses 0", args{space.Point{2, 2}, space.LineString{{1, 1}, {3, 3}}}, false, false},
		{"Crosses 1",
			args{space.Polygon{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.Crosses(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topological.Crosses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Topological.Crosses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopological_Disjoint(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Disjoint 0", args{space.Point{2, 2}, space.LineString{{1, 1}, {3, 3}}}, false, false},
		{"Disjoint 1",
			args{space.Polygon{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.Disjoint(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topological.Disjoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Topological.Disjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopological_Intersects(t *testing.T) {

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Intersects 0", args{space.Point{2, 2}, space.LineString{{1, 1}, {3, 3}}}, true, false},
		{"Intersects 1",
			args{space.Polygon{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.Intersects(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topological.Intersects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Topological.Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopological_Touches(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Touches 0", args{space.Point{2, 2}, space.LineString{{1, 1}, {3, 3}}}, false, false},
		{"Touches 1",
			args{space.Polygon{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.Touches(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topological.Touches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Topological.Touches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopological_Overlaps(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Touches 0", args{space.Point{2, 2}, space.LineString{{1, 1}, {3, 3}}}, false, false},
		{"Touches 1",
			args{space.Polygon{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tr.Overlaps(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topological.Overlaps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Topological.Overlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
