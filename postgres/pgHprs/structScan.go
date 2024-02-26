package pgHprs

import (
	"github.com/MaSTeR2W/SADEEM/postgres"

	"github.com/jmoiron/sqlx"
)

func QueryxAndStructScan[T any](sql string, args ...any) (*T, error) {
	rows, err := postgres.DB.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	var data T

	for rows.Next() {
		if err := rows.StructScan(&data); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &data, nil
}

func QueryxAndStructMultiScan[T any](sql string, args ...any) ([]*T, error) {
	rows, err := postgres.DB.Queryx(sql, args...)

	if err != nil {
		return nil, err
	}

	var records = make([]*T, 0)

	for rows.Next() {
		var data T
		if err := rows.StructScan(&data); err != nil {
			return nil, err
		}
		records = append(records, &data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func StmtQueryxAndStructScan[T any](stmt *sqlx.Stmt, args ...any) (*T, error) {
	rows, err := stmt.Queryx(args...)

	if err != nil {
		return nil, err
	}

	var data T

	for rows.Next() {
		if err := rows.StructScan(&data); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &data, nil
}

func StmtQueryxAndStructMultiScan[T any](stmt *sqlx.Stmt, args ...any) ([]*T, error) {
	rows, err := stmt.Queryx(args...)

	if err != nil {
		return nil, err
	}

	var records = make([]*T, 0)

	for rows.Next() {
		var data T
		if err := rows.StructScan(&data); err != nil {
			return nil, err
		}
		records = append(records, &data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func TxQueryxStructScan[T any](tx *sqlx.Tx, sql string, args ...any) (*T, error) {
	rows, err := tx.Queryx(sql, args...)

	if err != nil {
		return nil, err
	}

	var data T

	for rows.Next() {
		if err := rows.StructScan(&data); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &data, nil
}
