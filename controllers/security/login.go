package security

import (
	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/errors"
	"github.com/MaSTeR2W/SADEEM/models/security"
	"github.com/MaSTeR2W/SADEEM/models/user"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	var body = map[string]any{}
	var lang = c.QueryParam("lang")
	c.Bind(&body)

	var err error

	if err = security.LoginValidator.Validate(body, lang); err != nil {
		return err
	}

	var userId int
	var password []byte
	var salt []byte

	err = pgHprs.QueryxAndScan(
		"SELECT user_id, password, salt FROM users WHERE email=$1",
		[]any{body["email"]},
		&userId,
		&password,
		&salt,
	)

	if err != nil {
		return err
	}

	if userId == 0 {
		return &errors.HTTP404Err{
			Message: incorrectPasswordOrEmail(lang),
		}
	}

	if !user.ComparePassword(body["password"].(string), salt, password) {
		return &errors.HTTP404Err{
			Message: incorrectPasswordOrEmail(lang),
		}
	}

	var token string
	token, err = security.CreateToken(userId)

	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{
		"token": token,
	})

}

func incorrectPasswordOrEmail(lang string) string {
	if lang == "ar" {
		return "كلمة المرور أو البريد الإلكتروني غير صحيح"
	}

	return "Incorrect Password or email"
}
