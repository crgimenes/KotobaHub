package db

import (
	"database/sql"
)

var migrations = []string{}

func runMigration() error {
	// sqlite
	const sqlCreateMigration = `CREATE TABLE IF NOT EXISTS migration (
		value INTEGER NOT NULL PRIMARY KEY);`

	_, err := DB.db.Exec(sqlCreateMigration)
	if err != nil {
		return err
	}

	const sqlReadMigration = `SELECT value FROM migration;`

	// read last migration from database
	var lastMigration int
	err = DB.db.QueryRow(sqlReadMigration).Scan(&lastMigration)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	// run migrations
	for i := lastMigration + 1; i < len(migrations); i++ {
		_, err = DB.db.Exec(migrations[i])
		if err != nil {
			return err
		}
	}

	return nil
}
