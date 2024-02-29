package user

import (
	"strconv"

	"github.com/MaSTeR2W/SADEEM/models/security"
	"github.com/labstack/echo/v4"
)

func Me(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.SetParamValues(
			append(
				c.ParamValues(),
				strconv.FormatUint(
					uint64(c.Get("authData").(*security.AuthData).UserId),
					10,
				),
			)...,
		)
		c.SetParamNames(append(c.ParamNames(), "userId")...)

		return next(c)
	}
}
