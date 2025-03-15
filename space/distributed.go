package space

import (
    "sync"
)

// WorkerPool manages a pool of workers for distributed geometry processing.
type WorkerPool struct {
    workers int
    tasks   chan func() Geometry
    results chan Geometry
    wg      sync.WaitGroup
}

// NewWorkerPool initializes a worker pool with a specified number of workers.
func NewWorkerPool(workers int) *WorkerPool {
    pool := &WorkerPool{
        workers: workers,
        tasks:   make(chan func() Geometry),
        results: make(chan Geometry, workers),
    }
    pool.start()
    return pool
}

// start launches the worker goroutines.
func (p *WorkerPool) start() {
    for i := 0; i < p.workers; i++ {
        p.wg.Add(1)
        go func() {
            defer p.wg.Done()
            for task := range p.tasks {
                result := task()
                p.results <- result
            }
        }()
    }
}

// AddTask adds a task to the worker pool.
func (p *WorkerPool) AddTask(task func() Geometry) {
    p.tasks <- task
}

// Wait waits for all tasks to complete and returns merged results as a Geometry.
func (p *WorkerPool) Wait() Geometry {
    close(p.tasks)
    p.wg.Wait()
    close(p.results)

    var results []Geometry
    for result := range p.results {
        results = append(results, result)
    }
    return mergeGeometries(results)
}

// mergeGeometries combines multiple Geometry results into a single Geometry.
func mergeGeometries(geometries []Geometry) Geometry {
    if len(geometries) == 0 {
        return nil
    }
    if len(geometries) == 1 {
        return geometries[0]
    }
    // For simplicity, assume results are Polygons or Rings and create a Collection
    collection := Collection{}
    for _, g := range geometries {
        collection = append(collection, g)
    }
    return collection
}