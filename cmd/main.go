package main

import (
	"encoding/hex"
	"fmt"

	"github.com/smallevilbeast/gobase/binary"
)

func main() {
	w := binary.NewBinaryWriter()
	w.WriteUint8(0xAB)
	w.WriteUint16(0)
	w.WriteVarint32(197)
	w.WriteHex("AB 1F 04 20")
	fmt.Println(hex.EncodeToString(w.ToBytes()))
	fmt.Println(w.Size())
}
