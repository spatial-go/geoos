package sweepline

// OverlapAction An action taken when a SweepLineIndex detects that two Interval.
type OverlapAction interface {
	Overlap(s0, s1 *Interval)
}

// Interval ...
type Interval struct {
	Min, Max float64
	Item     interface{}
}

// CoordinatesOverlapAction ...
type CoordinatesOverlapAction struct {
}

// Overlap ...
func (c *CoordinatesOverlapAction) Overlap(s0, s1 *Interval) {
	//TODO
}
