package routes

import (
	"github.com/MaSTeR2W/SADEEM/controllers/classification"
	"github.com/MaSTeR2W/SADEEM/controllers/security"
	"github.com/labstack/echo/v4"
)

func classificationRoutes(g *echo.Group) {
	var classesGps = g.Group("/classifications")

	classesGps.POST(
		"",
		classification.Create,
		security.Authz("manager"),
	)

	classesGps.GET("", classification.GetMany)

	classesGps.PATCH(
		"/:classificationId",
		classification.Update,
		security.Authz("manager"),
		classification.DoesExist(true),
	)

	classesGps.GET(
		"/:classificationId",
		classification.GetOne,
		security.Authz("manager"),
		classification.DoesExist(false),
	)

	classesGps.DELETE(
		"/:classificationId",
		classification.Delete,
		security.Authz("manager"),
		classification.DoesExist(true),
	)
}
