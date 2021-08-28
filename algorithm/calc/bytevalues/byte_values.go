// Package bytevalues ead and write primitive datatypes from/to byte
package bytevalues

import "math"

// Byte order: big endian or little endian
const (
	BigEndian    = 0
	LittleEndian = 1
)

// GetInt32 Returns int32  read primitive datatypes from/to byte
// sequences, allowing the byte order to be specified
func GetInt32(buf []byte, byteOrder int) uint32 {
	if byteOrder == BigEndian {
		return uint32((buf[0]&0xff))<<24 | uint32((buf[1]&0xff))<<16 | uint32((buf[2]&0xff))<<8 | uint32(buf[3]&0xff)
	}
	return uint32((buf[3]&0xff))<<24 | uint32((buf[2]&0xff))<<16 | uint32((buf[1]&0xff))<<8 | uint32(buf[0]&0xff)

}

// PutInt32  write primitive datatypes from/to byte
// sequences, allowing the byte order to be specified
func PutInt32(buf []byte, intValue int32, byteOrder int) {
	if byteOrder == BigEndian {
		buf[0] = byte(intValue >> 24)
		buf[1] = byte(intValue >> 16)
		buf[2] = byte(intValue >> 8)
		buf[3] = byte(intValue)
	} else { // LITTLE_ENDIAN
		buf[0] = byte(intValue)
		buf[1] = byte(intValue >> 8)
		buf[2] = byte(intValue >> 16)
		buf[3] = byte(intValue >> 24)
	}
}

// GetInt64 Returns int64  read primitive datatypes from/to byte
// sequences, allowing the byte order to be specified
func GetInt64(buf []byte, byteOrder int) uint64 {
	if byteOrder == BigEndian {
		return uint64((buf[0]&0xff))<<56 | uint64(buf[1]&0xff)<<48 |
			uint64(buf[2]&0xff)<<40 | uint64(buf[3]&0xff)<<32 |
			uint64(buf[4]&0xff)<<24 | uint64(buf[5]&0xff)<<16 |
			uint64(buf[6]&0xff)<<8 | uint64(buf[7]&0xff)
	}
	return uint64(buf[7]&0xff)<<56 | uint64(buf[6]&0xff)<<48 |
		uint64(buf[5]&0xff)<<40 | uint64(buf[4]&0xff)<<32 |
		uint64(buf[3]&0xff)<<24 | uint64(buf[2]&0xff)<<16 |
		uint64(buf[1]&0xff)<<8 | uint64(buf[0]&0xff)
}

// PutInt64  write primitive datatypes from/to byte
// sequences, allowing the byte order to be specified
func PutInt64(buf []byte, intValue int64, byteOrder int) {
	if byteOrder == BigEndian {
		buf[0] = byte(intValue >> 56)
		buf[1] = byte(intValue >> 48)
		buf[2] = byte(intValue >> 40)
		buf[3] = byte(intValue >> 32)
		buf[4] = byte(intValue >> 24)
		buf[5] = byte(intValue >> 16)
		buf[6] = byte(intValue >> 8)
		buf[7] = byte(intValue)
	} else { // LITTLE_ENDIAN
		buf[0] = byte(intValue)
		buf[1] = byte(intValue >> 8)
		buf[2] = byte(intValue >> 16)
		buf[3] = byte(intValue >> 24)
		buf[4] = byte(intValue >> 32)
		buf[5] = byte(intValue >> 40)
		buf[6] = byte(intValue >> 48)
		buf[7] = byte(intValue >> 56)
	}
}

// GetFloat32 Returns float32  read primitive datatypes from/to byte
// sequences, allowing the byte order to be specified
func GetFloat32(buf []byte, byteOrder int) float32 {
	longVal := GetInt32(buf, byteOrder)
	return math.Float32frombits(longVal)
}

// PutFloat32  write primitive datatypes from/to byte
// sequences, allowing the byte order to be specified
func PutFloat32(buf []byte, floatValue float32, byteOrder int) {
	longVal := math.Float32bits(floatValue)
	PutInt32(buf, int32(longVal), byteOrder)
}

// GetFloat64 Returns float64  read primitive datatypes from/to byte
// sequences, allowing the byte order to be specified
func GetFloat64(buf []byte, byteOrder int) float64 {
	longVal := GetInt64(buf, byteOrder)
	return math.Float64frombits(longVal)
}

// PutFloat64  write primitive datatypes from/to byte
// sequences, allowing the byte order to be specified
func PutFloat64(buf []byte, floatValue float64, byteOrder int) {
	longVal := math.Float64bits(floatValue)
	PutInt64(buf, int64(longVal), byteOrder)
}
