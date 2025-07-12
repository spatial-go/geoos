package main

import (
"fmt"

"github.com/spatial-go/geoos/geoencoding"
"github.com/spatial-go/geoos/planar"
)

func main() {
// First, choose the default algorithm.
strategy := planar.NormalStrategy()

// Create two geometries for intersection
// Example 1: Line intersecting with a polygon
polygon, _ := geoencoding.Decode([]byte(`POLYGON((0 0,0 1,0 2,0 3,1 3,1 2,1 1,2 1,2 2,3 2,3 1,4 1,4 2,5 2,5 1,5 0,4 0,3 0,2 0,1 0,0 0))`), geoencoding.WKT)
line, _ := geoencoding.Decode([]byte(`LINESTRING(-5 0, 5 5)`), geoencoding.WKT)

// Calculate intersection
intersection, err := strategy.Intersection(polygon, line)
if err != nil {
fmt.Printf("Error: %v\n", err)
return
}

// Output result
result := string(geoencoding.Encode(intersection, geoencoding.WKT))
fmt.Printf("Intersection result: %s\n", result)
// Output: LINESTRING(0 2.5,1 3)

// Example 2: Two lines intersecting at a point
line1, _ := geoencoding.Decode([]byte(`LINESTRING(0 0, 2 2)`), geoencoding.WKT)
line2, _ := geoencoding.Decode([]byte(`LINESTRING(0 2, 2 0)`), geoencoding.WKT)

intersection2, err := strategy.Intersection(line1, line2)
if err != nil {
fmt.Printf("Error: %v\n", err)
return
}

result2 := string(geoencoding.Encode(intersection2, geoencoding.WKT))
fmt.Printf("Lines intersection: %s\n", result2)
// Output: POINT(1 1)
}
