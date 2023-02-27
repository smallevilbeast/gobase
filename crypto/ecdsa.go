package crypto

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type EcdsaPublicKey struct {
	publicKey *ecdsa.PublicKey
}

func NewEcdsaPublicKeyFromPEM(pemKey string) (*EcdsaPublicKey, error) {
	// PEM 格式的公钥
	pemPublicKey := []byte(pemKey)
	// 解码 PEM 格式的公钥
	block, _ := pem.Decode(pemPublicKey)
	if block == nil {
		return nil, errors.New("failed to decode PEM public key")
	}

	// 解析 ASN.1 DER 编码的公钥
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 将公钥转换为 ecdsa.PublicKey 类型
	ecdsaPubKey, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	return &EcdsaPublicKey{
		publicKey: ecdsaPubKey,
	}, nil
}

func (key *EcdsaPublicKey) Verify(hash []byte, sig []byte) bool {
	return ecdsa.VerifyASN1(key.publicKey, hash, sig)
}
