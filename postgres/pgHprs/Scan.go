package pgHprs

import (
	"github.com/MaSTeR2W/SADEEM/postgres"

	"github.com/jmoiron/sqlx"
)

func QueryxAndScan(query string, dest ...any) error {
	rows, err := postgres.DB.Queryx(query)

	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(dest...); err != nil {
			return err
		}
	}

	return rows.Err()
}

func StmtQueryxAndScan(stmt *sqlx.Stmt, args []any, dest []any) error {
	rows, err := stmt.Queryx(args...)

	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(dest...); err != nil {
			return err
		}
	}

	return rows.Err()
}
