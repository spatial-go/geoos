package encode

import (
	"github.com/spatial-go/geoos/geoencoding/geobuf/protogeo"
	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

// EncodingConfig ...
type EncodingConfig struct {
	Dimension uint
	Precision uint
	Keys      protogeo.KeyStore
}

// EncodingOption ...
type EncodingOption func(o *EncodingConfig)

// WithPrecision ...
func WithPrecision(precision uint) EncodingOption {
	return func(o *EncodingConfig) {
		o.Precision = uint(protogeo.DecodePrecision(uint32(precision)))
	}
}

// WithDimension ...
func WithDimension(dimension uint) EncodingOption {
	return func(o *EncodingConfig) {
		o.Dimension = dimension
	}
}

// WithKeyStore ...
func WithKeyStore(store protogeo.KeyStore) EncodingOption {
	return func(o *EncodingConfig) {
		o.Keys = store
	}
}

// FromAnalysis ...
func FromAnalysis(obj interface{}) EncodingOption {
	return func(o *EncodingConfig) {
		analyze(obj, o)
	}
}

func analyze(obj interface{}, opts *EncodingConfig) {
	opts.Dimension = 2
	switch t := obj.(type) {
	case *geojson.FeatureCollection:
		for _, feature := range t.Features {
			analyze(feature, opts)
		}
	case *geojson.Feature:
		analyze(geojson.NewGeometry(t.Geometry.Geometry()), opts)
		for key := range t.Properties {
			opts.Keys.Add(key)
		}
	case *geojson.Geometry:
		switch t.Type {
		case space.TypePoint:
			updatePrecision(t.Coordinates.(space.Point), opts)
		case space.TypeMultiPoint:
			coords := t.Coordinates.(space.MultiPoint)
			for _, coord := range coords {
				updatePrecision(coord, opts)
			}
		case space.TypeLineString:
			coords := t.Coordinates.(space.LineString)
			for _, coord := range coords {
				updatePrecision(coord, opts)
			}
		case space.TypeMultiLineString:
			lines := t.Coordinates.(space.MultiLineString)
			for _, line := range lines {
				for _, coord := range line {
					updatePrecision(coord, opts)
				}
			}
		case space.TypePolygon:
			lines := t.Coordinates.(space.Polygon)
			for _, line := range lines {
				for _, coord := range line {
					updatePrecision(coord, opts)
				}
			}
		case space.TypeMultiPolygon:
			polygons := t.Coordinates.(space.MultiPolygon)
			for _, rings := range polygons {
				for _, ring := range rings {
					for _, coord := range ring {
						updatePrecision(coord, opts)
					}
				}
			}
		}
	}

}

func updatePrecision(point space.Point, opt *EncodingConfig) {
	for _, val := range point {
		e := protogeo.GetPrecision(val)
		if e > opt.Precision {
			opt.Precision = e
		}
	}
}
