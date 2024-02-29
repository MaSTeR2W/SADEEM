package user

import (
	"github.com/MaSTeR2W/SADEEM/postgres"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
	"github.com/jmoiron/sqlx"
)

var Create *sqlx.Stmt

var GetOne *sqlx.Stmt

var Delete *sqlx.Stmt

func init() {
	go prepareStmts()
}

func prepareStmts() {
	// migrations should take a place first
	<-postgres.Migrated

	Create = pgHprs.MustPreparex("INSERT INTO users (name, email,	image,password, salt, user_type) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id, name, email, image, user_type;")

	GetOne = pgHprs.MustPreparex("SELECT * FROM joined_users_classifications WHERE user_id=$1")

	Delete = pgHprs.MustPreparex("DELETE FROM users WHERE user_id=$1")
}
