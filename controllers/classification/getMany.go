package classification

import (
	"github.com/MaSTeR2W/SADEEM/models/classification"
	"github.com/MaSTeR2W/SADEEM/models/security"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/labstack/echo/v4"
)

func GetMany(c echo.Context) error {
	var sql = "SELECT * FROM classifications"
	var sqlCount = "SELECT COUNT(*) AS count FROM classifications"

	if c.Get("authData").(*security.AuthData).UserType != "manager" {
		sql += " WHERE enabled=TRUE"
		sqlCount += " WHERE enabled=TRUE"
	}

	var classes, err = pgHprs.QueryxAndStructMultiScan[classification.Classification](sql)

	if err != nil {
		return err
	}

	var count int

	err = pgHprs.QueryxAndScan(sqlCount, []any{}, &count)

	if err != nil {
		return err
	}

	return c.JSON(200, map[string]any{
		"count": count,
		"data":  classes,
	})
}
