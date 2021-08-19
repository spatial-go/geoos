package sweepline

// Sweepline const parameter.
const (
	InsertStatus = iota + 1
	DeleteStatus
)

// Event ...
type Event struct {
	xValue           float64
	eventType        int
	InsertEvent      *Event // null if this is an INSERT event
	DeleteEventIndex int

	SweepInt *Interval
}

// SweepLineEvent create a default SweepLineEvent.
func (s *Event) SweepLineEvent(x float64, insertEvent *Event, sweepInt *Interval) *Event {
	s.xValue = x
	s.InsertEvent = insertEvent
	s.eventType = InsertStatus
	if s.InsertEvent != nil {
		s.eventType = DeleteStatus
	}
	s.SweepInt = sweepInt
	return s
}

// IsInsert ...
func (s *Event) IsInsert() bool {
	return s.InsertEvent == nil
}

// IsDelete ...
func (s *Event) IsDelete() bool {
	return s.InsertEvent != nil
}

// Compare ProjectionEvents are ordered first by their x-value, and then by their eventType.
//  It is important that Insert events are sorted before Delete events, so that
//  items whose Insert and Delete events occur at the same x-value will be correctly handled.
func (s *Event) Compare(pe *Event) int {
	if s.xValue < pe.xValue {
		return -1
	}
	if s.xValue > pe.xValue {
		return 1
	}
	if s.eventType < pe.eventType {
		return -1
	}
	if s.eventType > pe.eventType {
		return 1
	}
	return 0
}
