package routes

import (
	"github.com/MaSTeR2W/SADEEM/controllers/security"
	"github.com/MaSTeR2W/SADEEM/controllers/user_classes"
	"github.com/labstack/echo/v4"
)

func user_classesRoutes(g *echo.Group) {
	var ucGps = g.Group(
		"/:userId/classifications/:classificationId",
		security.Authz("manager"),
	)

	ucGps.POST("", user_classes.Set, user_classes.DoesExist)

	ucGps.DELETE("", user_classes.Delete, user_classes.DoesExist)
}
