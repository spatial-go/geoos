package subdivision

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/algorithm/subdivision/quadedge"
)

// DelaunayTriangulation ...
type DelaunayTriangulation struct {
	sites       []matrix.Matrix
	sitesEnv    *envelope.Envelope
	subdivision *quadedge.Subdivision
}

func (d *DelaunayTriangulation) computeEnvelope() {
	d.sitesEnv = envelope.Empty()
	for _, site := range d.sites {
		d.sitesEnv.ExpandToIncludeMatrix(site)
	}
}

func (d *DelaunayTriangulation) create() {
	if d.subdivision != nil {
		return
	}
	d.computeEnvelope()
	d.subdivision = quadedge.NewQuadEdgeSubdivision(d.sitesEnv, 0.0)
	triangulator := NewIncrementalDelaunayTriangulator(d.subdivision)
	triangulator.insertSites(d.sites)
}

// Subdivision ...
func (d *DelaunayTriangulation) Subdivision() *quadedge.Subdivision {
	d.create()
	return d.subdivision
}
