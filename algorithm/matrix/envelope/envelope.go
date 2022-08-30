// Package envelope Defines a rectangular region of the 2D coordinate plane.
package envelope

import (
	"fmt"
	"math"
	"strconv"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Envelope Defines a rectangular region of the 2D coordinate plane.
//  It is often used to represent the bounding box of a  Geometry,
//  Envelopes support infinite or half-infinite regions, by using the values of
//  Double.POSITIVE_INFINITY and Double.NEGATIVE_INFINITY.
//  Envelope objects may have a null value.
//  When Envelope objects are created or initialized,
//  the supplies extent values are automatically sorted into the correct order.
type Envelope struct {
	MaxX, MinX, MaxY, MinY float64
}

// HashCode Computes a hash code for a double value, using the algorithm from
// Joshua Bloch's book
func (e *Envelope) HashCode() int {
	//Algorithm from Effective Java by Joshua Bloch [Jon Aquino]
	result := 17
	result = 37*result + e.hashCode(e.MinX)
	result = 37*result + e.hashCode(e.MaxX)
	result = 37*result + e.hashCode(e.MinY)
	result = 37*result + e.hashCode(e.MaxY)
	return result
}

// hashCode Computes a hash code for a double value, using the algorithm from
// Joshua Bloch's book
func (e *Envelope) hashCode(x float64) int {
	f, _ := strconv.Atoi(strconv.FormatFloat(x, 'f', 0, 64))
	return f
}

// IsIntersects Test the point q to see whether it intersects the Envelope defined by p1-p2
func IsIntersects(p1, p2, q matrix.Matrix) bool {
	//OptimizeIt shows that Math#min and Math#max here are a bottleneck.
	if q[0] >= math.Min(p1[0], p2[0]) && q[0] <= math.Max(p1[0], p2[0]) &&
		q[1] >= math.Min(p1[1], p2[1]) && q[1] <= math.Max(p1[1], p2[1]) {
		return true
	}
	return false
}

// IsIntersectsTwo  Tests whether the envelope defined by p1-p2
// and the envelope defined by q1-q2 intersect.
func IsIntersectsTwo(p1, p2, q1, q2 matrix.Matrix) bool {
	minQ := math.Min(q1[0], q2[0])
	maxQ := math.Max(q1[0], q2[0])
	minP := math.Min(p1[0], p2[0])
	maxP := math.Max(p1[0], p2[0])

	if minP > maxQ || maxP < minQ {
		return false
	}

	minQ = math.Min(q1[1], q2[1])
	maxQ = math.Max(q1[1], q2[1])
	minP = math.Min(p1[1], p2[1])
	maxP = math.Max(p1[1], p2[1])
	if minP > maxQ || maxP < minQ {
		return false
	}
	return true
}

// Empty  Creates a null Envelope.
func Empty() *Envelope {
	el := &Envelope{}
	el.SetToNil()
	return el
}

// FourFloat  Creates an Envelope for a region defined by maximum and minimum values.
func FourFloat(x1, x2, y1, y2 float64) *Envelope {
	el := &Envelope{}
	el.initXY(x1, x2, y1, y2)
	return el
}

// TwoMatrix  Creates an Envelope for a region defined by two matrix.
func TwoMatrix(p1, p2 matrix.Matrix) *Envelope {
	el := &Envelope{}
	el.initXY(math.Min(p1[0], p2[0]), math.Max(p1[0], p2[0]), math.Min(p1[1], p2[1]), math.Max(p1[1], p2[1]))
	return el
}

// MatrixList Creates an Envelope from a matrix list
func MatrixList(ps []matrix.Matrix) *Envelope {
	el := &Envelope{}
	el.SetToNil()
	for _, p := range ps {
		el.ExpandToIncludeMatrix(p)
	}
	return el
}

// PolygonMatrixList Creates an Envelope from a polygon matrix list
func PolygonMatrixList(ps []matrix.PolygonMatrix) *Envelope {
	el := &Envelope{}
	el.SetToNil()
	for _, polygon := range ps {
		for _, part := range polygon {
			for _, p := range part {
				el.ExpandToInclude(p[0], p[1])
			}
		}
	}
	return el
}

// Matrix  Creates an Envelope for a region defined by a single matrix.
func Matrix(p matrix.Matrix) *Envelope {
	el := &Envelope{}
	el.initXY(p[0], p[0], p[1], p[1])
	return el
}

// Bound  Creates an Envelope for a region defined by a single matrix.
func Bound(b []matrix.Matrix) *Envelope {
	el := TwoMatrix(b[0], b[1])
	return el
}

// Env  Create an Envelope from an existing Envelope.
func Env(env *Envelope) *Envelope {
	el := &Envelope{}
	el.initEnvelope(env)
	return el
}

// initXY Initialize an Envelope for a region defined by maximum and minimum values.
func (e *Envelope) initXY(x1, x2, y1, y2 float64) {
	if x1 < x2 {
		e.MinX = x1
		e.MaxX = x2
	} else {
		e.MinX = x2
		e.MaxX = x1
	}
	if y1 < y2 {
		e.MinY = y1
		e.MaxY = y2
	} else {
		e.MinY = y2
		e.MaxY = y1
	}
}

// Copy Creates a copy of this envelope object.
func (e *Envelope) Copy() *Envelope {
	return Env(e)
}

// initEnvelope Initialize an Envelope from an existing Envelope.
func (e *Envelope) initEnvelope(env *Envelope) {
	e.MinX = env.MinX
	e.MaxX = env.MaxX
	e.MinY = env.MinY
	e.MaxY = env.MaxY
}

// SetToNil Makes this Envelope a "null" envelope, that is, the envelope
//  of the empty geometry.
func (e *Envelope) SetToNil() {
	e.MinX = 0
	e.MaxX = -1
	e.MinY = 0
	e.MaxY = -1
}

// IsNil  Returns true if this Envelope is a "nil" envelope.
func (e *Envelope) IsNil() bool {
	if e == nil {
		return true
	}
	return e.MaxX < e.MinX
}

// Width Returns the difference between the maximum and minimum x values.
func (e *Envelope) Width() float64 {
	if e.IsNil() {
		return 0
	}
	return e.MaxX - e.MinX
}

// Height  Returns the difference between the maximum and minimum y values.
func (e *Envelope) Height() float64 {
	if e.IsNil() {
		return 0
	}
	return e.MaxY - e.MinY
}

// Diameter  Gets the length of the diameter (diagonal) of the envelopArea.
func (e *Envelope) Diameter() float64 {
	if e.IsNil() {
		return 0
	}
	w := e.Width()
	h := e.Height()
	return math.Sqrt(w*w + h*h)
}

// Area Gets the area of this envelope.
func (e *Envelope) Area() float64 {
	return e.Width() * e.Height()
}

// MinExtent  Gets the minimum extent of this envelope across both dimensions.
func (e *Envelope) MinExtent() float64 {
	if e.IsNil() {
		return 0.0
	}
	w := e.Width()
	h := e.Height()
	if w < h {
		return w
	}
	return h
}

// MaxExtent Gets the maximum extent of this envelope across both dimensions.
func (e *Envelope) MaxExtent() float64 {
	if e.IsNil() {
		return 0.0
	}
	w := e.Width()
	h := e.Height()
	if w > h {
		return w
	}
	return h
}

// ExpandToIncludeMatrix Enlarges this Envelope so that it contains
func (e *Envelope) ExpandToIncludeMatrix(p matrix.Matrix) {
	e.ExpandToInclude(p[0], p[1])
}

// ExpandBy Expands this envelope by a given distance in all directions.
// Both positive and negative distances are supported.
func (e *Envelope) ExpandBy(distance float64) {
	e.ExpandByXY(distance, distance)
}

// ExpandByXY  Expands this envelope by a given distance in all directions.
// Both positive and negative distances are supported.
func (e *Envelope) ExpandByXY(deltaX, deltaY float64) {
	if e.IsNil() {
		return
	}

	e.MinX -= deltaX
	e.MaxX += deltaX
	e.MinY -= deltaY
	e.MaxY += deltaY

	// check for envelope disappearing
	if e.MinX > e.MaxX || e.MinY > e.MaxY {
		e.SetToNil()
	}
}

// ExpandToInclude  Enlarges this Envelope so that it contains the given point.
// Has no effect if the point is already on or within the envelope.
func (e *Envelope) ExpandToInclude(x, y float64) {
	if e.IsNil() {
		e.MinX = x
		e.MaxX = x
		e.MinY = y
		e.MaxY = y
	} else {
		if x < e.MinX {
			e.MinX = x
		}
		if x > e.MaxX {
			e.MaxX = x
		}
		if y < e.MinY {
			e.MinY = y
		}
		if y > e.MaxY {
			e.MaxY = y
		}
	}
}

// ExpandToIncludeEnv  Enlarges this Envelope so that it contains  the other Envelope.
//  Has no effect if other is wholly on or within the envelope.
func (e *Envelope) ExpandToIncludeEnv(other *Envelope) {
	if other.IsNil() {
		return
	}
	if e.IsNil() {
		e.MinX = other.MinX
		e.MaxX = other.MaxX
		e.MinY = other.MinY
		e.MaxY = other.MaxY
	} else {
		if other.MinX < e.MinX {
			e.MinX = other.MinX
		}
		if other.MaxX > e.MaxX {
			e.MaxX = other.MaxX
		}
		if other.MinY < e.MinY {
			e.MinY = other.MinY
		}
		if other.MaxY > e.MaxY {
			e.MaxY = other.MaxY
		}
	}
}

// Translate Translates this envelope by given amounts in the X and Y direction.
func (e *Envelope) Translate(transX, transY float64) {
	if e.IsNil() {
		return
	}
	e.initXY(e.MinX+transX, e.MaxX+transX,
		e.MinY+transY, e.MaxY+transY)
}

// Centre Computes the coordinate of the centre of this envelope (as long as it is non-null
func (e *Envelope) Centre() matrix.Matrix {
	if e.IsNil() {
		return nil
	}
	return matrix.Matrix{
		(e.MinX + e.MaxX) / 2.0,
		(e.MinY + e.MaxY) / 2.0,
	}
}

// Intersection Computes the intersection of two  Envelopes.
func (e *Envelope) Intersection(env *Envelope) *Envelope {
	if e.IsNil() || env.IsNil() || !e.IsIntersects(env) {
		return &Envelope{}
	}
	e.MinX = math.Max(e.MinX, env.MinX)
	e.MinY = math.Max(e.MinY, env.MinY)
	e.MaxX = math.Min(e.MaxX, env.MaxX)
	e.MaxY = math.Min(e.MaxY, env.MaxY)
	return FourFloat(e.MinX, e.MaxX, e.MinY, e.MaxY)
}

// IsIntersects Tests if the region defined by other
// intersects the region of this Envelope.
func (e *Envelope) IsIntersects(other *Envelope) bool {
	if e.IsNil() || other.IsNil() {
		return false
	}
	return !(other.MinX > e.MaxX ||
		other.MaxX < e.MinX ||
		other.MinY > e.MaxY ||
		other.MaxY < e.MinY)
}

// Disjoint Tests if the region defined by other
// is disjoint from the region of this Envelope.
func (e *Envelope) Disjoint(other *Envelope) bool {
	if e.IsNil() || other.IsNil() {
		return true
	}
	return other.MinX > e.MaxX ||
		other.MaxX < e.MinX ||
		other.MinY > e.MaxY ||
		other.MaxY < e.MinY
}

// Overlaps overlaps may be changed to be a true overlap check; that is,
// whether the intersection is two-dimensional.
func (e *Envelope) Overlaps(other *Envelope) bool {
	return e.IsIntersects(other)
}

// Contains Tests if the Envelope other lies wholely inside this Envelope (inclusive of the boundary).
func (e *Envelope) Contains(other *Envelope) bool {
	return e.Covers(other)
}

// Covers Tests if the Envelope other
// lies wholely inside this Envelope (inclusive of the boundary).
func (e *Envelope) Covers(other *Envelope) bool {
	if e.IsNil() || other.IsNil() {
		return false
	}
	return other.MinX >= e.MinX &&
		other.MaxX <= e.MaxX &&
		other.MinY >= e.MinY &&
		other.MaxY <= e.MaxY
}

// Distance Computes the distance between this and another Envelope.
func (e *Envelope) Distance(env *Envelope) float64 {
	if e.IsIntersects(env) {
		return 0
	}
	dx := 0.0
	if e.MaxX < env.MinX {
		dx = env.MinX - e.MaxX
	} else if e.MinX > env.MaxX {
		dx = e.MinX - env.MaxX
	}
	dy := 0.0
	if e.MaxY < env.MinY {
		dy = env.MinY - e.MaxY
	} else if e.MinY > env.MaxY {
		dy = e.MinY - env.MaxY
	}
	// if either is zero, the envelopes overlap either vertically or horizontally
	if dx == 0.0 {
		return dy
	}
	if dy == 0.0 {
		return dx
	}
	return math.Sqrt(dx*dx + dy*dy)
}

// Equals ...
func (e *Envelope) Equals(other *Envelope) bool {
	if e.IsNil() {
		return other.IsNil()
	}
	return e.MaxX == other.MaxX &&
		e.MaxY == other.MaxY &&
		e.MinX == other.MinX &&
		e.MinY == other.MinY
}

// Proximity ...
func (e *Envelope) Proximity(other *Envelope) bool {
	if e.IsNil() {
		return other.IsNil()
	}
	return matrix.Matrix{e.MaxX, e.MaxY}.Proximity(matrix.Matrix{other.MaxX, other.MaxY}) &&
		matrix.Matrix{e.MinX, e.MinY}.Proximity(matrix.Matrix{other.MinX, other.MinY})
}

// ToString ...
func (e *Envelope) ToString() string {
	return fmt.Sprintf("Env[%v : %v, %v : %v]", e.MinX, e.MaxX, e.MinY, e.MaxY)
}

// CompareTo Compares two envelopes using lexicographic ordering.
func (e *Envelope) CompareTo(other *Envelope) int {
	// compare nulls if present
	if e.IsNil() {
		if other.IsNil() {
			return 0
		}
		return -1
	}
	if other.IsNil() {
		return 1
	}

	// compare based on numerical ordering of ordinates
	if (e.MinX < other.MinX) || (e.MinY < other.MinY) ||
		(e.MaxX < other.MaxX) || (e.MaxY < other.MaxY) {
		return -1
	}
	if (e.MinX > other.MinX) || (e.MinY > other.MinY) ||
		(e.MaxX > other.MaxX) || (e.MaxY > other.MaxY) {
		return 1
	}
	return 0

}

// ToMatrix ...
func (e *Envelope) ToMatrix() *matrix.PolygonMatrix {
	return &matrix.PolygonMatrix{
		{
			{e.MinX, e.MinY},
			{e.MinX, e.MaxY},
			{e.MaxX, e.MaxY},
			{e.MaxX, e.MinY},
			{e.MinX, e.MinY},
		},
	}
}
