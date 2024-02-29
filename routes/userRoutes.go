package routes

import (
	"github.com/MaSTeR2W/SADEEM/controllers/security"
	"github.com/MaSTeR2W/SADEEM/controllers/user"
	"github.com/labstack/echo/v4"
)

func userRoutes(g *echo.Group) {
	var userGrp = g.Group("/users")

	userGrp.POST("", user.Create, security.Authz("manager"))

	userGrp.GET("", user.GetMany, security.Authz("manager"))

	userGrp.GET(
		"/me",
		user.GetOne,
		user.Me,
		user.DoesExist(false),
	)

	userGrp.GET(
		"/:userId",
		user.GetOne,
		security.Authz("manager"),
		user.DoesExist(false),
	)

	userGrp.PATCH(
		"/:userId",
		user.Update,
		security.Authz("manager"),
		user.DoesExist(true),
	)

	userGrp.DELETE(
		"/:userId",
		user.Delete,
		security.Authz("manager"),
		user.DoesExist(true),
	)

	user_classesRoutes(userGrp)
}
