package encryption

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAESCryptor_Encrypt(t *testing.T) {
	key := "secret"
	plain := "foo bar"

	iv, err := GenerateRandomIVKey(aes.BlockSize)
	require.NoError(t, err)
	cryptor := NewAESCryptor(key, iv, aes.BlockSize)
	enc, err := cryptor.Encrypt(plain)
	require.NoError(t, err)
	require.NotEmpty(t, enc)
	require.NotEqual(t, plain, enc)

	dec, err := cryptor.Decrypt(enc)
	require.NoError(t, err)
	require.Equal(t, plain, dec)

}
