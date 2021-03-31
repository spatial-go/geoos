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

// GetStringCoding
// 需要说明的是，isGBK()是通过双字节是否落在gbk的编码范围内实现的，
// 而utf-8编码格式的每个字节都是落在gbk的编码范围内，
// 所以只有先调用isUtf8()先判断不是utf-8编码，再调用isGBK()才有意义
func GetStringCoding(dataStr string) string {
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

// IsGBK 判断是否是GBK编码
func IsGBK(data []byte) bool {
	length := len(data)
	var i = 0
	for i < length-1 {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
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
	var num int = 0
	//8bit中首个0bit前有多少个1bits
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

// IsUTF8 判断是否是UTF8编码
func IsUTF8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
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
