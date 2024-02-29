package security

import (
	"crypto/ecdsa"
	"os"
	"time"

	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/errors"
	"github.com/MaSTeR2W/SADEEM/envVars"
	_ecdsa "github.com/MaSTeR2W/SADEEM/helpers/ecdsa"

	"github.com/golang-jwt/jwt/v5"
)

var priKey, pubKey = func() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	envVars.Read()
	ecc := _ecdsa.NewWithPriKey(os.Getenv("EC_PRIVATE_KEY"))

	return ecc.PriKey, ecc.PubKey
}()

func CreateToken(userId int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString(priKey)

	if err != nil {
		return "", err
	}

	return tokenStr, err
}

func VerifyToken(tokenStr string, lang string) (int, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		},
	)

	if err != nil || !token.Valid {
		var msg string
		if lang == "ar" {
			msg = "الرمز غير صالح"
		} else {
			msg = "Invalid token"
		}
		return 0, &errors.HTTP401Err{
			Message: msg,
		}
	}

	return int(token.Claims.(jwt.MapClaims)["userId"].(float64)), nil
}
