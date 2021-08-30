// Package index define spatial index interface.
package index

import (
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

// SpatialIndex The basic operations supported
// implementing spatial index algorithms.
//  A spatial index typically provides a primary filter for range rectangle queries.
//  A secondary filter is required to test for exact intersection.
//  The secondary filter may consist of other kinds of tests,
//  such as testing other spatial relationships.
type SpatialIndex interface {
	// Insert Adds a spatial item with an extent specified by the given Envelope to the index
	Insert(itemEnv *envelope.Envelope, item interface{}) error

	// Query Queries the index for all items whose extents intersect the given search  Envelope
	// Note that some kinds of indexes may also return objects which do not in fact
	//  intersect the query envelope.
	Query(searchEnv *envelope.Envelope) interface{}
	// QueryVisitor Queries the index for all items whose extents intersect the given search Envelope,
	// and applies an  ItemVisitor to them.
	// Note that some kinds of indexes may also return objects which do not in fact
	// intersect the query envelope.
	QueryVisitor(searchEnv *envelope.Envelope, visitor ItemVisitor) error

	// Remove Removes a single item from the tree.
	Remove(itemEnv *envelope.Envelope, item interface{}) bool
}
