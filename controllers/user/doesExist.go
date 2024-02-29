package user

import (
	"strconv"

	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/errors"
	"github.com/MaSTeR2W/SADEEM/models/user"
	"github.com/MaSTeR2W/SADEEM/postgres"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func DoesExist(forUpdate bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var userId = c.Param("userId")
			var intUserId, err = strconv.Atoi(userId)

			if err != nil || intUserId < 1 {
				return &errors.HTTP400Err{
					Field:   "userId",
					Message: InvalidUserIdErr(c.QueryParam("lang")),
				}
			}

			var fUser *user.User

			if forUpdate {
				var tx *sqlx.Tx
				tx, err = postgres.DB.Beginx()
				if err != nil {
					return err
				}
				defer tx.Rollback()
				c.Set("tx", tx)
				fUser, err = pgHprs.TxQueryxStructScan[user.User](tx, "SELECT user_id, name, email, image, user_type FROM users WHERE user_id="+userId+" FOR UPDATE;")

			} else {
				fUser, err = pgHprs.QueryxAndStructScan[user.User]("SELECT * FROM joined_users_classifications WHERE user_id=" + userId)
			}

			if err != nil {
				return err
			}

			if fUser.UserId == 0 {
				return &errors.HTTP404Err{
					Message: UserNotFoundErr(c.QueryParam("lang")),
				}
			}

			c.Set("fUser", fUser)
			return next(c)
		}
	}
}

func InvalidUserIdErr(lang string) string {
	if lang == "ar" {
		return "يجب أن يكون رقم المستخدم عدداً صحيحاً موجباً"
	}

	return "User id should be a positive integer"
}

func UserNotFoundErr(lang string) string {
	if lang == "ar" {
		return "لم يتم العثور على المستخدم"
	}

	return "User not found"
}
