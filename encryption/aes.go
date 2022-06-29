package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/sirupsen/logrus"
)

// AESCryptor aes cryptor
type AESCryptor struct {
	encryptionKey []byte
	iv            string
	blockSize     int
}

// NewAESCryptor create new AESCryptor.
// The key will be hashed using sha256 and will be used as the private key
func NewAESCryptor(key string, iv string, blockSize int) *AESCryptor {
	return &AESCryptor{
		encryptionKey: generateEncryptionKey(key),
		iv:            iv,
		blockSize:     blockSize,
	}
}

// Encrypt encrypt plain text using CBC
func (c *AESCryptor) Encrypt(plainText string) (encryptedText string, err error) {
	ivKey, err := hex.DecodeString(c.iv)
	if err != nil {
		logrus.Error(err)
		return
	}

	bPlaintext := c.pkcs5Padding([]byte(plainText), c.blockSize, len(plainText))
	block, err := aes.NewCipher(c.encryptionKey)
	if err != nil {
		logrus.Error(err)
		return
	}

	ciphertext := make([]byte, len(bPlaintext))

	mode := cipher.NewCBCEncrypter(block, ivKey)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypt cipherText using CBC
func (c *AESCryptor) Decrypt(cipherText string) (plainText string, err error) {
	ivKey, err := c.generateIVKey(c.iv)
	if err != nil {
		return
	}

	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(c.encryptionKey)
	if err != nil {
		return
	}

	mode := cipher.NewCBCDecrypter(block, ivKey)
	mode.CryptBlocks(cipherTextDecoded, cipherTextDecoded)

	return string(c.pkcs5Unpadding(cipherTextDecoded)), nil
}

func (c *AESCryptor) generateIVKey(iv string) (bIv []byte, err error) {
	if len(iv) > 0 {
		ivKey, err := hex.DecodeString(iv)
		if err != nil {
			return nil, fmt.Errorf("unable to hex decode iv")
		}
		return ivKey, nil
	}

	ivKey, err := GenerateRandomIVKey(c.blockSize)
	if err != nil {
		return nil, fmt.Errorf("unable to generate random iv key")
	}

	return hex.DecodeString(ivKey)
}

func (c *AESCryptor) pkcs5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func (c *AESCryptor) pkcs5Unpadding(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}

	length := len(src)
	unpadding := int(src[length-1])
	cutLen := (length - unpadding)
	// check boundaries
	if cutLen < 0 || cutLen > length {
		return src
	}

	return src[:cutLen]
}

func generateEncryptionKey(key string) []byte {
	hash := sha256.New()
	hash.Write([]byte(key))
	return hash.Sum(nil)
}
