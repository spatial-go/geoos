// Example  This is an example .
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spatial-go/geoos/geoencoding"
)

func main() {
	rawJSON := []byte(`
  { "type": "FeatureCollection",
	"features": [
	  { "type": "Feature",
		"geometry": {"type": "Point", "coordinates": [102.0, 0.5]},
		"properties": {"prop0": "value0"}
	  }
	]
  }`)

	buf := new(bytes.Buffer)
	buf.Write(rawJSON)

	if geomFeatureCollection, err := geoencoding.ReadGeoJSON(buf, geoencoding.GeoJSON); err != nil {
		fmt.Printf("ReadGeoJSON error:%v", err)
	} else {
		fmt.Println(geomFeatureCollection)

		filename := "../data/geojson.json"
		if file, err := os.Create(filename); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()

			if err := geoencoding.WriteGeoJSON(file, geomFeatureCollection, geoencoding.GeoJSON); err != nil {
				fmt.Printf("WriteGeoJSON error:%v", err)
			}
		}

		filename = "../data/geojson.json"
		if file, err := os.Open(filename); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()

			if geomFeatureCollection, err := geoencoding.ReadGeoJSON(file, geoencoding.GeoJSON); err != nil {
				fmt.Printf("ReadGeoJSON error:%v", err)
			} else {
				buf := new(bytes.Buffer)
				_ = geoencoding.WriteGeoJSON(buf, geomFeatureCollection, geoencoding.GeoJSON)
				bufStr := buf.String()
				fmt.Println(bufStr)
			}
		}

		filename = "../data/geobuf.proto"
		if file, err := os.Create(filename); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()
			if err := geoencoding.WriteGeoJSON(file, geomFeatureCollection, geoencoding.Geobuf); err != nil {
				fmt.Printf("WriteGeoJSON error:%v", err)
			}
		}

		filename = "../data/geobuf.proto"
		if file, err := os.Open(filename); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()

			if geomFeatureCollection, err := geoencoding.ReadGeoJSON(file, geoencoding.Geobuf); err != nil {
				fmt.Printf("ReadGeoJSON error:%v", err)
			} else {
				buf := new(bytes.Buffer)
				_ = geoencoding.WriteGeoJSON(buf, geomFeatureCollection, geoencoding.GeoJSON)
				bufStr := buf.String()
				fmt.Println(bufStr)
			}
		}

		filename = "../data/geocsv.csv"
		if file, err := os.Create(filename); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()
			if err := geoencoding.WriteGeoJSON(file, geomFeatureCollection, geoencoding.GeoCSV); err != nil {
				fmt.Printf("WriteGeoJSON error:%v", err)
			}
		}

		filename = "../data/geocsv.csv"
		if file, err := os.Open(filename); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()

			if geomFeatureCollection, err := geoencoding.ReadGeoJSON(file, geoencoding.GeoCSV); err != nil {
				fmt.Printf("ReadGeoJSON error:%v", err)
			} else {
				buf := new(bytes.Buffer)
				_ = geoencoding.WriteGeoJSON(buf, geomFeatureCollection, geoencoding.GeoJSON)
				bufStr := buf.String()
				fmt.Println(bufStr)
			}
		}

		filename = "../data/geowkb.wkb"
		if file, err := os.Create(filename); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()
			if err := geoencoding.WriteGeoJSON(file, geomFeatureCollection, geoencoding.WKB); err != nil {
				fmt.Printf("WriteGeoJSON error:%v", err)
			}
		}

		filename = "../data/geowkb.wkb"
		if file, err := os.Open(filename); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()

			if geomFeatureCollection, err := geoencoding.ReadGeoJSON(file, geoencoding.WKB); err != nil {
				fmt.Printf("ReadGeoJSON error:%v", err)
			} else {
				buf := new(bytes.Buffer)
				_ = geoencoding.WriteGeoJSON(buf, geomFeatureCollection, geoencoding.GeoJSON)
				bufStr := buf.String()
				fmt.Println(bufStr)
			}
		}

		filename = "../data/geowkt.wkt"
		if file, err := os.Create(filename); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()
			if err := geoencoding.WriteGeoJSON(file, geomFeatureCollection, geoencoding.WKT); err != nil {
				fmt.Printf("WriteGeoJSON error:%v", err)
			}
		}

		filename = "../data/geowkt.wkt"
		if file, err := os.Open(filename); err != nil {
			fmt.Println(err)
		} else {
			defer file.Close()

			if geomFeatureCollection, err := geoencoding.ReadGeoJSON(file, geoencoding.WKT); err != nil {
				fmt.Printf("ReadGeoJSON error:%v", err)
			} else {
				buf := new(bytes.Buffer)
				_ = geoencoding.WriteGeoJSON(buf, geomFeatureCollection, geoencoding.GeoJSON)
				bufStr := buf.String()
				fmt.Println(bufStr)
			}
		}
	}
}
