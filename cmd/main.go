package main

import (
	"encoding/hex"
	"fmt"

	"github.com/smallevilbeast/gobase/binary"
	"github.com/smallevilbeast/gobase/crypto"
)

func testBinaryWriter() {
	w := binary.NewBinaryWriter()
	w.WriteUint8(0xAB)
	w.WriteUint16(0)
	w.WriteUvarint32(197)
	w.WriteHex("AB 1F 04 20")
	fmt.Println(hex.EncodeToString(w.ToBytes()))
	fmt.Println(w.Size())
}

func testBinaryReader() {
	buf, _ := hex.DecodeString("010203c501c501ff01")
	r := binary.NewBinaryReader(buf)
	u8, _ := r.ReadUint8()
	u16, _ := r.ReadUint16()
	uvarint, _ := r.ReadUvarint32()
	end, _ := r.ReadToEnd()
	fmt.Printf("u8: 0x%x, u16: 0x%x, varint: %d, end: %s\n", u8, u16, uvarint, hex.EncodeToString(end))
}

func testAesCbcCrypto() {
	aes := crypto.NewAesCBCPKCS5()
	key := []byte("0123456789abcdef")
	plainText := []byte("0123456789abcdef")
	cipherText, _ := aes.Encrypt(plainText, key, key)
	fmt.Printf("cipherText: %s\n", hex.EncodeToString(cipherText))

	newPlainText, _ := aes.Decrypt(cipherText, key, key)
	fmt.Printf("plainText: %s\n", string(newPlainText))
}

func testEcdhKey() {
	ecdh := crypto.NewEcdhKeyP256()
	key, _ := ecdh.GetPublicKey()
	fmt.Println(hex.EncodeToString(key))
	peerKey, _ := hex.DecodeString("04ce7827f9014801097e7d86761f97cf74ba5ceeca8871ee35008af981ca199a10d08638bc28ee7d311917bbb934a631a40233b4011ad13f895c689ea46edb0298")
	out, _ := ecdh.ComputeDh(peerKey, crypto.KDFTypeSHA256)
	fmt.Println(hex.EncodeToString(out))
}

func testAesEcbEncrypt() {
	pt := []byte("Some11111111111111111111111111111111111")
	key := []byte("0123456789abcdef")
	aes := crypto.NewAesECBPKCS7()
	cipherText, err := aes.Encrypt(pt, key)
	fmt.Println(err)
	fmt.Println(hex.EncodeToString(cipherText))
}
func main() {
	testAesEcbEncrypt()
}
