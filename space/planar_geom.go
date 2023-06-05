package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/coordtransform"
	"github.com/spatial-go/geoos/space/spaceerr"
)

type PlanarGeom[T Geometry] struct {
	geom T
}

func (p *PlanarGeom[T]) Dimensions() int {
	return p.geom.Dimensions()
}

// Centroid Computes the centroid point of a geometry.
func (p *PlanarGeom[T]) Centroid() Point {
	if p.geom.IsEmpty() {
		return nil
	}
	cent := &buffer.CentroidComputer{}
	cent.Add(p.geom.ToMatrix())
	m := cent.GetCentroid()
	return Point(m)
}

// Distance returns distance Between the two Geometry.
func (p *PlanarGeom[T]) Distance(to Geometry, f measure.DistanceFunc) (float64, error) {
	if p.geom.IsEmpty() ||
		to == nil || to.IsEmpty() {
		return 0, nil
	}
	if p.geom.IsEmpty() != to.IsEmpty() {
		return 0, spaceerr.ErrNilGeometry
	}
	return f(p.geom.ToMatrix(), to.ToMatrix()), nil
}

// bufferInMeter ...
func (p *PlanarGeom[T]) bufferInMeter(width float64, quadsegs int) Geometry {

	switch p.geom.GeoJSONType() {
	case TypeLineString:
		ls := p.geom.Geom().(LineString)
		distances := make([]float64, len(ls))

		for i := range distances {
			distances[i] = measure.MercatorDistance(width, Point(ls[i]).Lat())
		}

		transformer := coordtransform.NewTransformer(coordtransform.LLTOMERCATOR)
		geomMatrix, _ := transformer.TransformGeometry(ls.ToMatrix())

		lbuffer := &buffer.VariableLineBuffer{Line: geomMatrix.(matrix.LineMatrix), QuadrantSegments: quadsegs}
		resultMatrix := lbuffer.DistancesBuffer(distances)
		geometry := TransGeometry(resultMatrix)
		if geometry != nil {
			transformer.CoordType = coordtransform.MERCATORTOLL
			geomMatrix, _ = transformer.TransformGeometry(geometry.ToMatrix())
			geometry = TransGeometry(geomMatrix)
		}
		return geometry
	default:
		centroid := p.geom.Centroid()
		width = measure.MercatorDistance(width, centroid.Lat())
		transformer := coordtransform.NewTransformer(coordtransform.LLTOMERCATOR)
		geomMatrix, _ := transformer.TransformGeometry(p.geom.ToMatrix())
		geometry := TransGeometry(geomMatrix)
		geometry = geometry.Buffer(width, quadsegs)
		if geometry != nil {
			transformer.CoordType = coordtransform.MERCATORTOLL
			geomMatrix, _ = transformer.TransformGeometry(geometry.ToMatrix())
			geometry = TransGeometry(geomMatrix)
		}
		return geometry
	}
}

// bufferInOriginal ...
func (p *PlanarGeom[T]) bufferInOriginal(width float64, quadsegs int) Geometry {
	buff := buffer.Buffer(p.geom.ToMatrix(), width, quadsegs)

	result := buff
	switch b := result.(type) {
	case matrix.LineMatrix:
		return LineString(b)
	case matrix.PolygonMatrix:
		return Polygon(b)
	}
	return nil
}
