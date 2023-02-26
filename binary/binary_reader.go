package binary

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type BinaryReader struct {
	reader *bytes.Reader
}

func NewBinaryReader(buf []byte) *BinaryReader {
	return &BinaryReader{
		reader: bytes.NewReader(buf),
	}
}

func (br *BinaryReader) ReadBytes(size uint32) ([]byte, error) {
	buf := make([]byte, size)
	_, err := br.reader.Read(buf)
	return buf, err
}

func (br *BinaryReader) readNumber(data any, littleEndian ...bool) error {
	var order binary.ByteOrder
	if len(littleEndian) > 0 && littleEndian[0] {
		order = binary.LittleEndian
	} else {
		order = binary.BigEndian
	}
	err := binary.Read(br.reader, order, data)
	return err

}

func (br *BinaryReader) ReadUint8() (uint8, error) {
	var value uint8
	err := br.readNumber(&value, true)
	return value, err
}

func (br *BinaryReader) ReadUint16(littleEndian ...bool) (uint16, error) {
	var value uint16
	err := br.readNumber(&value, littleEndian...)
	return value, err
}

func (br *BinaryReader) ReadUint32(littleEndian ...bool) (uint32, error) {
	var value uint32
	err := br.readNumber(&value, littleEndian...)
	return value, err
}

func (br *BinaryReader) ReadUint64(littleEndian ...bool) (uint64, error) {
	var value uint64
	err := br.readNumber(&value, littleEndian...)
	return value, err
}

func (br *BinaryReader) ReadUvarint32() (uint32, error) {
	value, err := br.ReadUvarint64()
	if err != nil {
		return 0, err
	}

	if value>>32 != 0 {
		return 0, errors.New("is 64 bit")
	}
	return uint32(value), err
}

func (br *BinaryReader) ReadUvarint64() (uint64, error) {
	return binary.ReadUvarint(br.reader)
}

func (br *BinaryReader) Discard(size uint32) error {
	_, err := br.reader.Seek(int64(size), io.SeekCurrent)
	return err
}

func (br *BinaryReader) ReadToEnd() ([]byte, error) {
	size := br.reader.Len()
	if size <= 0 {
		return nil, errors.New("buffer is empty")
	}
	buf := make([]byte, br.reader.Len())
	_, err := br.reader.Read(buf)
	return buf, err
}

func (br *BinaryReader) Size() int64 {
	return br.reader.Size()
}

func (br *BinaryReader) Len() int {
	return br.reader.Len()
}
