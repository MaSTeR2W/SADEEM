package classification

import (
	"github.com/MaSTeR2W/SADEEM/models/classification"
	"github.com/MaSTeR2W/SADEEM/models/security"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/labstack/echo/v4"
)

func GetMany(c echo.Context) error {
	var sql = "SELECT * FROM classifications"

	if c.Get("authData").(*security.AuthData).UserType != "manager" {
		sql += " WHERE enabled=TRUE"
	}
	var classes, err = pgHprs.QueryxAndStructMultiScan[classification.Classification](sql)

	if err != nil {
		return err
	}

	return c.JSON(200, classes)
}
