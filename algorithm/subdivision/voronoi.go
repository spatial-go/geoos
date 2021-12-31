package subdivision

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/algorithm/overlay"
	"github.com/spatial-go/geoos/algorithm/subdivision/quadedge"
)

const DefaultTolerance = 0.000000001

// Voronoi ...
type Voronoi struct {
	sites       []matrix.Matrix
	envelope    *envelope.Envelope
	subdivision *quadedge.QuadEdgeSubdivision
	result      []matrix.PolygonMatrix
}

// NewVoronoi return a new voronoi
func NewVoronoi() *Voronoi {
	return &Voronoi{}
}

// AddSites add some sites to voronoi
func (v *Voronoi) AddSites(sites []matrix.Matrix) {
	// TODO: should remove duplicate sites
	v.sites = append(v.sites, sites...)
	v.clearResult()
}

// ClearSites clear all sites of voronoi
func (v *Voronoi) ClearSites(sites []matrix.Matrix) {
	v.sites = v.sites[:0]
	v.clearResult()
}

// GetSites return sites of voronoi
func (v *Voronoi) GetSites() []matrix.Matrix {
	return v.sites
}

// SetEnvelope set envelope of voronoi
func (v *Voronoi) SetEnvelope(env envelope.Envelope) {
	v.envelope = &env
	v.clearResult()
}

// GetEnvelope return envelope of voronoi
func (v *Voronoi) GetEnvelope() envelope.Envelope {
	return *v.envelope
}

// clearResult clear result of voronoi
func (v *Voronoi) clearResult() {
	v.result = nil
}

// GetResult return result of voronoi
func (v *Voronoi) GetResult() []matrix.PolygonMatrix {
	if v.result != nil {
		return v.result
	}
	if len(v.sites) == 0 {
		return v.result
	}
	if v.envelope.IsNil() {
		v.envelope = envelope.Empty()
		for _, site := range v.sites {
			v.envelope.ExpandToIncludeMatrix(site)
		}
		if v.envelope.IsNil() {
			return v.result
		}
	}
	v.subdivision = quadedge.NewQuadEdgeSubdivision(v.envelope, 0.0)
	triangulator := NewIncrementalDelaunayTriangulator(v.subdivision)
	triangulator.insertSites(v.sites)

	polygons := v.subdivision.GetVoronoiCellPolygons()

	v.result = clipPolygons(polygons, v.envelope)
	return v.result
}

func clipPolygons(polygons []matrix.PolygonMatrix, env *envelope.Envelope) (clippedPolygons []matrix.PolygonMatrix) {
	if env.IsNil() {
		return
	}
	if len(polygons) == 0 {
		return
	}
	clippedPolygons = make([]matrix.PolygonMatrix, 0, len(polygons))
	clipEnv := env.ToMatrix()
	for _, polygon := range polygons {
		polygonEnv := envelope.Bound(polygon.Bound())
		if env.Contains(polygonEnv) {
			clippedPolygons = append(clippedPolygons, polygon)
		} else if env.IsIntersects(polygonEnv) {
			clipPolygon, _ := overlay.Intersection(*clipEnv, polygon)
			if clipPolygon != nil {
				clippedPolygons = append(clippedPolygons, clipPolygon.(matrix.PolygonMatrix))
			}
		}
	}
	return
}
