package classification

import "github.com/labstack/echo/v4"

func GetOne(c echo.Context) error {
	return c.JSON(200, c.Get("fClass"))
}
