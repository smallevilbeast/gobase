package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

func AesGCMEncrypt(plainText []byte, key []byte, nonce []byte, aad []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	dst := gcm.Seal(nil, nonce, plainText, aad)
	return dst, nil
}

func AesGCMDecrypt(cipherText []byte, key []byte, nonce []byte, aad []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm.Open(nil, nonce, cipherText, aad)
}
