package postgres

import (
	"os"
	"strings"

	"github.com/MaSTeR2W/SADEEM/envVars"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB_URL string

var DB *sqlx.DB = func() *sqlx.DB {
	envVars.Read()
	return sqlx.MustConnect("pgx", constructConnStr(
		ConnOpt{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			DB:       os.Getenv("DB_NAME"),
		},
	))
}()

type ConnOpt struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
	Opts     map[string]string
}

func constructConnStr(opts ConnOpt) string {
	var connStr = "postgres://" + ifEmpty(opts.User, "postgres") +
		":" + opts.Password +
		"@" + ifEmpty(opts.Host, "localhost") +
		":" + ifEmpty(opts.Port, "5432") +
		"/" + opts.DB

	if opts.Opts != nil {
		connStr += "?"
		var queryStr = make([]string, 0, len(opts.Opts))
		for k, v := range opts.Opts {
			queryStr = append(queryStr, k+"="+v)
		}

		connStr += strings.Join(queryStr, ",")
	}
	DB_URL = connStr
	return connStr
}

func ifEmpty(str1, str2 string) string {
	if str1 == "" {
		return str2
	}
	return str1
}
