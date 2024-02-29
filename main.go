package main

import (
	_ "github.com/MaSTeR2W/SADEEM/config"

	"github.com/MaSTeR2W/SADEEM/controllers/HTTPErr"
	"github.com/MaSTeR2W/SADEEM/helpers/absPath"
	"github.com/MaSTeR2W/SADEEM/postgres"
	"github.com/MaSTeR2W/SADEEM/routes"
	"github.com/labstack/echo/v4"
)

// calculate the path at runtime
var __dirname = absPath.ToMe()

func main() {
	postgres.Migrate(__dirname+"/migrations", 9)

	var e = echo.New()
	e.HTTPErrorHandler = HTTPErr.Handler

	routes.SetRoutes(e)

	e.Logger.Fatal(e.Start("localhost:1323"))

}
