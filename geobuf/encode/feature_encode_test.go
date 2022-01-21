package encode

import (
	"github.com/spatial-go/geoos/geobuf/decode"
	"github.com/spatial-go/geoos/geojson"
	"testing"
)

func TestEncodeFeature(t *testing.T) {
	rawJSON := `
    {
      "type": "Feature",
      "properties": {},
      "geometry": {
        "type": "Polygon",
        "coordinates": [
          [
            [
              67.92629420757294,
              57.12477998587717
            ],
            [
              67.92551100254059,
              57.124471326735076
            ],
            [
              67.92604207992552,
              57.124279141441335
            ],
            [
              67.9270076751709,
              57.12427622953526
            ],
            [
              67.92707204818726,
              57.12453538827802
            ],
            [
              67.92629420757294,
              57.12477998587717
            ]
          ]
        ]
      }
    }
`

	f, err := geojson.UnmarshalFeature([]byte(rawJSON))
	if err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	feature := Encode(f)
	t.Log(feature)
	fe := decode.Decode(feature)
	t.Log(fe)
}
