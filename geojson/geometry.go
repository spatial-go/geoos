package geojson

import (
	"encoding/json"
	"errors"

	"github.com/spatial-go/geoos"
)

// ErrInvalidGeometry will be returned if a the json of the geometry is invalid.
var ErrInvalidGeometry = errors.New("geojson: invalid geometry")

// A Geometry matches the structure of a GeoJSON Geometry.
type Geometry struct {
	Type        string         `json:"type"`
	Coordinates geoos.Geometry `json:"coordinates,omitempty"`
	Geometries  []*Geometry    `json:"geometries,omitempty"`
}

// NewGeometry will create a Geometry object but will convert
// the input into a GoeJSON geometry. For example, it will convert
// Rings and Bounds into Polygons.
func NewGeometry(g geoos.Geometry) *Geometry {
	jg := &Geometry{}
	switch g := g.(type) {
	case geoos.Ring:
		jg.Coordinates = geoos.Polygon{g}
	case geoos.Bound:
		jg.Coordinates = g.ToPolygon()
	case geoos.Collection:
		for _, c := range g {
			jg.Geometries = append(jg.Geometries, NewGeometry(c))
		}
		jg.Type = g.GeoJSONType()
	default:
		jg.Coordinates = g
	}

	if jg.Coordinates != nil {
		jg.Type = jg.Coordinates.GeoJSONType()
	}
	return jg
}

// Geometry returns the geoos.Geometry for the geojson Geometry.
// This will convert the "Geometries" into a geoos.Collection if applicable.
func (g Geometry) Geometry() geoos.Geometry {
	if g.Coordinates != nil {
		return g.Coordinates
	}

	c := make(geoos.Collection, 0, len(g.Geometries))
	for _, geom := range g.Geometries {
		c = append(c, geom.Geometry())
	}
	return c
}

// MarshalJSON will marshal the geometry into the correct json structure.
func (g Geometry) MarshalJSON() ([]byte, error) {
	if g.Coordinates == nil && len(g.Geometries) == 0 {
		return []byte(`null`), nil
	}

	ng := &jsonGeometryMarshall{}
	switch g := g.Coordinates.(type) {
	case geoos.Ring:
		ng.Coordinates = geoos.Polygon{g}
	case geoos.Bound:
		ng.Coordinates = g.ToPolygon()
	case geoos.Collection:
		ng.Geometries = make([]*Geometry, 0, len(g))
		for _, c := range g {
			ng.Geometries = append(ng.Geometries, NewGeometry(c))
		}
		ng.Type = g.GeoJSONType()
	default:
		ng.Coordinates = g
	}

	if ng.Coordinates != nil {
		ng.Type = ng.Coordinates.GeoJSONType()
	}

	if len(g.Geometries) > 0 {
		ng.Geometries = g.Geometries
		ng.Type = geoos.Collection{}.GeoJSONType()
	}
	return json.Marshal(ng)
}

// UnmarshalGeometry decodes the data into a GeoJSON feature.
// Alternately one can call json.Unmarshal(g) directly for the same result.
func UnmarshalGeometry(data []byte) (*Geometry, error) {
	g := &Geometry{}
	err := json.Unmarshal(data, g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// UnmarshalJSON will unmarshal the correct geometry from the json structure.
func (g *Geometry) UnmarshalJSON(data []byte) error {
	jg := &jsonGeometry{}
	err := json.Unmarshal(data, jg)
	if err != nil {
		return err
	}

	switch jg.Type {
	case "Point":
		p := geoos.Point{}
		err = json.Unmarshal(jg.Coordinates, &p)
		g.Coordinates = p
	case "MultiPoint":
		mp := geoos.MultiPoint{}
		err = json.Unmarshal(jg.Coordinates, &mp)
		g.Coordinates = mp
	case "LineString":
		ls := geoos.LineString{}
		err = json.Unmarshal(jg.Coordinates, &ls)
		g.Coordinates = ls
	case "MultiLineString":
		mls := geoos.MultiLineString{}
		err = json.Unmarshal(jg.Coordinates, &mls)
		g.Coordinates = mls
	case "Polygon":
		p := geoos.Polygon{}
		err = json.Unmarshal(jg.Coordinates, &p)
		g.Coordinates = p
	case "MultiPolygon":
		mp := geoos.MultiPolygon{}
		err = json.Unmarshal(jg.Coordinates, &mp)
		g.Coordinates = mp
	case "GeometryCollection":
		g.Geometries = jg.Geometries
	default:
		return ErrInvalidGeometry
	}

	g.Type = g.Geometry().GeoJSONType()

	return nil
}

// A Point is a helper type that will marshal to/from a GeoJSON Point geometry.
type Point geoos.Point

// Geometry will return the geoos.Geometry version of the data.
func (p Point) Geometry() geoos.Geometry {
	return geoos.Point(p)
}

// MarshalJSON will convert the Point into a GeoJSON Point geometry.
func (p Point) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: geoos.Point(p)})
}

// UnmarshalJSON will unmarshal the GeoJSON Point geometry.
func (p *Point) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}

	point, ok := g.Coordinates.(geoos.Point)
	if !ok {
		return errors.New("geojson: not a Point type")
	}

	*p = Point(point)
	return nil
}

// A MultiPoint is a helper type that will marshal to/from a GeoJSON MultiPoint geometry.
type MultiPoint geoos.MultiPoint

// Geometry will return the geoos.Geometry version of the data.
func (mp MultiPoint) Geometry() geoos.Geometry {
	return geoos.MultiPoint(mp)
}

// MarshalJSON will convert the MultiPoint into a GeoJSON MultiPoint geometry.
func (mp MultiPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: geoos.MultiPoint(mp)})
}

// UnmarshalJSON will unmarshal the GeoJSON MultiPoint geometry.
func (mp *MultiPoint) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}

	multiPoint, ok := g.Coordinates.(geoos.MultiPoint)
	if !ok {
		return errors.New("geojson: not a MultiPoint type")
	}

	*mp = MultiPoint(multiPoint)
	return nil
}

// A LineString is a helper type that will marshal to/from a GeoJSON LineString geometry.
type LineString geoos.LineString

// Geometry will return the Geometry version of the data.
func (ls LineString) Geometry() geoos.Geometry {
	return geoos.LineString(ls)
}

// MarshalJSON will convert the LineString into a GeoJSON LineString geometry.
func (ls LineString) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: geoos.LineString(ls)})
}

// UnmarshalJSON will unmarshal the GeoJSON MultiPoint geometry.
func (ls *LineString) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}

	lineString, ok := g.Coordinates.(geoos.LineString)
	if !ok {
		return errors.New("geojson: not a LineString type")
	}

	*ls = LineString(lineString)
	return nil
}

// A MultiLineString is a helper type that will marshal to/from a GeoJSON MultiLineString geometry.
type MultiLineString geoos.MultiLineString

// Geometry will return the geoos.Geometry version of the data.
func (mls MultiLineString) Geometry() geoos.Geometry {
	return geoos.MultiLineString(mls)
}

// MarshalJSON will convert the MultiLineString into a GeoJSON MultiLineString geometry.
func (mls MultiLineString) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: geoos.MultiLineString(mls)})
}

// UnmarshalJSON will unmarshal the GeoJSON MultiPoint geometry.
func (mls *MultiLineString) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}
	multilineString, ok := g.Coordinates.(geoos.MultiLineString)
	if !ok {
		return errors.New("geojson: not a MultiLineString type")
	}

	*mls = MultiLineString(multilineString)
	return nil
}

// A Polygon is a helper type that will marshal to/from a GeoJSON Polygon geometry.
type Polygon geoos.Polygon

// Geometry will return the geoos.Geometry version of the data.
func (p Polygon) Geometry() geoos.Geometry {
	return geoos.Polygon(p)
}

// MarshalJSON will convert the Polygon into a GeoJSON Polygon geometry.
func (p Polygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: geoos.Polygon(p)})
}

// UnmarshalJSON will unmarshal the GeoJSON Polygon geometry.
func (p *Polygon) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}
	polygon, ok := g.Coordinates.(geoos.Polygon)
	if !ok {
		return errors.New("geojson: not a Polygon type")
	}

	*p = Polygon(polygon)
	return nil
}

// A MultiPolygon is a helper type that will marshal to/from a GeoJSON MultiPolygon geometry.
type MultiPolygon geoos.MultiPolygon

// Geometry will return the geoos.Geometry version of the data.
func (mp MultiPolygon) Geometry() geoos.Geometry {
	return geoos.MultiPolygon(mp)
}

// MarshalJSON will convert the MultiPolygon into a GeoJSON MultiPolygon geometry.
func (mp MultiPolygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(Geometry{Coordinates: geoos.MultiPolygon(mp)})
}

// UnmarshalJSON will unmarshal the GeoJSON MultiPolygon geometry.
func (mp *MultiPolygon) UnmarshalJSON(data []byte) error {
	g := &Geometry{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}

	multiPolygon, ok := g.Coordinates.(geoos.MultiPolygon)
	if !ok {
		return errors.New("geojson: not a MultiPolygon type")
	}

	*mp = MultiPolygon(multiPolygon)
	return nil
}

type jsonGeometry struct {
	Type        string           `json:"type"`
	Coordinates nocopyRawMessage `json:"coordinates"`
	Geometries  []*Geometry      `json:"geometries,omitempty"`
}

type jsonGeometryMarshall struct {
	Type        string         `json:"type"`
	Coordinates geoos.Geometry `json:"coordinates,omitempty"`
	Geometries  []*Geometry    `json:"geometries,omitempty"`
}

type nocopyRawMessage []byte

func (m *nocopyRawMessage) UnmarshalJSON(data []byte) error {
	*m = data
	return nil
}
