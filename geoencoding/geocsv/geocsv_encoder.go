package geocsv

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

type GeocsvEncoder struct {
	geojson.BaseEncoder
}

// Encode Returns string of that encode geometry  by codeType.
func (e *GeocsvEncoder) Encode(g space.Geometry) []byte {
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
func (e *GeocsvEncoder) Decode(s []byte) (space.Geometry, error) {
	b := bytes.NewReader(s)
	options := Options{
		XField: "x",
		YField: "y",
	}
	if gc, err := ReadByte(b, options); err != nil {
		log.Printf("GeoCSV.Read() error = %v", err)
		return nil, err
	} else {
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
}

// Read Returns geometry from reader.
func (e *GeocsvEncoder) Read(r io.Reader) (space.Geometry, error) {
	if b, err := e.ReadBytes(r); err != nil {
		return nil, err
	} else {
		return e.Decode(b)
	}
}

// Write write geometry to reader.
func (e *GeocsvEncoder) Write(w io.Writer, g space.Geometry) error {
	b := e.Encode(g)
	return e.WriteBytes(w, b)
}

// WriteGeoJSON write geometry to writer.
func (e *GeocsvEncoder) WriteGeoJSON(w io.Writer, g *geojson.FeatureCollection) error {
	colls := space.Collection{}
	for _, v := range g.Features {
		colls = append(colls, v.Geometry.Geometry())
	}
	return e.Write(w, colls)
}

// ReadGeoJSON Returns geometry from reader .
func (e *GeocsvEncoder) ReadGeoJSON(r io.Reader) (*geojson.FeatureCollection, error) {
	if geom, err := e.Read(r); err != nil {
		return nil, err
	} else {
		return geojson.GeometryToFeatureCollection(geom), nil
	}
}
