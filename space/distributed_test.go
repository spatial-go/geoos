package space

import (
    "testing"
    "github.com/spatial-go/geoos/algorithm/matrix"
)

func TestRingBufferInMeterDistributed(t *testing.T) {
    tests := []struct {
        name    string
        ring    Ring
        width   float64
        quadsegs int
        workers  int
        wantEmpty bool
    }{
        {
            name:    "Valid Ring with 2 Workers",
            ring:    Ring{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {0, 0}},
            width:   10,
            quadsegs: 8,
            workers:  2,
            wantEmpty: false,
        },
        {
            name:    "Single Point Ring (Fallback)",
            ring:    Ring{{0, 0}},
            width:   10,
            quadsegs: 8,
            workers:  2,
            wantEmpty: true, // Should fallback and handle empty/invalid gracefully
        },
        {
            name:    "One Worker (Fallback)",
            ring:    Ring{{0, 0}, {1, 0}, {2, 0}, {0, 0}},
            width:   10,
            quadsegs: 8,
            workers:  1,
            wantEmpty: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := tt.ring.BufferInMeterDistributed(tt.width, tt.quadsegs, tt.workers)
            if result == nil || result.IsEmpty() != tt.wantEmpty {
                t.Errorf("%s: expected empty=%v, got %v", tt.name, tt.wantEmpty, result)
            }
        })
    }
}

func TestPolygonBufferInMeterDistributed(t *testing.T) {
    tests := []struct {
        name    string
        poly    Polygon
        width   float64
        quadsegs int
        workers  int
        wantEmpty bool
    }{
        {
            name:    "Valid Polygon with 2 Workers",
            poly:    Polygon{{{0, 0}, {1, 0}, {1, 1}, {0, 1}, {0, 0}}},
            width:   10,
            quadsegs: 8,
            workers:  2,
            wantEmpty: false,
        },
        {
            name:    "Empty Polygon",
            poly:    Polygon{},
            width:   10,
            quadsegs: 8,
            workers:  2,
            wantEmpty: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := tt.poly.BufferInMeterDistributed(tt.width, tt.quadsegs, tt.workers)
            if result == nil || result.IsEmpty() != tt.wantEmpty {
                t.Errorf("%s: expected empty=%v, got %v", tt.name, tt.wantEmpty, result)
            }
        })
    }
}

func TestCollectionBufferInMeterDistributed(t *testing.T) {
    ring := Ring{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {0, 0}}
    poly := Polygon{{{0, 0}, {1, 0}, {1, 1}, {0, 1}, {0, 0}}}
    tests := []struct {
        name    string
        coll    Collection
        width   float64
        quadsegs int
        workers  int
        wantEmpty bool
    }{
        {
            name:    "Valid Collection with 2 Workers",
            coll:    Collection{ring, poly},
            width:   10,
            quadsegs: 8,
            workers:  2,
            wantEmpty: false,
        },
        {
            name:    "Empty Collection",
            coll:    Collection{},
            width:   10,
            quadsegs: 8,
            workers:  2,
            wantEmpty: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := tt.coll.BufferInMeterDistributed(tt.width, tt.quadsegs, tt.workers)
            if result == nil || result.IsEmpty() != tt.wantEmpty {
                t.Errorf("%s: expected empty=%v, got %v", tt.name, tt.wantEmpty, result)
            }
        })
    }
}

func BenchmarkBufferInMeterDistributed(b *testing.B) {
    ring := Ring(make(LineString, 1000))
    for i := 0; i < 1000; i++ {
        ring[i] = Point{float64(i), 0}
    }
    poly := Polygon{matrix.LineMatrix(ring)}
    coll := Collection{ring, poly}

    b.Run("RingSingle", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            ring.BufferInMeter(10, 8)
        }
    })
    b.Run("RingDistributed4", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            ring.BufferInMeterDistributed(10, 8, 4)
        }
    })
    b.Run("PolySingle", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            poly.BufferInMeter(10, 8)
        }
    })
    b.Run("PolyDistributed4", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            poly.BufferInMeterDistributed(10, 8, 4)
        }
    })
    b.Run("CollSingle", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            coll.BufferInMeter(10, 8)
        }
    })
    b.Run("CollDistributed4", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            coll.BufferInMeterDistributed(10, 8, 4)
        }
    })
}