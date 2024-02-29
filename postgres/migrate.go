package postgres

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Migrated = make(chan bool, 1)

func Migrate(migrationsAbsPath string, version uint) {
	driver, err := pgx.WithInstance(DB.DB, &pgx.Config{})

	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationsAbsPath, "postgres", driver)
	if err != nil {
		panic(err)
	}

	curVer, _, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		panic(err)
	}

	if curVer != version {
		err = m.Migrate(version)
		if err != nil {
			panic(err)
		}
	}
	Migrated <- true
	close(Migrated)
}
