// package graph ...

package graph

import "github.com/spatial-go/geoos/algorithm/matrix"

// Relate explain
const (
	RPP1 = iota
	RPP2

	RPL1
	RPL2
	RPL3

	RPA1
	RPA2
	RPA3

	RLL1
	RLL2
	RLL3
	RLL4
	RLL5
	RLL6
	RLL7
	RLL8
	RLL9
	RLL10
	RLL11
	RLL12
	RLL13
	RLL14
	RLL15
	RLL16
	RLL17
	RLL18
	RLL19
	RLL20
	RLL21
	RLL22
	RLL23
	RLL24
	RLL25
	RLL26
	RLL27
	RLL28
	RLL29
	RLL30
	RLL31
	RLL32
	RLL33
	RLL34
	RLL35

	RLA1
	RLA2
	RLA3
	RLA4
	RLA5
	RLA6
	RLA7
	RLA8
	RLA9
	RLA10
	RLA11
	RLA12
	RLA13
	RLA14
	RLA15
	RLA16
	RLA17
	RLA18
	RLA19
	RLA20
	RLA21
	RLA22
	RLA23
	RLA24
	RLA25
	RLA26
	RLA27
	RLA28
	RLA29
	RLA30
	RLA31

	RAA1
	RAA2
	RAA3
	RAA4
	RAA5
	RAA6
	RAA7
	RAA8
	RAA9
	RAA10
	RAA11
)

// RelateStrings array
var (
	RelateStrings = []string{
		// point point
		"FF0FFF0F2", "0FFFFFFF2",
		// point line
		"FF0FFF102", "0FFFFF102", "F0FFFF102",
		// point polygon
		"FF0FFF212", "0FFFFF212", "F0FFFF212",

		// line line
		// num 8
		"FF1FF0102", "0F1FF0102", "1F1FF0102", "F01FF0102", "F01FF0102",
		"001FF0102", "001FF01F2", "1F10F0102", "1F10FF102", "1FF0FF102",
		"F010FF102", "F010F0102", "F010F0102", "0010FF102", "0010F0102",
		"0010F0102", "1010FF102", "1010F0102", "1010FF102", "FF1F0F102",
		"FF1F00102", "0F1F0F102", "0F1F00102", "1F1F0F102", "1FFF0F102",
		"1F1F00102", "F01F00102", "001F00102", "1F100F102", "1FF00F102",
		"F0100F102", "00100F102", "10100F102", "FF10F0102", "0F10FF102",

		// line polygon
		// num 41
		"FF1FF0212", "1FF0FF212", "F01FF0212", "F11FF0212", "10F0FF212",
		"11F0F0212", "101FF0212", "1010F0212", "1010FF212", "111FF0212",
		"1110FF212", "1110F0212", "FF1F0F212", "FF1F00212", "1FF00F212",
		"1FFF0F212", "F11F0F212", "F11F00212", "10FF0F212", "101F00212",
		"101F00212", "10F00F212", "10100F212", "F1FF0F212", "F11F0F212",
		"F11F00212", "11FF0F212", "101F0F212", "1F1F00212", "11F00F212",
		"11100F212",

		//polygon polygon
		// num 72
		"FF2FF1212", "FF2F01212", "FF2F11212", "212101212", "212FF1FF2",
		"2FF1FF212", "2F2F01FF2", "2FF10F212", "2FFF1FFF2", "2FF11F212",
		"212F11FF2",
	}
)

// RelateStringsTransposeByRing line relate to ring relate
// Model definition: boundary of point is nil,   boundary of  line is boundary,two point
// boundary of  ring is  nil, boundary of  polygon is  ring
// interior is Except boundary
// exterior exterior boundary and interior
func RelateStringsTransposeByRing(rs string, inputType int) string {
	if inputType < 1 {
		return rs
	}
	rsb := []byte(rs)
	switch inputType {
	case 1: // A is ring
		rsb[3] = 'F'
		rsb[4] = 'F'
		rsb[5] = 'F'
	case 2: // B is ring
		rsb[1] = 'F'
		rsb[4] = 'F'
		rsb[7] = 'F'
	case 3: // A and B is ring
		rsb[1] = 'F'
		rsb[3] = 'F'
		rsb[4] = 'F'
		rsb[5] = 'F'
		rsb[7] = 'F'
	}
	return string(rsb)
}

// IMTransposeByRing line relate to ring relate
// Model definition: boundary of point is nil,   boundary of  line is boundary,two point
// boundary of  ring is  nil, boundary of  polygon is  ring
// interior is Except boundary
// exterior exterior boundary and interior
func IMTransposeByRing(im *matrix.IntersectionMatrix, inputType int) *matrix.IntersectionMatrix {
	if inputType < 1 {
		return im
	}
	switch inputType {
	case 1: // A is ring
		im.Set(1, 0, -1)
		im.Set(1, 1, -1)
		im.Set(1, 2, -1)
	case 2: // B is ring
		im.Set(0, 1, -1)
		im.Set(1, 1, -1)
		im.Set(2, 1, -1)
	case 3: // A and B is ring
		im.Set(1, 0, -1)
		im.Set(1, 1, -1)
		im.Set(1, 2, -1)
		im.Set(0, 1, -1)
		im.Set(2, 1, -1)
	}
	return im
}
