package clusters

// func TestDistance(t *testing.T) {
// 	p1 := geoos.Point{2, 2}
// 	p2 := geoos.Point{3, 5}

// 	d := p1.Distance(p2.Coordinates())
// 	if d != 10 {
// 		t.Errorf("Expected distance of 10, got %f", d)
// 	}
// }

// func TestCenter(t *testing.T) {
// 	var o Observations
// 	o = append(o, Coordinates{1, 1})
// 	o = append(o, Coordinates{3, 2})
// 	o = append(o, Coordinates{5, 3})

// 	m, err := o.Center()
// 	if err != nil {
// 		t.Errorf("Could not retrieve center: %v", err)
// 		return
// 	}

// 	if m[0] != 3 || m[1] != 2 {
// 		t.Errorf("Expected coordinates [3 2], got %v", m)
// 	}
// }

// func TestAverageDistance(t *testing.T) {
// 	var o Observations
// 	o = append(o, Coordinates{1, 1})
// 	o = append(o, Coordinates{3, 2})
// 	o = append(o, Coordinates{5, 3})

// 	d := AverageDistance(o[0], o[1:])
// 	if d != 12.5 {
// 		t.Errorf("Expected average distance of 12.5, got %v", d)
// 	}

// 	d = AverageDistance(o[1], Observations{o[1]})
// 	if d != 0 {
// 		t.Errorf("Expected average distance of 0, got %v", d)
// 	}
// }

// func TestClusters(t *testing.T) {
// 	var o Observations
// 	o = append(o, Coordinates{1, 1})
// 	o = append(o, Coordinates{3, 2})
// 	o = append(o, Coordinates{5, 3})

// 	c, err := New(2, o)
// 	if err != nil {
// 		t.Errorf("Error seeding clusters: %v", err)
// 		return
// 	}

// 	if len(c) != 2 {
// 		t.Errorf("Expected 2 clusters, got %d", len(c))
// 		return
// 	}

// 	c[0].Append(o[0])
// 	c[1].Append(o[1])
// 	c[1].Append(o[2])
// 	c.Recenter()

// 	if n := c.Nearest(o[1]); n != 1 {
// 		t.Errorf("Expected nearest cluster 1, got %d", n)
// 	}

// 	nc, d := c.Neighbour(o[0], 0)
// 	if nc != 1 {
// 		t.Errorf("Expected neighbouring cluster 1, got %d", nc)
// 	}
// 	if d != 12.5 {
// 		t.Errorf("Expected neighbouring cluster distance 12.5, got %f", d)
// 	}

// 	if pp := c[1].PointsInDimension(0); pp[0] != 3 || pp[1] != 5 {
// 		t.Errorf("Expected [3 5] as points in dimension 0, got %v", pp)
// 	}
// 	if pp := c.CentersInDimension(0); pp[0] != 1 || pp[1] != 4 {
// 		t.Errorf("Expected [1 4] as centers in dimension 0, got %v", pp)
// 	}

// 	c.Reset()
// 	if len(c[0].Observations) > 0 {
// 		t.Errorf("Expected empty cluster 1, found %d observations", len(c[0].Observations))
// 	}
// }
