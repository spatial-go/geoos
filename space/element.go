package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/coordtransform"
	"github.com/spatial-go/geoos/space/spaceerr"
)

// const Coordinate System
const (
	BJ54 = iota + 1000000
	XA80
	CGCS2000

	// WGS84 World Geodetic System一1984 Coordinate System
	WGS84 = 4326

	// PseudoMercator  WGS 84 / Pseudo-Mercator
	PseudoMercator = 3857

	//GCJ02 Guojia cehui ju 02 ,unit degree
	GCJ02 = 104326

	//GCJ02Web Guojia cehui ju 02 Mercator, unit m
	GCJ02Web = 103857

	// BD09 Guojia cehui ju 02+BD ,unit degree
	BD09 = 114326

	// BD09Web Guojia cehui ju 02+BD, unit m
	BD09Web = 113857
)

var projectionCoordinateSystem = []int{PseudoMercator, GCJ02Web, BD09Web}

// Line  straight line  .
type Line struct {
	Start, End Point
}

// GeometryValid describes a geographic Element Valid
type GeometryValid struct {
	Geometry
	coordinateSystem int
}

// CreateElementValid Returns valid geom element. returns nil if geom is invalid.
func CreateElementValid(geom Geometry) (*GeometryValid, error) {
	return CreateElementValidWithCoordSys(geom, GCJ02)
}

// CreateElementValidWithCoordSys Returns valid geom element. returns nil if geom is invalid.
func CreateElementValidWithCoordSys(geom Geometry, coordSys int) (*GeometryValid, error) {
	geom = geom.Filter(&matrix.UniqueArrayFilter{})
	if geom.IsValid() {
		return &GeometryValid{geom, coordSys}, nil
	}
	return nil, spaceerr.ErrNotValidGeometry
}

// CoordinateSystem return Coordinate System.
func (g *GeometryValid) CoordinateSystem() int {
	return g.coordinateSystem
}

// IsProjection returns true if the coordinateSystem is projection.
func (g *GeometryValid) IsProjection() bool {
	for i := range projectionCoordinateSystem {
		if projectionCoordinateSystem[i] == g.coordinateSystem {
			return true
		}
	}
	return false
}

// Geom return Geometry without Coordinate System.
func (g *GeometryValid) Geom() Geometry {
	return g.Geometry
}

func defaultCoordinateSystem() int {
	return GCJ02
}

// Centroid Computes the centroid point of a geometry.
func Centroid(geom Geometry) Point {
	cent := &buffer.CentroidComputer{}

	if geom == nil || geom.IsEmpty() {
		return nil
	}
	cent.Add(geom.ToMatrix())
	m := cent.GetCentroid()
	return Point(m)
}

// Distance returns distance Between the two Geometry.
func Distance(from, to Geometry, f measure.Distance) (float64, error) {
	if from == nil || from.IsEmpty() ||
		to == nil || to.IsEmpty() {
		return 0, nil
	}
	if from.IsEmpty() != to.IsEmpty() {
		return 0, spaceerr.ErrNilGeometry
	}
	return f(from.ToMatrix(), to.ToMatrix()), nil
}

// TransGeometry trans steric to geometry.
func TransGeometry(inputGeom matrix.Steric) Geometry {
	switch g := inputGeom.(type) {
	case matrix.Matrix:
		return Point(g)
	case matrix.LineMatrix:
		if len(g) == 1 {
			return Point(matrix.Matrix(g[0]))
		}
		return LineString(g)
	case matrix.PolygonMatrix:
		return Polygon(g)
	case matrix.Collection:
		multiType := ""
		for _, v := range g {
			if multiType == "" {
				multiType = TransGeometry(v).GeoJSONType()
			}
			if multiType != TransGeometry(v).GeoJSONType() {
				multiType = ""
				break
			}
		}
		switch multiType {
		case TypeLineString:
			var coll MultiLineString
			for _, v := range g {
				coll = append(coll, TransGeometry(v).(LineString))
			}
			return coll
		case TypePoint:
			var coll MultiPoint
			for _, v := range g {
				coll = append(coll, TransGeometry(v).(Point))
			}
			return coll
		case TypePolygon:
			var coll MultiPolygon
			for _, v := range g {
				coll = append(coll, TransGeometry(v).(Polygon))
			}
			return coll
		default:
			var coll Collection
			for _, v := range g {
				coll = append(coll, TransGeometry(v))
			}
			return coll
		}
	default:
		return nil
	}
}

// BufferInMeter ...
func BufferInMeter(geometry Geometry, width float64, quadsegs int) Geometry {
	centroid := geometry.Centroid()
	width = measure.MercatorDistance(width, centroid.Lat())
	transformer := coordtransform.NewTransformer(coordtransform.LLTOMERCATOR)
	geomMatrix, _ := transformer.TransformGeometry(geometry.ToMatrix())
	geometry = TransGeometry(geomMatrix)
	geometry = geometry.Buffer(width, quadsegs)
	if geometry != nil {
		transformer.CoordType = coordtransform.MERCATORTOLL
		geomMatrix, _ = transformer.TransformGeometry(geometry.ToMatrix())
		geometry = TransGeometry(geomMatrix)
	}
	return geometry
}
