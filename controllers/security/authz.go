package security

import (
	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/errors"
	"github.com/MaSTeR2W/SADEEM/models/security"
	"github.com/labstack/echo/v4"
)

func Authz(userType string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Get("authData").(*security.AuthData).UserType != userType {
				return &errors.HTTP403Err{
					Message: notAuthzErr(c.QueryParam("lang")),
				}
			}
			return next(c)
		}
	}
}

func notAuthzErr(lang string) string {
	if lang == "ar" {
		return "ليس لديك إذن الولوج"
	}
	return "You do not have access permission"
}
