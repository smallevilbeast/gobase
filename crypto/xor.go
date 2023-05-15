package crypto

func XorCrypto(data []byte, key []byte) []byte {
	result := make([]byte, len(data))
	keyLen := len(key)

	for i, b := range data {
		result[i] = b ^ key[i%keyLen]
	}

	return result
}
