package sharedpaths

import (
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay"
)

// Computer returns a collection containing paths shared by the two input geometries.
// Those going in the same direction are in the first element of the collection,
// those going in the opposite direction are in the second element.
// The paths themselves are given in the direction of the first geometry.
type Computer struct {
	g1, g2 matrix.Steric
}

// SharedPaths create  SharedPaths with geom.
func SharedPaths(g1, g2 matrix.Steric) (forwDir matrix.Collection, backDir matrix.Collection, err error) {
	sp := &Computer{g1, g2}
	return sp.SharedPaths()
}

// SharedPaths get SharedPaths returns forwDir,backDir.
func (s *Computer) SharedPaths() (forwDir matrix.Collection, backDir matrix.Collection, err error) {
	if !s.checkLinealInput() {
		return nil, nil, algoerr.ErrNotMatchType
	}
	r1, r2 := s.findLinearIntersections()
	for _, v := range r1 {
		forwDir = append(forwDir, v)
	}
	for _, v := range r2 {
		backDir = append(backDir, v)
	}
	return

}

func (s *Computer) findLinearIntersections() (forwDir []matrix.LineMatrix, backDir []matrix.LineMatrix) {
	if ml, ok := s.g1.(matrix.Collection); ok {
		for _, v := range ml {
			r1, r2 := findLinearIntersections(v.(matrix.LineMatrix), s.g2)
			forwDir = append(forwDir, r1...)
			backDir = append(backDir, r2...)
		}
	} else {
		r1, r2 := findLinearIntersections(s.g1.(matrix.LineMatrix), s.g2)
		forwDir = append(forwDir, r1...)
		backDir = append(backDir, r2...)
	}
	return
}

func findLinearIntersections(g1 matrix.LineMatrix, g2 matrix.Steric) (forwDir []matrix.LineMatrix, backDir []matrix.LineMatrix) {
	if ml, ok := g2.(matrix.Collection); ok {
		for _, v := range ml {
			ils := overlay.IntersectLine(g1, v.(matrix.LineMatrix))
			for _, il := range ils {
				if !il.Ips[1].IsEntering {
					forwDir = append(forwDir, matrix.LineMatrix{il.Ips[0].Matrix, il.Ips[1].Matrix})
				} else {
					backDir = append(backDir, matrix.LineMatrix{il.Ips[0].Matrix, il.Ips[1].Matrix})
				}
			}
		}
	} else {
		ils := overlay.IntersectLine(g1, g2.(matrix.LineMatrix))
		for _, il := range ils {
			if !il.Ips[1].IsEntering {
				forwDir = append(forwDir, matrix.LineMatrix{il.Ips[0].Matrix, il.Ips[1].Matrix})
			} else {
				backDir = append(backDir, matrix.LineMatrix{il.Ips[0].Matrix, il.Ips[1].Matrix})
			}
		}
	}
	return
}

func (s *Computer) checkLinealInput() bool {
	return checkLinealInput(s.g1) && checkLinealInput(s.g2)
}
func checkLinealInput(g matrix.Steric) bool {
	switch m := g.(type) {
	case matrix.LineMatrix:
		return true
	case matrix.Collection:
		for _, v := range m {
			if _, ok := v.(matrix.LineMatrix); !ok {
				return false
			}
		}
		return true
	}
	return false
}
