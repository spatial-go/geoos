package index

import (
	"fmt"
)

// ErrHPRInsert ...
var ErrHPRInsert = fmt.Errorf("hpr tree is built")

// ErrHPRNotIsIntersects ...
var ErrHPRNotIsIntersects = fmt.Errorf("hpr tree totalExtent is not Intersects")

// ErrRTreeQueried ...
var ErrRTreeQueried = fmt.Errorf("Index cannot be added to once it has been queried")

// ErrTreeIsNil ...
var ErrTreeIsNil = fmt.Errorf("Index is nil")

// ErrNotMatchType ...
var ErrNotMatchType = fmt.Errorf("Operation does not support not match type arguments")
