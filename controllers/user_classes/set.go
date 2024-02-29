package user_classes

import (
	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr/errors"
	"github.com/MaSTeR2W/SADEEM/models/classification"
	"github.com/MaSTeR2W/SADEEM/models/user"
	"github.com/MaSTeR2W/SADEEM/models/user_classes"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Set(c echo.Context) error {

	var (
		tx     = c.Get("tx").(*sqlx.Tx)
		fUser  = c.Get("fUser").(*user.User)
		fClass = c.Get("fClass").(*classification.Classification)
	)

	user_class, err := pgHprs.StmtQueryxAndStructScan[user_classes.User_Class](
		tx.Stmtx(user_classes.DoesClassHoldedByUser),
		fUser.UserId, fClass.ClassId,
	)

	if err != nil {
		return err
	}

	if user_class.UserId != 0 {
		return &errors.HTTP400Err{
			Message: alreadyAssignedClassErr(c.QueryParam("lang")),
		}
	}

	var newClasses []*classification.Classification

	newClasses, err = pgHprs.TxQueryxAndStructMultiScan[classification.Classification](
		tx,
		"SELECT * FROM add_class_to_user($1, $2)",
		fUser.UserId,
		fClass.ClassId,
	)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	fUser.Classifications = newClasses

	return c.JSON(201, fUser)

}

func alreadyAssignedClassErr(lang string) string {
	if lang == "ar" {
		return "المستخدم لديه هذا التصنيف بالفعل"
	}
	return "The user already have this classification"
}
