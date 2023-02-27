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

func main() {
	testBinaryWriter()
	testBinaryReader()

	testAesCbcCrypto()
}
