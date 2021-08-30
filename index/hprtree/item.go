package hprtree

import (
	"fmt"

	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

// Item ...
type Item struct {
	Env  *envelope.Envelope
	Item interface{}
}

// ToString ...
func (it *Item) ToString() string {
	return fmt.Sprintf("Item:%v " + it.Env.ToString())
}
