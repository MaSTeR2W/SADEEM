package classification

import (
	"strconv"

	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/errors"
	"github.com/MaSTeR2W/SADEEM/models/classification"
	"github.com/MaSTeR2W/SADEEM/postgres"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func DoesExist(forUpdate bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var classId = c.Param("classificationId")
			var lang = c.QueryParam("lang")

			var intClassId, err = strconv.Atoi(classId)

			if err != nil || intClassId < 1 {
				return &errors.HTTP400Err{
					Field:   "classificationId",
					Message: InvalidClassIdErr(lang),
				}
			}

			var fClass *classification.Classification

			if forUpdate {
				var tx *sqlx.Tx
				tx, err = postgres.DB.Beginx()

				if err != nil {
					return err
				}

				defer tx.Rollback()
				c.Set("tx", tx)

				fClass, err = pgHprs.TxQueryxStructScan[classification.Classification](tx, "SELECT * FROM classifications WHERE class_id="+classId+" FOR UPDATE;")

			} else {

				fClass, err = pgHprs.QueryxAndStructScan[classification.Classification]("SELECT * FROM classifications WHERE class_id=" + classId)

			}

			if err != nil {
				return err
			}

			if fClass.ClassId == 0 {
				return &errors.HTTP404Err{
					Message: ClassNotFoundErr(lang),
				}
			}

			c.Set("fClass", fClass)

			return next(c)
		}
	}
}

func InvalidClassIdErr(lang string) string {
	if lang == "ar" {
		return "يجب أن يكون رقم التصنيف عدداً صحيحاً موجباً"
	}

	return "Classification id should be a positive integer"
}

func ClassNotFoundErr(lang string) string {
	if lang == "ar" {
		return "لم يتم العثور على التصنيف"
	}

	return "Classification not found"
}
