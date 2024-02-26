package pgHprs

import (
	"github.com/MaSTeR2W/SADEEM/postgres"

	"github.com/jmoiron/sqlx"
)

func QueryxAndMapScan(sql string) (map[string]any, error) {

	rows, err := postgres.DB.Queryx(sql)
	if err != nil {
		return nil, err
	}

	var t = map[string]any{}

	for rows.Next() {
		if err := rows.MapScan(t); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return t, nil
}

func QueryxAndMyMapScan(sql string, m map[string]any) error {

	rows, err := postgres.DB.Queryx(sql)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.MapScan(m); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

func StmtQueyxAndMapScan(stmt *sqlx.Stmt, args ...any) (map[string]any, error) {
	rows, err := stmt.Queryx(args...)
	if err != nil {
		return nil, err
	}

	var m = map[string]any{}

	for rows.Next() {
		if err := rows.MapScan(m); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return m, nil
}

func StmtQueyxAndMyMapScan(stmt *sqlx.Stmt, m map[string]any, args ...any) error {
	rows, err := stmt.Queryx(args...)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.MapScan(m); err != nil {
			return err
		}
	}

	return rows.Err()
}
