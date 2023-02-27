package crypto

import (
	"crypto/sha256"
	"golang.org/x/crypto/hkdf"
)

func HKDFExpendSHA256(key []byte, info []byte, length int) ([]byte, error) {
	newKey := make([]byte, length)
	_, err := hkdf.Expand(sha256.New, key, info).Read(newKey)
	if err != nil {
		return nil, err
	}
	return newKey, nil
}
