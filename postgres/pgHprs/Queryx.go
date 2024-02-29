package pgHprs

import (
	"github.com/MaSTeR2W/SADEEM/postgres"

	"github.com/jmoiron/sqlx"
)

func Queryx(query string, args ...any) error {
	rows, err := postgres.DB.Queryx(query, args...)

	if err != nil {
		return err
	}

	for rows.Next() {
	}

	return rows.Err()
}

func StmtQueryx(stmt *sqlx.Stmt, args ...any) error {
	rows, err := stmt.Queryx(args...)

	if err != nil {
		return err
	}

	for rows.Next() {
	}

	return rows.Err()
}

func TxQueryx(tx *sqlx.Tx, sql string, args ...any) error {
	rows, err := tx.Queryx(sql, args...)

	if err != nil {
		return err
	}

	for rows.Next() {
	}

	return rows.Err()
}
