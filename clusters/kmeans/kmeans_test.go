package kmeans

import (
	"math/rand"
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/clusters"
)

const (
	randomSeed = int64(42)
)

func TestNewErrors(t *testing.T) {
	_, err := NewWithOptions(0.00)
	if err == nil {
		t.Errorf("Expected invalid options to return an error, got nil")
	}

	_, err = NewWithOptions(1.00)
	if err == nil {
		t.Errorf("Expected invalid options to return an error, got nil")
	}
}

func TestPartitioningError(t *testing.T) {
	km := New()
	d := clusters.Points{}
	if _, err := km.Partition(d, 1); err == nil {
		t.Errorf("Expected error partitioning with empty data set, got nil")
		return
	}

	d = clusters.Points{
		geoos.Point{
			0.1,
			0.1,
		},
	}
	if _, err := km.Partition(d, 0); err == nil {
		t.Errorf("Expected error partitioning with 0 clusters, got nil")
		return
	}

	if _, err := km.Partition(d, 2); err == nil {
		t.Errorf("Expected error partitioning with more clusters than data points, got nil")
		return
	}
}

func TestDimensions(t *testing.T) {
	var d clusters.Points
	for x := 0; x < 255; x += 32 {
		for y := 0; y < 255; y += 32 {
			d = append(d, geoos.Point{
				float64(x) / 255.0,
				float64(y) / 255.0,
			})
		}
	}

	k := 4
	km := New()
	clusters, err := km.Partition(d, k)
	if err != nil {
		t.Errorf("Unexpected error partitioning: %v", err)
		return
	}

	if len(clusters) != k {
		t.Errorf("Expected %d clusters, got: %d", k, len(clusters))
	}
}

func benchmarkPartition(size, partitions int, b *testing.B) {
	rand.Seed(randomSeed)
	var d clusters.Points

	for i := 0; i < size; i++ {
		d = append(d, geoos.Point{
			rand.Float64(),
			rand.Float64(),
		})
	}

	for j := 0; j < b.N; j++ {
		km := New()
		km.Partition(d, partitions)
	}
}

func BenchmarkPartition32Points(b *testing.B)    { benchmarkPartition(32, 16, b) }
func BenchmarkPartition512Points(b *testing.B)   { benchmarkPartition(512, 16, b) }
func BenchmarkPartition4096Points(b *testing.B)  { benchmarkPartition(4096, 16, b) }
func BenchmarkPartition65536Points(b *testing.B) { benchmarkPartition(65536, 16, b) }
