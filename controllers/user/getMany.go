package user

import (
	"github.com/MaSTeR2W/SADEEM/models/user"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/labstack/echo/v4"
)

func GetMany(c echo.Context) error {
	var users, err = pgHprs.QueryxAndStructMultiScan[user.User]("SELECT * FROM joined_users_classifications")

	if err != nil {
		return err
	}

	return c.JSON(200, users)
}
