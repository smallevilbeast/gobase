package crypto

import (
	"crypto/aes"
	"errors"
	"github.com/andreburgaud/crypt2go/ecb"
)

type AesECB struct {
	padding CryptoPadding
}

func NewAesECBPKCS5() *AesECB {
	return &AesECB{
		padding: &PKCS5Padding{},
	}
}

func NewAesECBPKCS7() *AesECB {
	return &AesECB{
		padding: &PKCS7Padding{},
	}
}

func (crypto *AesECB) Encrypt(plainText []byte, key []byte) ([]byte, error) {
	return AesECBEncrypt(plainText, key, crypto.padding)
}

func (crypto *AesECB) Decrypt(cipherText []byte, key []byte) ([]byte, error) {
	return AesECBDecrypt(cipherText, key, crypto.padding)
}

func AesECBEncrypt(plainText []byte, key []byte, padding CryptoPadding) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plainText = padding.Padding(plainText, blockSize)

	cipherText := make([]byte, len(plainText))
	mode := ecb.NewECBEncrypter(block)
	mode.CryptBlocks(cipherText, plainText)

	return cipherText, nil

}

func AesECBDecrypt(cipherText []byte, key []byte, padding CryptoPadding) ([]byte, error) {
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

	mode := ecb.NewECBDecrypter(block)
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
