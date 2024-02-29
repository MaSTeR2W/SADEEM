package config

import (
	"github.com/MaSTeR2W/SADEEM/models/user"
	"github.com/MaSTeR2W/SADEEM/postgres"
	"github.com/MaSTeR2W/SADEEM/postgres/pgHprs"
)

func init() {
	go CreateFirstManager()
}

func CreateFirstManager() {
	<-postgres.Migrated

	var count int

	var err = pgHprs.QueryxAndScan("SELECT COUNT(*) FROM users LIMIT 1", []any{}, &count)

	if err != nil {
		panic(err)
	}
	if count != 0 {
		return
	}

	var password, salt []byte
	password, salt, err = user.GenerateSaltAndHashPassword("12345678")

	if err != nil {
		panic(err)
	}

	err = pgHprs.Queryx(
		"INSERT INTO users (name, email,	image,password, salt, user_type) VALUES ($1, $2, $3, $4, $5, $6)",
		"sadeem",
		"sadeem@sadeem.ly",
		"/imgs/20240229_221829_0dd2394ea28f.png",
		password,
		salt,
		"manager",
	)

	if err != nil {
		panic(err)
	}

}
