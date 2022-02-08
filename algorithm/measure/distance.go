package measure

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
	"github.com/spatial-go/geoos/coordtransform"
)

const (

	// R radius of earth.
	R = 6371000.0 //6378137.0
	// E is eccentricity.
	E = 0.006694379990141317
)

// Distance is a func of measure distance.
type Distance func(from, to matrix.Steric) float64

// MercatorDistance scale factor is changed along the meridians as a function of latitude
// https://gis.stackexchange.com/questions/110730/mercator-scale-factor-is-changed-along-the-meridians-as-a-function-of-latitude
// https://gis.stackexchange.com/questions/93332/calculating-distance-scale-factor-by-latitude-for-mercator
func MercatorDistance(distance float64, lat float64) float64 {
	k, c := MercatorFctor(lat)
	distance = distance * k * c
	return distance
}

// MercatorFctor scale factor the meridians as a function of latitude
func MercatorFctor(lat float64) (float64, float64) {
	lat = lat * math.Pi / 180
	k := (1 / math.Cos(lat))
	c := math.Sqrt(1 - math.Pow(E, 2)*math.Pow(math.Sin(lat), 2))
	return k, c
}

// SpheroidDistance Calculate distance, return unit: meter
func SpheroidDistance(fromSteric, toSteric matrix.Steric) float64 {
	if to, ok := toSteric.(matrix.Matrix); ok {
		if from, ok := fromSteric.(matrix.Matrix); ok {
			rad := math.Pi / 180.0
			lat0 := from[1] * rad
			lng0 := from[0] * rad
			lat1 := to[1] * rad
			lng1 := to[0] * rad
			theta := lng1 - lng0
			dist := math.Acos(math.Sin(lat0)*math.Sin(lat1) + math.Cos(lat0)*math.Cos(lat1)*math.Cos(theta))
			return dist * R
		}
	}
	trans := coordtransform.NewTransformer(coordtransform.LLTOMERCATOR)
	from, _ := trans.TransformGeometry(fromSteric)
	to, _ := trans.TransformGeometry(toSteric)
	locMatrix := []matrix.Matrix{{0, 0}, {0, 0}}
	dist := distanceCompute(from, to, locMatrix)

	trans = coordtransform.NewTransformer(coordtransform.MERCATORTOLL)
	loc0, _ := trans.TransformGeometry(locMatrix[0])
	loc1, _ := trans.TransformGeometry(locMatrix[1])
	k, _ := MercatorFctor((loc0.(matrix.Matrix)[1] + loc1.(matrix.Matrix)[1]) / 2)
	return dist / k
}

// PlanarDistance returns Distance of form to.
func PlanarDistance(fromSteric, toSteric matrix.Steric) float64 {
	locMatrix := []matrix.Matrix{{0, 0}, {0, 0}}
	return distanceCompute(fromSteric, toSteric, locMatrix)
}

// distanceCompute returns Distance of form to.
func distanceCompute(fromSteric, toSteric matrix.Steric, locMatrix []matrix.Matrix) float64 {
	switch to := toSteric.(type) {
	case matrix.Matrix:
		if from, ok := fromSteric.(matrix.Matrix); ok {
			locMatrix[0], locMatrix[1] = from, to
			return math.Sqrt((from[0]-to[0])*(from[0]-to[0]) + (from[1]-to[1])*(from[1]-to[1]))
		}
		return distanceCompute(to, fromSteric, locMatrix)

	case matrix.LineMatrix:
		if from, ok := fromSteric.(matrix.Matrix); ok {
			return distanceLineToPoint(to, from, locMatrix)
		} else if from, ok := fromSteric.(matrix.LineMatrix); ok {
			return distanceLineAndLine(from, to, locMatrix)
		}
		return distanceCompute(to, fromSteric, locMatrix)
	case matrix.PolygonMatrix:
		if from, ok := fromSteric.(matrix.Matrix); ok {
			return distancePolygonToPoint(to, from, locMatrix)
		} else if from, ok := fromSteric.(matrix.LineMatrix); ok {
			return distancePolygonAndLine(to, from, locMatrix)
		} else if from, ok := fromSteric.(matrix.PolygonMatrix); ok {
			dist := math.MaxFloat64
			loc := []matrix.Matrix{{0, 0}, {0, 0}}
			for _, v := range from {
				if distP := distanceCompute(matrix.LineMatrix(v), to, loc); dist > distP {
					locMatrix[0], locMatrix[1] = loc[0], loc[1]
					dist = distP
				}
			}
			return dist
		}
		return PlanarDistance(to, fromSteric)
	case matrix.Collection:
		dist := math.MaxFloat64
		loc := []matrix.Matrix{{0, 0}, {0, 0}}
		for _, v := range to {
			if distP := distanceCompute(fromSteric, v, loc); dist > distP {
				locMatrix[0], locMatrix[1] = loc[0], loc[1]
				dist = distP
			}
		}
		return dist
	default:
		return 0
	}
}

// distanceSegmentToPoint Returns Distance of p,ab
func distanceSegmentToPoint(p, a, b matrix.Matrix) float64 {
	// if start = end, then just compute distance to one of the endpoints
	if a[0] == b[0] && a[1] == b[1] {
		return PlanarDistance(p, a)
	}
	// otherwise use comp.graphics.algorithms Frequently Asked Questions method
	//
	// (1) r = AC dot AB
	//         ---------
	//         ||AB||^2
	//
	// r has the following meaning:
	//   r=0 P = A
	//   r=1 P = B
	//   r<0 P is on the backward extension of AB
	//   r>1 P is on the forward extension of AB
	//   0<r<1 P is interior to AB

	len2 := (b[0]-a[0])*(b[0]-a[0]) + (b[1]-a[1])*(b[1]-a[1])
	r := ((p[0]-a[0])*(b[0]-a[0]) + (p[1]-a[1])*(b[1]-a[1])) / len2

	if r <= 0.0 {
		return PlanarDistance(p, a)
	}
	if r >= 1.0 {
		return PlanarDistance(p, b)
	}

	//
	// (2) s = (Ay-Cy)(Bx-Ax)-(Ax-Cx)(By-Ay)
	//         -----------------------------
	//                    L^2
	//
	// Then the distance from C to P = |s|*L.
	//
	// This is the same calculation .
	// Unrolled here for performance.
	//
	s := ((a[1]-p[1])*(b[0]-a[0]) - (a[0]-p[0])*(b[1]-a[1])) / len2
	return math.Abs(s) * math.Sqrt(len2)
}

// distanceLineToPoint Returns Distance of p,line
func distanceLineToPoint(line matrix.LineMatrix, pt matrix.Matrix, locMatrix []matrix.Matrix) (dist float64) {
	dist = math.MaxFloat64
	for i, v := range line {
		if i < len(line)-1 {
			if tmpDist := distanceSegmentToPoint(pt, v, line[i+1]); dist > tmpDist {
				locMatrix[0] = pt
				if PlanarDistance(pt, matrix.Matrix(v)) < PlanarDistance(pt, matrix.Matrix(line[i+1])) {
					locMatrix[1] = v
				} else {
					locMatrix[1] = line[i+1]
				}
				dist = tmpDist
			}
		}
	}
	return
}

// distancePolygonToPoint Returns Distance of p,polygon
func distancePolygonToPoint(poly matrix.PolygonMatrix, pt matrix.Matrix, locMatrix []matrix.Matrix) (dist float64) {
	dist = math.MaxFloat64
	loc := []matrix.Matrix{{0, 0}, {0, 0}}
	for _, v := range poly {
		tmpDist := distanceLineToPoint(v, pt, loc)
		if dist > tmpDist {
			locMatrix[0], locMatrix[1] = loc[0], loc[1]
			dist = tmpDist
		}
	}
	return
}

// distanceLineAndLine returns distance Between the two Geometry.
func distanceLineAndLine(from, to matrix.LineMatrix, locMatrix []matrix.Matrix) (dist float64) {
	dist = math.MaxFloat64
	loc := []matrix.Matrix{{0, 0}, {0, 0}}
	if mark := relate.IsIntersectionEdge(from, to); mark {
		return 0
	}
	for _, v := range from {
		if distP := distanceLineToPoint(to, matrix.Matrix(v), loc); dist > distP {
			locMatrix[0], locMatrix[1] = loc[0], loc[1]
			dist = distP
		}
	}
	for _, v := range to {
		if distP := distanceLineToPoint(from, matrix.Matrix(v), loc); dist > distP {
			locMatrix[0], locMatrix[1] = loc[0], loc[1]
			dist = distP
		}
	}
	return dist
}

// distancePolygonAndLine returns distance Between the two Geometry.
func distancePolygonAndLine(poly matrix.PolygonMatrix, line matrix.LineMatrix, locMatrix []matrix.Matrix) (dist float64) {
	dist = math.MaxFloat64
	loc := []matrix.Matrix{{0, 0}, {0, 0}}
	for _, v := range poly {
		if distP := distanceLineAndLine(matrix.LineMatrix(v), line, loc); dist > distP {
			locMatrix[0], locMatrix[1] = loc[0], loc[1]
			dist = distP
		}
	}
	return dist
}
