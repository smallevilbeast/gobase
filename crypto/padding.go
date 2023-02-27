package crypto

import (
	"bytes"
	"errors"
)

type CryptoPadding interface {
	Padding(plainText []byte, blockSize int) []byte
	UnPadding(cipherText []byte) ([]byte, error)
}

type PKCS5Padding struct{}
type PKCS7Padding struct{}

func (p *PKCS5Padding) Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainText, padtext...)
}

func (p *PKCS5Padding) UnPadding(cipherText []byte) ([]byte, error) {
	length := len(cipherText)
	unpadding := int(cipherText[length-1])
	if unpadding > length {
		return nil, errors.New("padding error")
	}

	return cipherText[:(length - unpadding)], nil
}

func (p *PKCS7Padding) Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	padtext := make([]byte, padding)
	for i := 0; i < padding; i++ {
		padtext[i] = byte(padding)
	}
	return append(plainText, padtext...)
}

func (p *PKCS7Padding) UnPadding(cipherText []byte) ([]byte, error) {
	length := len(cipherText)
	unpadding := int(cipherText[length-1])

	if unpadding > length {
		return nil, errors.New("padding error")
	}

	return cipherText[:(length - unpadding)], nil
}
