package classification

import (
	"github.com/MaSTeR2W/SADEEM/models/classification"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	var body = map[string]any{}
	var language = c.QueryParam("lang")
	c.Bind(&body)

	var err error

	if err = classification.CreationValidator.Validate(body, language); err != nil {
		return err
	}

	var class *classification.Classification

	class, err = pgHprs.StmtQueryxAndStructScan[classification.Classification](
		classification.Create,
		body["name"],
		body["enabled"],
	)

	if err != nil {
		return err
	}

	return c.JSON(201, class)

}
