//Package geocsv is a library for read csv file with geospatial data.
package geocsv

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spatial-go/geoos/geoencoding/geojson"
	"github.com/spatial-go/geoos/geoencoding/wkt"
	"github.com/spatial-go/geoos/space"
	"github.com/spatial-go/geoos/utils"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// default value of longitude and latitude
var defaultCoordValue = float64(-9999)

// GeoCSV a extension of the CSV with geospatial data
type GeoCSV struct {
	r       io.Reader
	headers []string
	rows    [][]string
	options Options
	coll    space.Collection
}

// Options an options of GeoCSV
type Options struct {
	Fields   []string
	XField   string
	YField   string
	WKTField string
}

// NewGeoCSV ...
func NewGeoCSV() (gc *GeoCSV) {
	gc = &GeoCSV{}
	return
}

func (gc *GeoCSV) readRecords() (err error) {
	if gc.r == nil {
		err = errors.New("file is nil")
		return
	}
	headerRead := false
	gbkDecoder := simplifiedchinese.GBK.NewDecoder()
	reader := csv.NewReader(gc.r)
	for {
		record, readErr := reader.Read()
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			err = readErr
			return
		}
		encodeValues := make([]string, 0, len(record))
		for _, value := range record {
			var encodeValue string
			coding := utils.GetStringEncoding(value)
			switch coding {
			case utils.UTF8:
				encodeValue = value
			case utils.GBK:
				encodingString, _ := gbkDecoder.Bytes([]byte(value))
				encodeValue = string(encodingString)
			default:
				if encodingString, decodeError := gbkDecoder.Bytes([]byte(value)); decodeError == nil {
					encodeValue = string(encodingString)
				} else {
					err = errors.New("file encoding is not supported")
					return
				}
			}
			encodeValue = strings.TrimSpace(encodeValue)
			// remove special characters, such as &#65279;
			encodeValue = strings.ReplaceAll(encodeValue, "\uFEFF", "")
			encodeValue = strings.TrimSpace(encodeValue)
			encodeValues = append(encodeValues, encodeValue)
		}
		if !headerRead {
			headerRead = true
			gc.headers = encodeValues
		} else {
			gc.rows = append(gc.rows, encodeValues)
		}
	}
	return
}

// Read read csv file with options
func Read(filePath string, options Options) (gc *GeoCSV, err error) {
	gc = NewGeoCSV()
	gc.options = options
	if file, fileerr := os.Open(filePath); fileerr == nil {
		gc.r = file
		defer file.Close()
		if err = gc.readRecords(); err != nil {
			return
		}
	}
	return
}

// ReadByte read csv file with options
func ReadByte(reader io.Reader, options Options) (gc *GeoCSV, err error) {
	gc = NewGeoCSV()
	gc.options = options
	gc.r = reader
	if err = gc.readRecords(); err != nil {
		return
	}
	return
}

// ToGeoJSON export geojson
func (gc *GeoCSV) ToGeoJSON() (features *geojson.FeatureCollection) {
	features = geojson.NewFeatureCollection()
	for _, row := range gc.rows {
		var (
			lng      = defaultCoordValue
			lat      = defaultCoordValue
			geometry *geojson.Geometry
		)
		properties := geojson.Properties{}

		for j, cell := range row {
			fieldName := gc.headers[j]
			if len(gc.options.WKTField) > 0 && fieldName == gc.options.WKTField {
				if wktGeometry, wktError := wkt.UnmarshalString(cell); wktError == nil {
					geometry = geojson.NewGeometry(wktGeometry)
				}
			} else if len(gc.options.XField) > 0 && fieldName == gc.options.XField {
				lng, _ = strconv.ParseFloat(cell, 64)
			} else if len(gc.options.YField) > 0 && fieldName == gc.options.YField {
				lat, _ = strconv.ParseFloat(cell, 64)
			}
			properties[fieldName] = cell
		}
		if geometry == nil && lng != defaultCoordValue && lat != defaultCoordValue {
			geometry = geojson.NewGeometry(space.Point{lng, lat})
		}
		if geometry != nil {
			feature := geojson.NewFeature(*geometry)
			feature.Properties = properties
			features.Features = append(features.Features, feature)
		}
	}
	return
}
