package routes

import (
	"github.com/MaSTeR2W/SADEEM/controllers/security"
	"github.com/labstack/echo/v4"
)

func registraionRoutes(g *echo.Group) {
	var regGrp = g.Group("/registration")

	regGrp.POST("/login", security.Login)
}
