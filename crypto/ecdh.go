package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
)

type KDFType int

const (
	KDFTypeMD5 KDFType = iota
	KDFTypeSHA1
	KDFTypeSHA256
)

type EcdhKey struct {
	curve      elliptic.Curve
	privateKey *ecdsa.PrivateKey
}

// NewEcdhKeyP256 ECDH_NID_prime256v1 = 415
func NewEcdhKeyP256() *EcdhKey {
	ecdhKey := &EcdhKey{
		curve:      elliptic.P256(),
		privateKey: nil,
	}
	err := ecdhKey.GenerateKey()
	if err != nil {
		return nil
	}
	return ecdhKey
}

// NewEcdhKeyP224 ECDH_NID_secp224r1 = 713
func NewEcdhKeyP224() *EcdhKey {
	ecdhKey := &EcdhKey{
		curve:      elliptic.P224(),
		privateKey: nil,
	}
	err := ecdhKey.GenerateKey()
	if err != nil {
		return nil
	}
	return ecdhKey
}

func (ecdh *EcdhKey) GenerateKey() error {
	privateKey, err := ecdsa.GenerateKey(ecdh.curve, rand.Reader)
	if err != nil {
		return err
	}
	ecdh.privateKey = privateKey
	return nil
}

func (ecdh *EcdhKey) GetPublicKey() ([]byte, error) {
	if ecdh.privateKey == nil {
		return nil, errors.New("private key is nil")
	}
	return elliptic.Marshal(
		ecdh.privateKey.PublicKey.Curve,
		ecdh.privateKey.PublicKey.X,
		ecdh.privateKey.PublicKey.Y), nil
}

func (ecdh *EcdhKey) ComputeDh(peerPublicKey []byte, kdfType KDFType) ([]byte, error) {
	x, y := elliptic.Unmarshal(ecdh.curve, peerPublicKey)
	if x == nil {
		return nil, errors.New("unmarshal peer public key error")
	}
	r, _ := ecdh.curve.ScalarMult(x, y, ecdh.privateKey.D.Bytes())
	secretKey := r.Bytes()
	switch kdfType {
	case KDFTypeMD5:
		result := md5.Sum(secretKey)
		return result[:], nil
	case KDFTypeSHA1:
		result := sha1.Sum(secretKey)
		return result[:], nil
	case KDFTypeSHA256:
		result := sha256.Sum256(secretKey)
		return result[:], nil
	}
	return nil, errors.New("kdf type not support")
}
