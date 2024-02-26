package pgHprs

import (
	"github.com/MaSTeR2W/SADEEM/postgres"

	"github.com/jmoiron/sqlx"
)

func MustPreparex(sql string) *sqlx.Stmt {

	stmt, err := postgres.DB.Preparex(sql)
	if err != nil {
		panic(err)
	}
	return stmt
}
