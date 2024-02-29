package HTTPErr

import (
	"fmt"

	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/errors"
	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/pgErr"
	"github.com/MaSTeR2W/SADEEM/helpers/image"
	"github.com/MaSTeR2W/SADEEM/validators"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
)

var message500Ar = []byte(`{"message":"خطأ داخلي في الخادم"}`)
var message500En = []byte(`{"message":"Internal server error"}`)

func Handler(err error, c echo.Context) {
	fmt.Println(err)
	switch e := err.(type) {
	case *pgconn.PgError:
		if e.Code == "23505" {
			pgErr.C23505(e, c)
		} else {
			unhandlerErr(c)
		}

	case *image.UnsupportedType, *validators.ValidationErr, *validators.ValidationErrs, *errors.HTTP400Err:
		c.JSON(400, e)
	case *errors.HTTP404Err:
		c.JSON(404, e)
	case *errors.HTTP401Err:
		c.JSON(401, e)
	case *errors.HTTP403Err:
		c.JSON(403, e)
	default:
		if eErr, ok := e.(*echo.HTTPError); ok {
			if eErr.Code == 404 {
				notFound(c)
				return
			}
		}
		unhandlerErr(c)

	}
}

func unhandlerErr(c echo.Context) {
	if c.QueryParam("lang") == "ar" {
		c.JSONBlob(500, message500Ar)
	} else {
		c.JSONBlob(500, message500En)
	}
}

var notFoundAr = []byte(`{"message":"هذا المسار غير متاح حالياً"}`)
var notFoundEn = []byte(`{"message":"This path is currently unavailable"}`)

func notFound(c echo.Context) {
	if c.QueryParam("lang") == "ar" {
		c.JSONBlob(404, notFoundAr)

	} else {
		c.JSONBlob(404, notFoundEn)

	}
}
