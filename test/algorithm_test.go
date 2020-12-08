package test

import (
	"testing"
)

func TestInterpolate(t *testing.T) {

	//wkt := coder.NewWKT()
	//g := wkt.FromWKTStr("LINESTRING(0 0, 1 1)")
	//defer wkt.Destroy()
	//
	//algorithm := NewAlgorithm()
	//defer algorithm.Destroy()
	//geometry, e := algorithm.Interpolate(g, 10)
	//if e != nil {
	//	t.Error(e.Error())
	//} else {
	//	t.Logf("%p", &geometry)
	//}

	//c:= geo.InitGeosContext()
	//fmt.Printf("%p",c)
	//geo.FinishGeosContext(c)
	//fmt.Printf("%p",c)

}

//func TestGeometryProject(t *testing.T) {
//	wkt := geo.NewWKT()
//	defer wkt.Destroy()
//	ls := geo.FromWKTStr(wkt.Handler,"LINESTRING(0 0, 1 1)")
//	pt := geo.FromWKTStr(wkt.Handler,"POINT(0 1)")
//	fmt.Printf("%p",ls)
//	fmt.Printf("%p",pt)
//	a := algorithm.NewAlgorithm()
//	defer a.Destroy()
//	project := a.Project(ls, pt)
//
//	expected := 0.7071067811865476
//	if expected != project {
//		t.Errorf("Geometry.Project(): want %v, got %v", expected, project)
//	}
//}
