package clusters

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/planar"
)

// coordinates
// Observation
// geoos.Point

// Points is a slice of observations
// Observations is a slice of observations
// type Points []geoos.Point

// A Cluster which data points gravitate around
type Cluster struct {
	Center geoos.Point
	Points Points
}

// Clusters is a slice of clusters
type Clusters []Cluster

// New sets up a new set of clusters and randomly seeds their initial positions
func New(k int, dataset Points) (Clusters, error) {
	var c Clusters
	if len(dataset) == 0 || len(dataset[0]) == 0 {
		return c, fmt.Errorf("there must be at least one dimension in the data set")
	}
	if k == 0 {
		return c, fmt.Errorf("k must be greater than 0")
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < k; i++ {
		var p geoos.Point
		p[0] = rand.Float64()
		p[1] = rand.Float64()

		c = append(c, Cluster{
			Center: p,
		})
	}
	return c, nil
}

// Append adds an observation to the Cluster
func (c *Cluster) Append(point geoos.Point) {
	c.Points = append(c.Points, point)
}

// Nearest returns the index of the cluster nearest to point
func (c Clusters) Nearest(point geoos.Point) int {
	var ci int
	dist := -1.0
	G := planar.GEOAlgorithm{}

	// Find the nearest cluster for this data point
	for i, cluster := range c {
		d, _ := G.Distance(point, cluster.Center)
		if dist < 0 || d < dist {
			dist = d
			ci = i
		}
	}
	return ci
}

// Neighbour returns the neighbouring cluster of a point along with the average distance to its points
func (c Clusters) Neighbour(point geoos.Point, fromCluster int) (int, float64) {
	var d float64
	nc := -1

	for i, cluster := range c {
		if i == fromCluster {
			continue
		}

		cd := AverageDistance(point, cluster.Points)
		if nc < 0 || cd < d {
			nc = i
			d = cd
		}
	}
	return nc, d
}

// Recenter recenters a cluster
func (c *Cluster) Recenter() {
	center, err := c.Points.Center()
	if err != nil {
		return
	}
	c.Center = center
}

// Recenter recenters all clusters
func (c Clusters) Recenter() {
	for i := 0; i < len(c); i++ {
		c[i].Recenter()
	}
}

// Reset clears all point assignments
func (c Clusters) Reset() {
	for i := 0; i < len(c); i++ {
		c[i].Points = Points{}
	}
}

// PointsInDimension returns all coordinates in a given dimension
// TODO ?
func (c Cluster) PointsInDimension(n int) (v []float64) {
	for _, p := range c.Points {
		v = append(v, p[n])
	}
	return v
}

// CentersInDimension returns all cluster centroids' coordinates in a given
// dimension
// TODO ?
func (c Clusters) CentersInDimension(n int) (v []float64) {
	for _, cl := range c {
		v = append(v, cl.Center[n])
	}
	return v
}
