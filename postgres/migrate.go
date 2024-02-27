package postgres

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateUp(migrationsAbsPath string, version uint) error {
	driver, err := pgx.WithInstance(DB.DB, &pgx.Config{})

	if err != nil {
		return nil
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationsAbsPath, "postgres", driver)
	if err != nil {
		panic(err)
	}

	curVer, _, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return err
	}

	if curVer < version {
		err = m.Migrate(version)
		if err != nil {
			return err
		}
	}

	return nil
}
