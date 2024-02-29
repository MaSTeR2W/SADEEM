package user

import (
	"os"

	"github.com/MaSTeR2W/SADEEM/helpers/image"
	"github.com/MaSTeR2W/SADEEM/models/user"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	var fUser = c.Get("fUser").(*user.User)
	var tx = c.Get("tx").(*sqlx.Tx)

	var err = pgHprs.StmtQueryx(
		tx.Stmtx(user.Delete),
		c.Param("userId"),
	)

	if err != nil {
		return err
	}

	os.Remove(image.ImgsFolderPath + fUser.Image[6:])
	err = tx.Commit()

	if err != nil {
		return err
	}

	return c.NoContent(204)
}
