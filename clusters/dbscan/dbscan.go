package dbscan

import (
	"github.com/spatial-go/geoos/clusters"
	"github.com/spatial-go/geoos/common"
	"github.com/spatial-go/geoos/space"
)

// DBSCAN in pseudocode (from http://en.wikipedia.org/wiki/DBSCAN):

// EpsFunction is a function that returns eps based on point pt
type EpsFunction func(pt space.Point) float64

// DBScan clusters incoming points into clusters with params (eps, minPoints)
//
// eps is clustering radius in km
// minPoints in minimum number of points in eps-neighbourhood (density)
func DBScan(points clusters.PointList, eps float64, minPoints int) (clusterArry []clusters.Cluster, noise []int) {
	visited := make([]bool, len(points))
	members := make([]bool, len(points))
	clusterArry = []clusters.Cluster{}
	noise = []int{}
	C := 0
	kdTree := NewKDTree(points)

	// Our SphericalDistanceFast returns distance which is not mutiplied
	// by EarthR * DegreeRad, adjust eps accordingly
	eps = eps / common.EarthR / common.DegreeRad

	// neighborUnique := bitset.New(uint(len(points)))
	neighborUnique := make(map[int]int)

	for i := 0; i < len(points); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true

		neighborPts := kdTree.InRange(points[i], eps, nil)
		if len(neighborPts) < minPoints {
			noise = append(noise, i)
		} else {
			cluster := clusters.Cluster{C: C, Points: []int{i}}
			members[i] = true
			C++
			// expandCluster goes here inline
			neighborUnique = make(map[int]int)
			for j := 0; j < len(neighborPts); j++ {
				neighborUnique[neighborPts[j]] = neighborPts[j]
			}

			for j := 0; j < len(neighborPts); j++ {
				k := neighborPts[j]
				if !visited[k] {
					visited[k] = true
					moreNeighbors := kdTree.InRange(points[k], eps, nil)
					if len(moreNeighbors) >= minPoints {
						for _, p := range moreNeighbors {
							if _, ok := neighborUnique[p]; !ok {
								neighborPts = append(neighborPts, p)
								neighborUnique[p] = p
							}
						}
					}
				}

				if !members[k] {
					cluster.Points = append(cluster.Points, k)
					members[k] = true
				}
			}
			clusterArry = append(clusterArry, cluster)
		}
	}
	return
}

// RegionQuery is simple way O(N) to find points in neighbourhood
//
// It is roughly equivalent to kdTree.InRange(points[i], eps, nil)
func RegionQuery(points clusters.PointList, P space.Point, eps float64) []int {
	result := []int{}

	for i := 0; i < len(points); i++ {
		// if points[i].sqDist(P) < eps*eps {
		dis := DistanceSphericalFast(points[i], P)
		if dis < eps*eps {
			result = append(result, i)
		}
	}
	return result
}
