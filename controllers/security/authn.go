package security

import (
	"strings"

	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/errors"
	"github.com/MaSTeR2W/SADEEM/models/security"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/labstack/echo/v4"
)

func Authn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var lang = c.QueryParam("lang")
		var authHeader = c.Request().Header.Get("Authorization")

		var tokenStr, found = strings.CutPrefix(authHeader, "Bearer ")

		if !found {
			return &errors.HTTP401Err{
				Message: missingTokenErr(lang),
			}
		}

		var userId, err = security.VerifyToken(tokenStr, lang)

		if err != nil {
			return err
		}

		var authData *security.AuthData

		authData, err = pgHprs.StmtQueryxAndStructScan[security.AuthData](security.GetAuthData, userId)

		if err != nil {
			return err
		}

		if authData.UserId == 0 {
			return &errors.HTTP401Err{
				Message: missingUserErr(lang),
			}
		}

		c.Set("authData", authData)

		return next(c)
	}
}

func missingTokenErr(lang string) string {
	if lang == "ar" {
		return "الرمز مفقود"
	}
	return "Token is missing"
}

func missingUserErr(lang string) string {
	if lang == "ar" {
		return "المستخدم الذي ينتمي إليه الرمز لم يعد موجوداً"
	}
	return "The user to which the token belongs no longer exists"
}
