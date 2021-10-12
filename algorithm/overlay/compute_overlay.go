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
	whichPlane := c.subjectPlane
	otherPlane := c.clippingPlane
	if !which {
		whichPlane = c.clippingPlane
		otherPlane = c.subjectPlane
	}
	inHole := otherPlane.IsVertexInHole(start)
	walkings := whichPlane.Rings
	for _, w := range walkings {
		if iter, err := SliceContains(w.Vertexes, start); err == nil {
			for {
				pol.AddPointWhich(&w.Vertexes[iter], which)
				if inHole {
					if w.IsClockwise {
						iter++
					} else {
						iter--
					}
				} else {
					if w.IsClockwise {
						iter--
					} else {
						iter++
					}
				}

				// 循环列表
				if iter == len(w.Vertexes) {
					iter = 0
				}
				if iter == -1 {
					iter = len(w.Vertexes) - 1
				}

				if w.Vertexes[iter].IsIntersectionPoint {
					break
				}
			}
			return &w.Vertexes[iter]
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
	whichPlane := c.subjectPlane
	if !which {
		whichPlane = c.clippingPlane
	}
	walkings := whichPlane.Rings
	for i, w := range walkings {
		if iter, err := SliceContains(w.Vertexes, start); err == nil {
			for {
				pol.AddPointWhich(&w.Vertexes[iter], which)
				inHole := (i > 0)
				if inHole {
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
				if iter == len(w.Vertexes) {
					iter = 0
				}
				if iter == -1 {
					iter = len(w.Vertexes) - 1
				}

				if w.Vertexes[iter].IsIntersectionPoint {
					break
				}
			}
			return &w.Vertexes[iter]
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
		if iter, err := SliceContains(w.Vertexes, start); err == nil {
			for {
				pol.AddPointWhich(&w.Vertexes[iter], which)

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
				if iter == len(w.Vertexes) {
					iter = 0
				}
				if iter == -1 {
					iter = len(w.Vertexes) - 1
				}

				if w.Vertexes[iter].IsIntersectionPoint {
					break
				}
			}
			return &w.Vertexes[iter]
		}
	}
	// should not happend
	return &Vertex{}
}
