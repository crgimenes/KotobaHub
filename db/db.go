package db

import (
	"KotobaHub/config"
	_ "embed"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type Database struct {
	db *sqlx.DB
}

var (
	DB *Database
)

func Open() error {
	db, err := sqlx.Open("sqlite", config.CFG.DBPath)
	if err != nil {
		return err
	}

	DB = &Database{
		db: db,
	}

	err = runMigration()

	return err
}

func (d *Database) Close() error {
	return d.db.Close()
}

func Close() {
	DB.Close()
}
