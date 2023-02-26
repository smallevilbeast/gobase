package main

import (
	"encoding/hex"
	"fmt"

	"github.com/smallevilbeast/gobase/binary"
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
	uvarint, _ := r.ReadUvarint()
	end, _ := r.ReadToEnd()
	fmt.Printf("u8: 0x%x, u16: 0x%x, varint: %d, end: %s\n", u8, u16, uvarint, hex.EncodeToString(end))
}

func main() {
	testBinaryWriter()
	testBinaryReader()
}
