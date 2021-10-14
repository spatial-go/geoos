package geoos

import (
	"testing"
)

func TestGeoosTestTag(t *testing.T) {
	t.Run("", func(t *testing.T) {
		if !GeoosTestTag {
			t.Errorf("target has be true")
		}
	})
}
