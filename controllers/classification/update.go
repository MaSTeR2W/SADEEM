package classification

import (
	"strings"

	"github.com/MaSTeR2W/SADEEM/helpers/hprFns"
	"github.com/MaSTeR2W/SADEEM/models/classification"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	var body = map[string]any{}
	var lang = c.QueryParam("lang")

	c.Bind(&body)

	var err error

	if err = classification.UpdateValidator.Validate(body, lang); err != nil {
		return err
	}

	var tx = c.Get("tx").(*sqlx.Tx)

	var keys, vals = hprFns.GetExistedKeysVals(body, "name", "enabled")

	body, err = pgHprs.TxQueryxAndMapScan(
		tx,
		"UPDATE classifications SET "+pgHprs.SetBindVars(keys...)+" WHERE class_id="+c.Param("classificationId")+" RETURNING "+strings.Join(keys, ","),
		vals...,
	)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return c.JSON(200, body)
}
