package user_classes

import (
	"strconv"

	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/errors"
	"github.com/MaSTeR2W/SADEEM/controllers/classification"
	"github.com/MaSTeR2W/SADEEM/controllers/user"
	mClassification "github.com/MaSTeR2W/SADEEM/models/classification"
	mUser "github.com/MaSTeR2W/SADEEM/models/user"
	"github.com/MaSTeR2W/SADEEM/postgres"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func DoesExist(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var lang = c.QueryParam("lang")
		var userId = c.Param("userId")

		var intUserId, err = strconv.Atoi(userId)

		if err != nil || intUserId < 1 {
			return &errors.HTTP400Err{
				Field:   "userId",
				Message: user.InvalidUserIdErr(lang),
			}
		}

		var classId = c.Param("classificationId")
		var intClassId int
		intClassId, err = strconv.Atoi(classId)

		if err != nil || intClassId < 1 {
			return &errors.HTTP400Err{
				Field:   "classId",
				Message: classification.InvalidClassIdErr(lang),
			}
		}
		var tx *sqlx.Tx
		tx, err = postgres.DB.Beginx()

		if err != nil {
			return err
		}

		defer tx.Rollback()
		c.Set("tx", tx)

		var fUser *mUser.User
		fUser, err = pgHprs.TxQueryxStructScan[mUser.User](
			tx,
			"SELECT user_id, name, email,image, user_type FROM users WHERE user_id="+userId+" FOR UPDATE",
		)

		if err != nil {
			return err
		}

		if fUser.UserId == 0 {
			return &errors.HTTP404Err{
				Message: user.UserNotFoundErr(lang),
			}
		}

		var fClass *mClassification.Classification

		fClass, err = pgHprs.TxQueryxStructScan[mClassification.Classification](tx, "SELECT * FROM classifications WHERE class_id="+classId+" FOR UPDATE;")

		if err != nil {
			return err
		}

		if fClass.ClassId == 0 {
			return &errors.HTTP404Err{
				Message: classification.ClassNotFoundErr(lang),
			}
		}

		c.Set("fUser", fUser)
		c.Set("fClass", fClass)

		return next(c)
	}
}
