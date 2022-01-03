package encryption

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRandomIVKey generate random IV value
func GenerateRandomIVKey(blockSize int) (bIv string, err error) {
	bytes := make([]byte, blockSize)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
