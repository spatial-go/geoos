package encode

import (
	"github.com/spatial-go/geoos/encoding/geobuf/utils"
	"github.com/spatial-go/geoos/encoding/geojson"
	geoos "github.com/spatial-go/geoos/space"
)

// EncodingConfig ...
type EncodingConfig struct {
	Dimension uint
	Precision uint
	Keys      utils.KeyStore
}

// EncodingOption ...
type EncodingOption func(o *EncodingConfig)

// WithPrecision ...
func WithPrecision(precision uint) EncodingOption {
	return func(o *EncodingConfig) {
		o.Precision = uint(utils.DecodePrecision(uint32(precision)))
	}
}

// WithDimension ...
func WithDimension(dimension uint) EncodingOption {
	return func(o *EncodingConfig) {
		o.Dimension = dimension
	}
}

// WithKeyStore ...
func WithKeyStore(store utils.KeyStore) EncodingOption {
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
		case geoos.TypePoint:
			updatePrecision(t.Coordinates.(geoos.Point), opts)
		case geoos.TypeMultiPoint:
			coords := t.Coordinates.(geoos.MultiPoint)
			for _, coord := range coords {
				updatePrecision(coord, opts)
			}
		case geoos.TypeLineString:
			coords := t.Coordinates.(geoos.LineString)
			for _, coord := range coords {
				updatePrecision(coord, opts)
			}
		case geoos.TypeMultiLineString:
			lines := t.Coordinates.(geoos.MultiLineString)
			for _, line := range lines {
				for _, coord := range line {
					updatePrecision(coord, opts)
				}
			}
		case geoos.TypePolygon:
			lines := t.Coordinates.(geoos.Polygon)
			for _, line := range lines {
				for _, coord := range line {
					updatePrecision(coord, opts)
				}
			}
		case geoos.TypeMultiPolygon:
			polygons := t.Coordinates.(geoos.MultiPolygon)
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

func updatePrecision(point geoos.Point, opt *EncodingConfig) {
	for _, val := range point {
		e := utils.GetPrecision(val)
		if e > opt.Precision {
			opt.Precision = e
		}
	}
}
