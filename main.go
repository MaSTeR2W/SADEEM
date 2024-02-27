package main

import (
	"github.com/MaSTeR2W/SADEEM/helpers/absPath"
	"github.com/MaSTeR2W/SADEEM/postgres"
)

// calculate the path at runtime
var __dirname = absPath.ToMe()

func main() {
	if err := postgres.MigrateUp(__dirname+"/migrations", 8); err != nil {
		panic(err)
	}
}
