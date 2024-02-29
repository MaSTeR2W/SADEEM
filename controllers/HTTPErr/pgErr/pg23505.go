package pgErr

import (
	"strings"

	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
)

func C23505(msg *pgconn.PgError, c echo.Context) {
	var lang = c.QueryParam("lang")
	var err errors.HTTP400Err

	err.Field = strings.Split(msg.ConstraintName, "_")[0]
	switch err.Field {
	case "email":
		if lang == "ar" {
			err.Message = "هذا البريد الإلكتروني قيد الاستخدام بالفعل."
		} else {
			err.Message = "This email address is already in use."
		}
	}

	c.JSON(400, err)
}
