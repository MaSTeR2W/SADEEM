package security

import (
	"github.com/MaSTeR2W/SADEEM/postgres"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
)

var GetAuthData *sqlx.Stmt

func init() {
	go prepare()
}

func prepare() {
	<-postgres.Migrated

	GetAuthData = pgHprs.MustPreparex("SELECT user_id, user_type FROM users WHERE user_id=$1")

}
