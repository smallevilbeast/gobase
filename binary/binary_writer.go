package binary

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strings"
)

type BinaryWriter struct {
	writer *bytes.Buffer
}

func NewBinaryWriter() *BinaryWriter {
	return &BinaryWriter{
		writer: bytes.NewBuffer([]byte{}),
	}
}

func (bw *BinaryWriter) writeNumber(data any, littleEndian ...bool) error {
	var order binary.ByteOrder
	if len(littleEndian) > 0 && littleEndian[0] {
		order = binary.LittleEndian
	} else {
		order = binary.BigEndian
	}
	return binary.Write(bw.writer, order, data)
}

func (bw *BinaryWriter) WriteUint8(value uint8) error {
	return bw.writeNumber(&value, true)
}

func (bw *BinaryWriter) WriteUint16(value uint16, littleEndian ...bool) error {
	return bw.writeNumber(&value, littleEndian...)
}

func (bw *BinaryWriter) WriteUint32(value uint32, littleEndian ...bool) error {
	return bw.writeNumber(&value, littleEndian...)
}

func (bw *BinaryWriter) WriteUint64(value uint64, littleEndian ...bool) error {
	return bw.writeNumber(&value, littleEndian...)
}

func (bw *BinaryWriter) WriteUvarint32(value uint32) error {
	buf := make([]byte, binary.MaxVarintLen32)
	n := binary.PutUvarint(buf, uint64(value))
	_, err := bw.writer.Write(buf[:n])
	return err
}

func (bw *BinaryWriter) WriteUvarint64(value uint64) error {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, value)
	_, err := bw.writer.Write(buf[:n])
	return err
}

func (bw *BinaryWriter) WriteBytes(value []byte) error {
	_, err := bw.writer.Write(value)
	return err
}

func (bw *BinaryWriter) WriteHex(hexValue string) error {
	hexValue = strings.Replace(hexValue, " ", "", -1)
	value, err := hex.DecodeString(hexValue)
	if err != nil {
		return err
	}
	_, err = bw.writer.Write(value)
	return err
}

func (bw *BinaryWriter) Size() int {
	return bw.writer.Len()
}

func (bw *BinaryWriter) ToBytes() []byte {
	return bw.writer.Bytes()
}
