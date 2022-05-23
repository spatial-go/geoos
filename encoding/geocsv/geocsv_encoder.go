package geocsv

import (
	"bytes"
	"fmt"
	"log"

	"github.com/spatial-go/geoos/space"
)

type GeocsvEncoder struct {
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
		if len(gc.rows) != 4 {
			log.Printf("length of rows is wrong")
		}
		features := gc.ToGeoJSON()
		if len(features.Features) != 4 {
			log.Printf("length of features is wrong")
		}
		coll := make(space.Collection, len(features.Features))
		for i, f := range features.Features {
			coll[i] = f.Geometry.Coordinates.(space.Point)
		}
		return coll, nil
	}
}
