package binary

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strings"
)

type BinaryWriter struct {
	w bytes.Buffer
}

func NewBinaryWriter() *BinaryWriter {
	return &BinaryWriter{}
}

func (bw *BinaryWriter) WriteUint8(value uint8) error {
	buf := []byte{value}
	_, err := bw.w.Write(buf)
	return err
}

func (bw *BinaryWriter) WriteUint16(value uint16, littleEndian ...bool) error {
	buf := make([]byte, 2)
	if len(littleEndian) > 0 && littleEndian[0] {
		binary.LittleEndian.PutUint16(buf, value)
	} else {
		binary.BigEndian.PutUint16(buf, value)
	}
	_, err := bw.w.Write(buf)
	return err
}

func (bw *BinaryWriter) WriteUint32(value uint32, littleEndian ...bool) error {
	buf := make([]byte, 4)
	if len(littleEndian) > 0 && littleEndian[0] {
		binary.LittleEndian.PutUint32(buf, value)
	} else {
		binary.BigEndian.PutUint32(buf, value)
	}
	_, err := bw.w.Write(buf)
	return err
}

func (bw *BinaryWriter) WriteUint64(value uint64, littleEndian ...bool) error {
	buf := make([]byte, 8)
	if len(littleEndian) > 0 && littleEndian[0] {
		binary.LittleEndian.PutUint64(buf, value)
	} else {
		binary.BigEndian.PutUint64(buf, value)
	}
	_, err := bw.w.Write(buf)
	return err
}

func (bw *BinaryWriter) WriteVarint32(value int32) error {
	buf := make([]byte, binary.MaxVarintLen32)
	n := binary.PutVarint(buf, int64(value))
	_, err := bw.w.Write(buf[:n])
	return err
}

func (bw *BinaryWriter) WriteVarint64(value int64) error {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, value)
	_, err := bw.w.Write(buf[:n])
	return err
}

func (bw *BinaryWriter) WriteBytes(value []byte) error {
	_, err := bw.w.Write(value)
	return err
}

func (bw *BinaryWriter) WriteHex(hexValue string) error {
	hexValue = strings.Replace(hexValue, " ", "", -1)
	value, err := hex.DecodeString(hexValue)
	if err != nil {
		return err
	}
	_, err = bw.w.Write(value)
	return err
}

func (bw *BinaryWriter) Size() int {
	return bw.w.Len()
}

func (bw *BinaryWriter) ToBytes() []byte {
	return bw.w.Bytes()
}
