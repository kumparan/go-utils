package encryption

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAESCryptor_Encrypt(t *testing.T) {
	key := "secret"
	plain := "foo bar"
	require.NotEqual(t, 0, len(plain)%aes.BlockSize)

	iv, err := GenerateRandomIVKey(aes.BlockSize)
	require.NoError(t, err)
	cryptor := NewAESCryptor(key, iv, aes.BlockSize)
	enc, err := cryptor.Encrypt(plain)
	require.NoError(t, err)
	require.NotEmpty(t, enc)
	require.NotEqual(t, plain, enc)
	require.NotEqual(t, len(plain), len(enc))
	require.Equal(t, 0, len(enc)%aes.BlockSize)

	dec, err := cryptor.Decrypt(enc)
	require.NoError(t, err)
	require.Equal(t, plain, dec)
}

func TestAESCryptor_PKCSUnpadding(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		cryptor := &AESCryptor{}
		res := cryptor.pkcs5Unpadding(nil)
		require.Nil(t, res)
	})

	emails := []string{
		`kumkum@email.com`,
		`kumkum@email.com   `,
		"",
	}

	iv, err := GenerateRandomIVKey(aes.BlockSize)
	require.NoError(t, err)

	for _, email := range emails {
		cryptor := NewAESCryptor("foobar", iv, aes.BlockSize)
		enc, err := cryptor.Encrypt(email)
		require.NoError(t, err)
		require.NotEqual(t, email, enc)

		decEmail, err := cryptor.Decrypt(enc)
		require.NoError(t, err)
		require.Equal(t, email, decEmail)
	}
}
