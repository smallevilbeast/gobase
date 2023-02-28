package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type AesCBC struct {
	padding CryptoPadding
}

func NewAesCBCPKCS5() *AesCBC {
	return &AesCBC{
		padding: &PKCS5Padding{},
	}
}

func NewAesCBCPKCS7() *AesCBC {
	return &AesCBC{
		padding: &PKCS7Padding{},
	}
}

func (crypto *AesCBC) Encrypt(plainText []byte, key []byte, iv []byte) ([]byte, error) {
	return AesCBCEncrypt(plainText, key, iv, crypto.padding)
}

func (crypto *AesCBC) Decrypt(cipherText []byte, key []byte, iv []byte) ([]byte, error) {
	return AesCBCDecrypt(cipherText, key, iv, crypto.padding)
}

func AesCBCEncrypt(plainText []byte, key []byte, iv []byte, padding CryptoPadding) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plainText = padding.Padding(plainText, blockSize)

	cipherText := make([]byte, len(plainText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, plainText)

	return cipherText, nil

}

func AesCBCDecrypt(cipherText []byte, key []byte, iv []byte, padding CryptoPadding) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	cipherTextSize := len(cipherText)
	if cipherTextSize < blockSize {
		return nil, errors.New("ciphertext too short")
	}

	if cipherTextSize%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, cipherTextSize)

	// 解密密文
	mode.CryptBlocks(plainText, cipherText)

	// 使用指定的填充方式进行去填充
	plainText, err = padding.UnPadding(plainText)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
