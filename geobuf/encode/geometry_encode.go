package encode

import (
	"github.com/spatial-go/geoos/geobuf/proto"
	math "github.com/spatial-go/geoos/geobuf/utils"

	"github.com/spatial-go/geoos/geojson"
	geoos "github.com/spatial-go/geoos/space"
)

// Geometry ...
func Geometry(g *geojson.Geometry, cfg *EncodingConfig) *proto.Data_Geometry {
	switch g.Type {
	case geoos.TypePoint:
		p := g.Coordinates.(geoos.Point)
		return &proto.Data_Geometry{
			Type:   proto.Data_Geometry_POINT,
			Coords: translateCoords(cfg.Precision, p[:]),
		}
	case geoos.TypeLineString:
		p := g.Coordinates.(geoos.LineString)
		return &proto.Data_Geometry{
			Type:   proto.Data_Geometry_LINESTRING,
			Coords: translateLine(cfg.Precision, cfg.Dimension, p, false),
		}
	case geoos.TypeMultiLineString:
		p := g.Coordinates.(geoos.MultiLineString)
		coords, lengths := translateMultiLine(cfg.Precision, cfg.Dimension, p)
		return &proto.Data_Geometry{
			Type:    proto.Data_Geometry_MULTILINESTRING,
			Coords:  coords,
			Lengths: lengths,
		}
	case geoos.TypePolygon:
		p := g.Coordinates.(geoos.Polygon)
		coords, lengths := translateMultiRing(cfg.Precision, cfg.Dimension, p)
		return &proto.Data_Geometry{
			Type:    proto.Data_Geometry_POLYGON,
			Coords:  coords,
			Lengths: lengths,
		}
	case geoos.TypeMultiPolygon:
		p := []geoos.Polygon(g.Coordinates.(geoos.MultiPolygon))
		coords, lengths := translateMultiPolygon(cfg.Precision, cfg.Dimension, p)
		return &proto.Data_Geometry{
			Type:    proto.Data_Geometry_MULTIPOLYGON,
			Coords:  coords,
			Lengths: lengths,
		}
	}
	return nil
}

func translateMultiLine(e uint, dim uint, lines []geoos.LineString) ([]int64, []uint32) {
	lengths := make([]uint32, len(lines))
	coords := []int64{}

	for i, line := range lines {
		lengths[i] = uint32(len(line))
		coords = append(coords, translateLine(e, dim, line, false)...)
	}
	return coords, lengths
}

func translateMultiPolygon(e uint, dim uint, polygons []geoos.Polygon) ([]int64, []uint32) {
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

func translateMultiRing(e uint, dim uint, lines geoos.Polygon) ([]int64, []uint32) {
	lengths := make([]uint32, len(lines))
	coords := []int64{}
	for i, line := range lines {
		lengths[i] = uint32(len(line) - 1)
		newLine := translateLine(e, dim, line, true)
		coords = append(coords, newLine...)
	}
	return coords, lengths
}

func translateLine(precision uint, dim uint, points geoos.LineString, isClosed bool) []int64 {
	sums := make([]int64, dim)
	ret := make([]int64, len(points)*int(dim))
	for i, point := range points {
		for j, p := range point {
			n := math.IntWithPrecision(p, precision) - sums[j]
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
		ret[i] = math.IntWithPrecision(p, precision)
	}
	return ret
}
