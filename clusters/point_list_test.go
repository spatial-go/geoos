package clusters

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/space"
)

func TestConvexHull(t *testing.T) {
	// empty list
	points := PointList{}
	expected := PointList{}
	testExample(t, points, expected)

	// single point
	points = PointList{{1, 1}}
	expected = PointList{{1, 1}}
	testExample(t, points, expected)

	// two points
	points = PointList{{1, 1}, {1, 2}}
	expected = PointList{{1, 1}, {1, 2}}
	testExample(t, points, expected)

	// line
	points = PointList{{1, 1}, {2, 2}, {3, 3}}
	expected = PointList{{1, 1}, {3, 3}} // intermediate point omitted
	testExample(t, points, expected)

	// triangle
	points = PointList{{2, 2}, {1, 2}, {1, 1}}
	expected = PointList{{1, 1}, {2, 2}, {1, 2}, {1, 1}} // closed polygon shape
	testExample(t, points, expected)

	// square
	points = PointList{{1, 1}, {2, 1}, {2, 2}, {1, 2}}
	expected = PointList{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}} // closed polygon-shape
	testExample(t, points, expected)

	// generate random point cloud with set outer bound
	size := 50
	min, max := 0., 90.
	points = make(PointList, size)
	for i := 0; i < size; i++ {
		points[i] = space.Point{min + rand.Float64()*max, min + rand.Float64()*max}
	}
	points = append(points, PointList{{min - 1, min - 1}, {min - 1, max + 1}, {max + 1, max + 1}, {max + 1, min - 1}}...)    // square outer boundary in clockwise order
	expected = PointList{{min - 1, min - 1}, {max + 1, min - 1}, {max + 1, max + 1}, {min - 1, max + 1}, {min - 1, min - 1}} // closed polygon-shape in counter-clockwise order
	testExample(t, points, expected)
}

func testExample(t *testing.T, points PointList, expected PointList) {
	hull := points.ConvexHull()
	if !reflect.DeepEqual(expected, hull) {
		t.Errorf("Expected %v but got %v", expected, hull)
		t.FailNow()
	}
}
