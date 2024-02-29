package user

import (
	"fmt"
	"os"

	"github.com/MaSTeR2W/SADEEM/helpers/hprFns"
	"github.com/MaSTeR2W/SADEEM/helpers/image"
	"github.com/MaSTeR2W/SADEEM/models/user"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	var language = c.QueryParam("lang")

	var body = map[string]any{}
	c.Bind(&body)

	var err error

	file, fileErr := c.FormFile("image")

	if file != nil {
		body["image"] = ""
	}

	if err = user.UpdateValidator.Validate(body, language); err != nil {
		return err
	}

	var fUser = c.Get("fUser").(*user.User)

	if file != nil {
		if fileErr != nil {
			return fileErr
		}

		var filename string

		filename, err = image.SaveImage(file, "image", 0, 0, language)

		body["image"] = "/imgs/" + filename

		// do not redeclare err again inside this block
		defer func() {
			if err != nil {
				os.Remove(image.ImgsFolderPath + filename)
				fmt.Println("not nil")
			} else {
				os.Remove(image.ImgsFolderPath + fUser.Image[6:])
				fmt.Println("nil")
			}
		}()

		if err != nil {
			return err
		}

	}

	if pass, ok := body["password"]; ok {
		body["password"], body["salt"], err = user.GenerateSaltAndHashPassword(pass.(string))

		if err != nil {
			return err
		}
	}

	var tx = c.Get("tx").(*sqlx.Tx)

	var keys, vals = hprFns.GetExistedKeysVals(
		body,
		"name",
		"email",
		"password",
		"salt",
		"image",
	)

	err = pgHprs.TxQueryx(
		tx,
		"UPDATE users SET "+pgHprs.SetBindVars(keys...)+" WHERE user_id="+fUser.StrUserId(),
		vals...,
	)

	if err != nil {
		return err
	}

	updatedUser, err := pgHprs.StmtQueryxAndStructScan[user.User](tx.Stmtx(user.GetOne), fUser.UserId)

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return c.JSON(200, updatedUser)
}
