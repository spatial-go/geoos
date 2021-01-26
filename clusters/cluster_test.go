package clusters

import (
	"testing"

	"github.com/spatial-go/geoos"
)

func TestCenter(t *testing.T) {
	var o Points
	o = append(o, geoos.Point{1, 1})
	o = append(o, geoos.Point{3, 2})
	o = append(o, geoos.Point{5, 3})

	m, err := o.Center()
	if err != nil {
		t.Errorf("Could not retrieve center: %v", err)
		return
	}

	if m[0] != 3 || m[1] != 2 {
		t.Errorf("Expected coordinates [3 2], got %v", m)
	}
}

func TestAverageDistance(t *testing.T) {
	var o Points
	o = append(o, geoos.Point{1, 1})
	o = append(o, geoos.Point{4, 5})
	o = append(o, geoos.Point{6, 13})

	d := AverageDistance(o[0], o[1:])
	if d != 9 {
		t.Errorf("Expected average distance of 12.5, got %v", d)
	}

	d = AverageDistance(o[1], Points{o[1]})
	if d != 0 {
		t.Errorf("Expected average distance of 0, got %v", d)
	}
}

func TestClusters(t *testing.T) {
	var o Points
	o = append(o, geoos.Point{1, 1})
	o = append(o, geoos.Point{4, 5})
	o = append(o, geoos.Point{6, 13})

	c, err := New(2, o)
	if err != nil {
		t.Errorf("Error seeding clusters: %v", err)
		return
	}

	if len(c) != 2 {
		t.Errorf("Expected 2 clusters, got %d", len(c))
		return
	}

	c[0].Append(o[0])
	c[1].Append(o[1])
	c[1].Append(o[2])
	c.Recenter()

	if n := c.Nearest(o[1]); n != 1 {
		t.Errorf("Expected nearest cluster 1, got %d", n)
	}

	nc, d := c.Neighbour(o[0], 0)
	if nc != 1 {
		t.Errorf("Expected neighbouring cluster 1, got %d", nc)
	}
	if d != 9 {
		t.Errorf("Expected neighbouring cluster distance 9, got %f", d)
	}

	if pp := c[1].PointsInDimension(0); pp[0] != 4 || pp[1] != 6 {
		t.Errorf("Expected [4 6] as points in dimension 0, got %v", pp)
	}
	if pp := c.CentersInDimension(0); pp[0] != 1 || pp[1] != 5 {
		t.Errorf("Expected [1 5] as centers in dimension 0, got %v", pp)
	}

	c.Reset()
	if len(c[0].Points) > 0 {
		t.Errorf("Expected empty cluster 1, found %d observations", len(c[0].Points))
	}
}
