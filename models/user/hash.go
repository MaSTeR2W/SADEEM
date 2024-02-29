package user

import (
	"bytes"

	"github.com/MaSTeR2W/SADEEM/helpers/argon"
)

var arg = argon.NewArgon2idHash(3, 12288, 1, 256)

func GenerateSaltAndHashPassword(password string) ([]byte, []byte, error) {

	salt, err := argon.RandomSecret(128)
	if err != nil {
		return nil, nil, err
	}

	return arg.GenerateHash([]byte(password), salt), salt, nil
}

func ComparePassword(enteredPass string, salt []byte, curPass []byte) bool {
	hashedPass := arg.GenerateHash([]byte(enteredPass), salt)

	return bytes.Equal(hashedPass, curPass)
}
