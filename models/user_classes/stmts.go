package user_classes

import (
	"github.com/MaSTeR2W/SADEEM/postgres"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
)

var DoesClassHoldedByUser *sqlx.Stmt

var DeleteUserClass *sqlx.Stmt

func init() {
	go prepare()
}

func prepare() {
	<-postgres.Migrated

	DoesClassHoldedByUser = pgHprs.MustPreparex("SELECT user_id, class_id FROM user_classifications WHERE user_id=$1 AND class_id=$2 FOR UPDATE")

	DeleteUserClass = pgHprs.MustPreparex("DELETE FROM user_classifications WHERE user_id=$1 AND class_id=$2")
}
