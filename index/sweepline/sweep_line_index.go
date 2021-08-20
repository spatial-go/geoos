//Package sweepline Contains struct which implement a sweepline algorithm for scanning geometric data structures.
package sweepline

import (
	"sort"
)

// Index A sweepline implements a sorted index on a set of intervals.
// It is used to compute all overlaps between the interval in the index.
type Index struct {
	events     LineEvents
	indexBuilt bool
	nOverlaps  int
}

// Add ...
func (s *Index) Add(sweepInt *Interval) {
	insertEvent := NewEvent(sweepInt.Min, nil, sweepInt)
	s.events = append(s.events, insertEvent)
	s.events = append(s.events, NewEvent(sweepInt.Max, insertEvent, sweepInt))
}

// buildIndex Because Delete Events have a link to their corresponding Insert event,
//  it is possible to compute exactly the range of events which must be
//  compared to a given Insert event object.
func (s *Index) buildIndex() {
	if s.indexBuilt {
		return
	}
	sort.Sort(s.events)
	for i, ev := range s.events {
		if ev.IsDelete() {
			ev.InsertEvent.DeleteEventIndex = i
		}
	}
	s.indexBuilt = true
}

// ComputeOverlaps  compute overlaps.
func (s *Index) ComputeOverlaps(action OverlapAction) {
	s.nOverlaps = 0
	s.buildIndex()

	for i, ev := range s.events {
		if ev.IsInsert() {
			s.processOverlaps(i, ev.DeleteEventIndex, ev.SweepInt, action)
		}
	}
}

func (s *Index) processOverlaps(start, end int, s0 *Interval, action OverlapAction) {
	/**
	 * Since we might need to test for self-intersections,
	 * include current insert event object in list of event objects to test.
	 * Last index can be skipped, because it must be a Delete event.
	 */
	for i := start; i < end; i++ {
		ev := s.events[i]
		if ev.IsInsert() {
			action.Overlap(s0, ev.SweepInt)
			s.nOverlaps++
		}
	}
}

// LineEvents ...
type LineEvents []*Event

// Len ...
func (it LineEvents) Len() int {
	return len(it)
}

// Less ...
func (it LineEvents) Less(i, j int) bool {
	return it[i].Compare(it[j]) == -1
}

// Swap ...
func (it LineEvents) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
