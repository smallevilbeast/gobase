package binary

import (
	"encoding/hex"
	"strings"
)

func FromHex(hexStr string) []byte {
	hexStr = strings.ReplaceAll(hexStr, " ", "")
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return []byte{}
	}
	return bytes
}
