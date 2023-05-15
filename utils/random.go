package utils

import (
	"crypto/rand"
)

func RandomBytes(size uint32) []byte {
	key := make([]byte, size)
	rand.Read(key)
	return key
}
