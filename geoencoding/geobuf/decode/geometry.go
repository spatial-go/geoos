package decode

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/geoencoding/geobuf/protogeo"
	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

// Geometry ...
func Geometry(geo *protogeo.Data_Geometry, precision, dimensions uint32) *geojson.Geometry {
	switch geo.Type {
	case protogeo.Data_Geometry_POINT:
		return geojson.NewGeometry(makePoint(geo.Coords, precision))
	case protogeo.Data_Geometry_MULTIPOINT:
		return geojson.NewGeometry(makeMultiPoint(geo.Coords, precision, dimensions))
	case protogeo.Data_Geometry_LINESTRING:
		return geojson.NewGeometry(makeLineString(geo.Coords, precision, dimensions))
	case protogeo.Data_Geometry_MULTILINESTRING:
		return geojson.NewGeometry(makeMultiLineString(geo.Lengths, geo.Coords, precision, dimensions))
	case protogeo.Data_Geometry_POLYGON:
		return geojson.NewGeometry(makePolygon(geo.Lengths, geo.Coords, precision, dimensions))
	case protogeo.Data_Geometry_MULTIPOLYGON:
		return geojson.NewGeometry(makeMultiPolygon(geo.Lengths, geo.Coords, precision, dimensions))
	}
	return &geojson.Geometry{}
}

func makePoint(inCords []int64, precision uint32) space.Point {
	return space.Point{makeCoords(inCords, precision)[0], makeCoords(inCords, precision)[1]}
}

func makeMultiPoint(inCords []int64, precision uint32, dimension uint32) space.MultiPoint {
	points := make(space.MultiPoint, len(inCords)/int(dimension))
	prevCords := [2]int64{}
	for i, j := 0, 1; j < len(inCords); i, j = i+2, j+2 {
		prevCords[0] += inCords[i]
		prevCords[1] += inCords[j]
		points[i/2] = makePoint(prevCords[:], precision)
	}
	return points
}

func makeMultiPolygon(lengths []uint32, inCords []int64, precision uint32, dimension uint32) space.MultiPolygon {
	polyCount := int(lengths[0])
	polygons := make([]space.Polygon, polyCount)
	lengths = lengths[1:]
	for i := 0; i < polyCount; i++ {
		ringCount := lengths[0]
		polygons[i] = makePolygon(lengths[1:ringCount+1], inCords, precision, dimension)
		skip := 0
		for i := 0; i < int(ringCount); i++ {
			skip += int(lengths[i]) * int(dimension)
		}

		lengths = lengths[ringCount:]
		inCords = inCords[skip:]
	}
	return polygons
}

func makePolygon(lengths []uint32, inCords []int64, precision uint32, dimension uint32) space.Polygon {
	lines := make(matrix.PolygonMatrix, len(lengths))
	for i, length := range lengths {
		l := int(length * dimension)
		lines[i] = makeRing(inCords[:l], precision, dimension)
		inCords = inCords[l:]
	}
	poly := space.Polygon(lines)
	return poly
}

func makeMultiLineString(lengths []uint32, inCords []int64, precision uint32, dimension uint32) space.MultiLineString {
	lines := make([]space.LineString, len(lengths))
	for i, length := range lengths {
		l := int(length * dimension)
		lines[i] = makeLineString(inCords[:l], precision, dimension)
		inCords = inCords[l:]
	}
	return lines
}

func makeRing(inCords []int64, precision uint32, dimension uint32) space.Ring {
	points := makeLine(inCords, precision, dimension)
	points = append(points, points[0])
	return points
}

func makeLineString(inCords []int64, precision uint32, dimension uint32) space.LineString {
	points := make(space.Ring, len(inCords)/int(dimension))
	prevCords := [2]int64{}
	for i, j := 0, 1; j < len(inCords); i, j = i+2, j+2 {
		prevCords[0] += inCords[i]
		prevCords[1] += inCords[j]
		points[i/2] = makePoint(prevCords[:], precision)
	}
	return space.LineString(makeLine(inCords, precision, dimension))
}

func makeLine(inCords []int64, precision uint32, dimension uint32) space.Ring {
	points := make(space.Ring, len(inCords)/int(dimension))
	prevCords := [2]int64{}
	for i, j := 0, 1; j < len(inCords); i, j = i+2, j+2 {
		prevCords[0] += inCords[i]
		prevCords[1] += inCords[j]
		points[i/2] = makePoint(prevCords[:], precision)
	}
	return points
}

func makeCoords(inCords []int64, precision uint32) [2]float64 {
	ret := [2]float64{}
	e := protogeo.DecodePrecision(precision)
	for i, val := range inCords {
		ret[i] = protogeo.FloatWithPrecision(val, uint32(e))
	}
	return ret
}
