// Package debugtools defines the method of using geojson for debugging.
package debugtools

import (
	"log"
	"os"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/geoencoding"
	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/space"
)

const dir = "/debugtools/debug_data/"

// WriteGeom Write the geom object to a file in geojson format.
func WriteGeom(filename string, geom space.Geometry) {
	file, err := os.Create(geoos.EnvPath() + dir + filename)
	defer file.Close()
	if err != nil {
		log.Println(err)
	} else {
		defer file.Close()

		if err := geoencoding.WriteGeoJSON(file, geojson.GeometryToFeatureCollection(geom), geoencoding.GeoJSON); err != nil {
			log.Println(err)
		}
	}
}

// WriteMatrix Write the matrix object to a file in geojson format.
func WriteMatrix(filename string, m matrix.Steric) {

	file, err := os.Create(geoos.EnvPath() + dir + filename)
	defer file.Close()
	if err != nil {
		log.Println(err)
	} else {
		defer file.Close()

		if err := geoencoding.WriteGeoJSON(file, geojson.GeometryToFeatureCollection(space.TransGeometry(m)), geoencoding.GeoJSON); err != nil {
			log.Println(err)
		}
	}
}
