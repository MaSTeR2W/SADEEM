package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"os"

	"github.com/MaSTeR2W/SADEEM/envVars"
)

var GCM cipher.AEAD = func() cipher.AEAD {
	// ==========================
	// because Golang is a stupid
	envVars.Read()
	// =========================

	key, err := base64.StdEncoding.DecodeString(os.Getenv("AES_KEY"))

	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err)
	}

	return gcm
}()

func Encrypt(text string) ([]byte, error) {

	var nonce = make([]byte, GCM.NonceSize())
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	cipherBytes := GCM.Seal(nil, nonce, []byte(text), nil)

	return append(nonce, cipherBytes...), nil
}

func Decrypt(nonceWithCipherBytes []byte) (string, error) {

	var nonce, cipherBytes = nonceWithCipherBytes[:GCM.NonceSize()], nonceWithCipherBytes[GCM.NonceSize():]

	plainBytes, err := GCM.Open(nil, nonce, cipherBytes, nil)

	if err != nil {
		return "", err
	}
	return string(plainBytes), nil
}

func DecryptMany(multi_nonceWithCipherBytes [][]byte) ([]string, error) {
	var plainTexts = []string{}
	for _, nonceWithCipherBytes := range multi_nonceWithCipherBytes {

		var nonce, cipherBytes = nonceWithCipherBytes[:GCM.NonceSize()], nonceWithCipherBytes[GCM.NonceSize():]

		plainBytes, err := GCM.Open(nil, nonce, cipherBytes, nil)

		if err != nil {
			return nil, err
		}
		plainTexts = append(plainTexts, string(plainBytes))
	}

	return plainTexts, nil
}
