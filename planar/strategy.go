package planar

import (
    "github.com/spatial-go/geoos/space"
)

type Strategy interface {
    Area(space.Geometry) (float64, error)
    ToMultiPart(g space.Geometry) (space.Geometry, error)
}

type normalStrategy struct{}

func NormalStrategy() Strategy {
    return normalStrategy{}
}

func (s normalStrategy) ToMultiPart(g space.Geometry) (space.Geometry, error) {
    return space.ToMultiPart(g)
}

func (s normalStrategy) Area(g space.Geometry) (float64, error) {
    return 0, nil // Placeholder
}