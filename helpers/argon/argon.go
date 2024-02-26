package argon

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

type Argon2idHash struct {
	// time represents the number of
	// passed over the specified memory.
	time uint32
	// cpu memory to be used.
	memory uint32
	// threads for parallelism aspect
	// of the algorithm.
	threads uint8
	// keyLen of the generate hash key.
	keyLen uint32
}

func (a *Argon2idHash) GenerateHash(password, salt []byte) []byte {
	return argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)
}

// Compare generated hash with store hash.
func (a *Argon2idHash) Compare(hash, password, salt []byte) error {
	// Generate hash for comparison.
	hashedPassword := a.GenerateHash(password, salt)
	// Compare the generated hash with the stored hash.
	// If they don't match return error.
	if !bytes.Equal(hash, hashedPassword) {
		return errors.New("hash doesn't match")
	}
	return nil
}

func NewArgon2idHash(time, memory uint32, threads uint8, keyLen uint32) *Argon2idHash {
	return &Argon2idHash{
		time:    time,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}
}

func RandomSecret(length uint8) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func MustGenerateRdmStr(length uint8) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)[:length]
}
