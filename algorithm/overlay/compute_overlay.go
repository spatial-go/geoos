package overlay

// ComputePolyOverlay overlay polygon.
type ComputePolyOverlay interface {
	Next(pol *Plane, start *Vertex) *Vertex
	Compute(pol *Plane, start *Vertex, which bool) *Vertex
}

// ComputeMergeOverlay merge overlay polygon.
type ComputeMergeOverlay struct {
	*PolygonOverlay
}

// Next overlay polygon.
func (c *ComputeMergeOverlay) Next(pol *Plane, start *Vertex) *Vertex {
	next := c.Compute(pol, start, true)
	next = c.Compute(pol, next, false)
	return next
}

// Compute overlay polygon.
func (c *ComputeMergeOverlay) Compute(pol *Plane, start *Vertex, which bool) *Vertex {
	// find in each edge
	walkings := c.subjectPlane.Rings
	if !which {
		walkings = c.clippingPlane.Rings
	}
	for _, w := range walkings {
		if iter, err := SliceContains(w.Vertexs, start); err == nil {
			for {
				pol.AddPointWhich(&w.Vertexs[iter], which)

				if w.IsClockwise {
					iter--
				} else {
					iter++
				}

				// 循环列表
				if iter == len(w.Vertexs) {
					iter = 0
				}
				if iter == -1 {
					iter = len(w.Vertexs) - 1
				}

				if w.Vertexs[iter].IsIntersectionPoint {
					break
				}
			}
			return &w.Vertexs[iter]
		}
	}
	// should not happend
	return &Vertex{}
}

// ComputeClipOverlay merge overlay polygon.
type ComputeClipOverlay struct {
	*PolygonOverlay
}

// Next overlay polygon.
func (c *ComputeClipOverlay) Next(pol *Plane, start *Vertex) *Vertex {
	next := c.Compute(pol, start, true)
	next = c.Compute(pol, next, false)
	return next
}

// Compute overlay polygon.
func (c *ComputeClipOverlay) Compute(pol *Plane, start *Vertex, which bool) *Vertex {
	// find in each edge
	walkings := c.subjectPlane.Rings
	if !which {
		walkings = c.clippingPlane.Rings
	}
	for _, w := range walkings {
		if iter, err := SliceContains(w.Vertexs, start); err == nil {
			for {
				pol.AddPointWhich(&w.Vertexs[iter], which)
				if w.IsClockwise {
					iter++
				} else {
					iter--
				}
				// 循环列表
				if iter == len(w.Vertexs) {
					iter = 0
				}
				if iter == -1 {
					iter = len(w.Vertexs) - 1
				}

				if w.Vertexs[iter].IsIntersectionPoint {
					break
				}
			}
			return &w.Vertexs[iter]
		}
	}
	// should not happend
	return &Vertex{}
}

// ComputeMainOverlay merge overlay polygon.
type ComputeMainOverlay struct {
	*PolygonOverlay
}

// Next overlay polygon.
func (c *ComputeMainOverlay) Next(pol *Plane, start *Vertex) *Vertex {
	next := c.Compute(pol, start, true)
	next = c.Compute(pol, next, false)
	return next
}

// Compute overlay polygon.
func (c *ComputeMainOverlay) Compute(pol *Plane, start *Vertex, which bool) *Vertex {
	// find in each edge
	walkings := c.subjectPlane.Rings
	if !which {
		walkings = c.clippingPlane.Rings
	}
	for _, w := range walkings {
		if iter, err := SliceContains(w.Vertexs, start); err == nil {
			for {
				pol.AddPointWhich(&w.Vertexs[iter], which)

				if which {
					if w.IsClockwise {
						iter--
					} else {
						iter++
					}
				} else {
					if w.IsClockwise {
						iter++
					} else {
						iter--
					}
				}

				// 循环列表
				if iter == len(w.Vertexs) {
					iter = 0
				}
				if iter == -1 {
					iter = len(w.Vertexs) - 1
				}

				if w.Vertexs[iter].IsIntersectionPoint {
					break
				}
			}
			return &w.Vertexs[iter]
		}
	}
	// should not happend
	return &Vertex{}
}
