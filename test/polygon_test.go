package test

import (
	"github.com/spatial-go/geos/base"
	"testing"
)
const bufferPoly = `POLYGON (( 10 130, 50 190, 110 190, 140 150, 150 80, 100 10, 20 40, 10 130 ),( 70 40, 100 50, 120 80, 80 110, 50 90, 70 40 ))`
const multiPoly = `POLYGON ((109.918212890625 32.94414888814148, 110.7366943359375 32.29177633471201, 110.7147216796875 33.00866349457558, 109.918212890625 32.94414888814148))
POLYGON ((110.9234619140625 32.97180377635759, 111.0223388671875 32.49586350791503, 111.0498046875 32.97180377635759, 110.9234619140625 32.97180377635759))`
const lineString =`LINESTRING(100 150,50 60, 70 80, 160 170)`
func TestBuffer(t *testing.T){

	geometry, _ := base.UnmarshalString(multiPoly)
	p := geometry.(base.Polygon)
	f, _ := p.Area()
	t.Log(f)

}

func TestBoundary(t *testing.T)  {
	geometry, _ := base.UnmarshalString(bufferPoly)
	p := geometry.(base.Polygon)
	boundary, e := p.Boundary()
	s := base.MarshalString(*boundary)
	t.Logf("测试 %s",s)

	if e!=nil{
		t.Error(e)
	}
	t.Log(boundary)
}