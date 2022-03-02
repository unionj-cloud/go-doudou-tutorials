package lib_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/internal/lib"
	"os"
	"testing"
)

func TestGenerateRandomKey(t *testing.T) {
	numbytes := 32
	key := lib.NewRandomKey()

	assert.True(t, len(key) == numbytes, "must be 32 bytes")
	assert.True(t, len(fmt.Sprintf("%x", key)) == numbytes*2, "must be a 64 characters string")
	assert.False(t, bytes.Equal(key, make([]byte, numbytes)), "must be different than zeroes")
}

func TestEncryptDecrypt(t *testing.T) {

	secretKey := lib.NewRandomKey()

	ps := fmt.Sprintf("%x", secretKey)
	// set PLATFORM_SECRET env variable
	os.Setenv("PLATFORM_SECRET", ps)

	appSecretKey := lib.NewRandomKey()
	fmt.Println(hex.EncodeToString(appSecretKey))
	fmt.Println(fmt.Sprintf("%x", appSecretKey))

	encryptedAppSecret, err := lib.Encrypt(appSecretKey)

	if err != nil {
		t.Fatal(err)
	}

	decryptedAppSecret, err := lib.Decrypt(encryptedAppSecret)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%x", encryptedAppSecret))
	fmt.Println(fmt.Sprintf("%x", decryptedAppSecret))

	assert.True(t, bytes.Equal(appSecretKey, decryptedAppSecret), "must decrypt to the same value")
}

func TestJwtToken(t *testing.T) {
	token, err := lib.JwtToken("secret", jwt.MapClaims{
		"clientId": "wordcloud-bff",
		"userId":   0,
	})
	require.NoError(t, err)
	fmt.Println(token)
}
