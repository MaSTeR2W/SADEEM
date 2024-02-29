package user

import (
	"fmt"
	"os"

	"github.com/MaSTeR2W/SADEEM/helpers/image"
	"github.com/MaSTeR2W/SADEEM/models/user"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	var language = c.QueryParam("lang")

	var body = map[string]any{}
	c.Bind(&body)

	var err error

	file, fileErr := c.FormFile("image")

	if file != nil {
		body["image"] = ""
	} else {
		delete(body, "image")
	}

	if err = user.CreationValidator.Validate(body, language); err != nil {
		return err
	}

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
				os.Remove(image.ImgsFolderPath + filename[6:])
				fmt.Println("not nil")
			} else {
				fmt.Println("nil")
			}
		}()

		if err != nil {
			return err
		}

	}

	password, salt, err := user.GenerateSaltAndHashPassword(body["password"].(string))

	if err != nil {
		fmt.Println(err)
		return err
	}

	user, err := pgHprs.StmtQueryxAndStructScan[user.User](
		user.Create,
		body["name"],
		body["email"],
		body["image"],
		password,
		salt,
		body["userType"],
	)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(201, user)
}
