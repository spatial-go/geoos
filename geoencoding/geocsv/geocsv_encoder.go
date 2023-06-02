package geocsv

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

// Encoder defines csv encoder.
type Encoder struct {
	geojson.BaseEncoder
}

// Encode Returns string of that encode geometry  by codeType.
func (e *Encoder) Encode(g space.Geometry) []byte {
	gc := NewGeoCSV()
	gc.options = Options{
		XField: "x",
		YField: "y",
	}
	gc.coll = g.(space.Collection)
	buf := new(bytes.Buffer)
	buf.WriteString("way_id,pt_id,x,y\n")
	for i, f := range gc.coll {
		str := fmt.Sprintf("%v,%v,%v,%v\n", i, i, f.(space.Point)[0], f.(space.Point)[1])
		buf.WriteString(str)
	}
	return buf.Bytes()
}

// Decode Returns geometry of that decode string by codeType.
func (e *Encoder) Decode(s []byte) (space.Geometry, error) {
	b := bytes.NewReader(s)
	options := Options{
		XField: "x",
		YField: "y",
	}
	gc, err := ReadByte(b, options)
	if err != nil {
		log.Printf("GeoCSV.Read() error = %v", err)
		return nil, err
	}
	if len(gc.headers) != 4 {
		log.Printf("length of headers is wrong")
	}

	features := gc.ToGeoJSON()

	coll := make(space.Collection, len(features.Features))
	for i, f := range features.Features {
		coll[i] = f.Geometry.Coordinates.(space.Point)
	}
	return coll, nil
}

// Read Returns geometry from reader.
func (e *Encoder) Read(r io.Reader) (space.Geometry, error) {
	b, err := e.ReadBytes(r)
	if err != nil {
		return nil, err
	}
	return e.Decode(b)
}

// Write write geometry to reader.
func (e *Encoder) Write(w io.Writer, g space.Geometry) error {
	b := e.Encode(g)
	return e.WriteBytes(w, b)
}

// WriteGeoJSON write geometry to writer.
func (e *Encoder) WriteGeoJSON(w io.Writer, g *geojson.FeatureCollection) error {
	colls := space.Collection{}
	for _, v := range g.Features {
		colls = append(colls, v.Geometry.Geometry())
	}
	return e.Write(w, colls)
}

// ReadGeoJSON Returns geometry from reader .
func (e *Encoder) ReadGeoJSON(r io.Reader) (*geojson.FeatureCollection, error) {
	geom, err := e.Read(r)
	if err != nil {
		return nil, err
	}
	return geojson.GeometryToFeatureCollection(geom), nil
}
