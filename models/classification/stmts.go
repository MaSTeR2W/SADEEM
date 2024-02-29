package classification

import (
	"github.com/MaSTeR2W/SADEEM/postgres"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
)

var Create *sqlx.Stmt

var Delete *sqlx.Stmt

func init() {
	go prepareStmts()
}

func prepareStmts() {
	<-postgres.Migrated

	Create = pgHprs.MustPreparex("INSERT INTO classifications (name, enabled) VALUES ($1, $2) RETURNING class_id, name, enabled")

	Delete = pgHprs.MustPreparex("DELETE FROM classifications WHERE class_id=$1")
}
