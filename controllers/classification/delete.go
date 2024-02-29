package classification

import (
	"github.com/MaSTeR2W/SADEEM/models/classification"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	var tx = c.Get("tx").(*sqlx.Tx)

	var err = pgHprs.StmtQueryx(
		tx.Stmtx(classification.Delete),
		c.Param("classificationId"),
	)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return c.NoContent(204)

}
