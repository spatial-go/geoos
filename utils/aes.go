package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AesEncryption ...
func AesEncryption(key, iv, plainText []byte) ([]byte, error) {

	block, err := aes.NewCipher(key[:16])

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

// AesDecryption ...
func AesDecryption(key, iv, cipherText []byte) ([]byte, error) {

	block, err := aes.NewCipher(key[:16])

	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// PKCS5Padding ...
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS5UnPadding ...
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	return src[:(length - int(src[length-1]))]
}
