package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"math/big"
)

// RSAPublicEncrypt 使用 RSA 公钥加密数据
func RSAPublicEncrypt(keyE string, keyN string, plaintext []byte) ([]byte, error) {
	// 从 keyN 和 keyE 构建 RSA PublicKey
	newKeyN, _ := new(big.Int).SetString(keyN, 16)
	newKeyE, _ := new(big.Int).SetString(keyE, 16)
	pubKey := &rsa.PublicKey{
		N: newKeyN,
		E: int(newKeyE.Int64()),
	}

	// 计算最大加密长度
	maxLen := pubKey.Size() - 11

	// 对 plaintext 分块加密
	var ciphertext []byte
	for len(plaintext) > 0 {
		var blockSize int
		if len(plaintext) > maxLen {
			blockSize = maxLen
		} else {
			blockSize = len(plaintext)
		}

		block, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plaintext[:blockSize])
		if err != nil {
			return nil, err
		}

		ciphertext = append(ciphertext, block...)
		plaintext = plaintext[blockSize:]
	}

	return ciphertext, nil
}
