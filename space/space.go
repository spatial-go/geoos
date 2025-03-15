// File: space/space.go
package space

import (
    "fmt"
)

// ToMultiPart converts a single-part geometry to its multi-part equivalent.
// If the geometry is already multi-part, it remains unchanged.
func ToMultiPart(g Geometry) (Geometry, error) {
    switch geom := g.Geom().(type) {
    case Point:
        // Convert single Point to MultiPoint
        return MultiPoint{geom}, nil
    case LineString:
        // Convert single LineString to MultiLineString
        return MultiLineString{geom}, nil
    case Polygon:
        // Convert single Polygon to MultiPolygon
        return MultiPolygon{geom}, nil
    case MultiPoint, MultiLineString, MultiPolygon:
        // Already multi-part, return unchanged
        return g, nil
    default:
        return nil, fmt.Errorf("unsupported geometry type: %T", geom)
    }
}