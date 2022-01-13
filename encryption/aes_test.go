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
