package crypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
)

func HmacSha256(key []byte, data []byte) []byte {
	hm := hmac.New(sha256.New, key)
	hm.Write(data)
	return hm.Sum(nil)
}

func HmacSha1(key []byte, data []byte) []byte {
	hm := hmac.New(sha1.New, key)
	hm.Write(data)
	return hm.Sum(nil)
}
