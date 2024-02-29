package routes

import (
	"github.com/MaSTeR2W/SADEEM/controllers/security"
	"github.com/labstack/echo/v4"
)

func SetRoutes(c *echo.Echo) {
	var apiGp = c.Group("/api/v1")

	registraionRoutes(apiGp)

	apiGp.Use(security.Authn)

	userRoutes(apiGp)
	classificationRoutes(apiGp)
}
