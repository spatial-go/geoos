package encode

import (
	"github.com/spatial-go/geoos/geoencoding/geobuf/protogeo"
	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

// Geometry ...
func Geometry(g *geojson.Geometry, cfg *EncodingConfig) *protogeo.Data_Geometry {
	switch g.Type {
	case space.TypePoint:
		p := g.Coordinates.(space.Point)
		return &protogeo.Data_Geometry{
			Type:   protogeo.Data_Geometry_POINT,
			Coords: translateCoords(cfg.Precision, p[:]),
		}
	case space.TypeLineString:
		p := g.Coordinates.(space.LineString)
		return &protogeo.Data_Geometry{
			Type:   protogeo.Data_Geometry_LINESTRING,
			Coords: translateLine(cfg.Precision, cfg.Dimension, p, false),
		}
	case space.TypeMultiLineString:
		p := g.Coordinates.(space.MultiLineString)
		coords, lengths := translateMultiLine(cfg.Precision, cfg.Dimension, p)
		return &protogeo.Data_Geometry{
			Type:    protogeo.Data_Geometry_MULTILINESTRING,
			Coords:  coords,
			Lengths: lengths,
		}
	case space.TypePolygon:
		p := g.Coordinates.(space.Polygon)
		coords, lengths := translateMultiRing(cfg.Precision, cfg.Dimension, p)
		return &protogeo.Data_Geometry{
			Type:    protogeo.Data_Geometry_POLYGON,
			Coords:  coords,
			Lengths: lengths,
		}
	case space.TypeMultiPolygon:
		p := []space.Polygon(g.Coordinates.(space.MultiPolygon))
		coords, lengths := translateMultiPolygon(cfg.Precision, cfg.Dimension, p)
		return &protogeo.Data_Geometry{
			Type:    protogeo.Data_Geometry_MULTIPOLYGON,
			Coords:  coords,
			Lengths: lengths,
		}
	}
	return nil
}

func translateMultiLine(e uint, dim uint, lines []space.LineString) ([]int64, []uint32) {
	lengths := make([]uint32, len(lines))
	coords := []int64{}

	for i, line := range lines {
		lengths[i] = uint32(len(line))
		coords = append(coords, translateLine(e, dim, line, false)...)
	}
	return coords, lengths
}

func translateMultiPolygon(e uint, dim uint, polygons []space.Polygon) ([]int64, []uint32) {
	lengths := []uint32{uint32(len(polygons))}
	coords := []int64{}
	for _, rings := range polygons {
		lengths = append(lengths, uint32(len(rings)))
		newLine, newLength := translateMultiRing(e, dim, rings)
		lengths = append(lengths, newLength...)
		coords = append(coords, newLine...)
	}
	return coords, lengths
}

func translateMultiRing(e uint, dim uint, lines space.Polygon) ([]int64, []uint32) {
	lengths := make([]uint32, len(lines))
	coords := []int64{}
	for i, line := range lines {
		lengths[i] = uint32(len(line) - 1)
		newLine := translateLine(e, dim, line, true)
		coords = append(coords, newLine...)
	}
	return coords, lengths
}

func translateLine(precision uint, dim uint, points space.LineString, isClosed bool) []int64 {
	sums := make([]int64, dim)
	ret := make([]int64, len(points)*int(dim))
	for i, point := range points {
		for j, p := range point {
			n := protogeo.IntWithPrecision(p, precision) - sums[j]
			ret[(int(dim)*i)+j] = n
			sums[j] = sums[j] + n
		}
	}
	if isClosed {
		return ret[:(len(ret) - int(dim))]
	}
	return ret
}

// Converts a floating point geojson point to int64 by multiplying it by a factor of 10,
// potentially truncating and rounding
func translateCoords(precision uint, point []float64) []int64 {
	ret := make([]int64, len(point))
	for i, p := range point {
		ret[i] = protogeo.IntWithPrecision(p, precision)
	}
	return ret
}
