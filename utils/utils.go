// Package utils A functions of utils.
package utils

import (
	"bytes"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	// GBK ...
	GBK string = "GBK"
	// UTF8 ...
	UTF8 string = "UTF8"
	// UNKNOWN ...
	UNKNOWN string = "UNKNOWN"
)

// GetStringEncoding determine string encoding, UTF8 or GBK or UNKNOWN
func GetStringEncoding(dataStr string) string {
	// filter special characters
	dataStr = strings.ReplaceAll(dataStr, "·", "")
	data := []byte(dataStr)
	if IsUTF8(data) {
		return UTF8
	} else if IsGBK(data) {
		return GBK
	} else {
		return UNKNOWN
	}
}

// IsGBK determine GBK encoding
func IsGBK(data []byte) bool {
	length := len(data)
	var i = 0
	for i < length-1 {
		if data[i] <= 0x7f {
			i++
			continue
		} else {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0x7f {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func preNUm(data byte) int {
	var mask byte = 0x80
	var num int
	for i := 0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}

// IsUTF8 determine UTF8 encoding
func IsUTF8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			i++
			for j := 0; j < num-1; j++ {
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			return false
		}
	}
	return true
}

// UTF82GBK convert UTF8 to GBK
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}
