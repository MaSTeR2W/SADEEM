package routes

import (
	"path"

	"github.com/MaSTeR2W/SADEEM/controllers/security"
	"github.com/MaSTeR2W/SADEEM/helpers/absPath"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var __dirname = absPath.ToMe()

func SetRoutes(c *echo.Echo) {
	var apiGp = c.Group("/api/v1")

	registraionRoutes(apiGp)

	apiGp.Use(security.Authn)

	apiGp.Use(middleware.Static(path.Join(__dirname, "../public")))

	userRoutes(apiGp)
	classificationRoutes(apiGp)
}
