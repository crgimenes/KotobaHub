package db

import (
	"database/sql"
	"log"
)

var migrations = []string{}

func runMigration() error {

	// begin transaction
	tx, err := DB.db.Begin()
	if err != nil {
		return err
	}

	// rollback transaction if error
	defer func() {
		if err != nil {
			e := tx.Rollback()
			if e != nil {
				log.Println(e)
			}
			return
		}
	}()

	// sqlite
	const sqlCreateMigration = `CREATE TABLE IF NOT EXISTS migration (
		value INTEGER NOT NULL PRIMARY KEY);`

	_, err = tx.Exec(sqlCreateMigration)
	if err != nil {
		return err
	}

	const sqlReadMigration = `SELECT value FROM migration;`

	// read last migration from database
	var lastMigration int
	err = tx.QueryRow(sqlReadMigration).Scan(&lastMigration)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	// run migrations
	for i := lastMigration + 1; i < len(migrations); i++ {
		_, err = tx.Exec(migrations[i])
		if err != nil {
			return err
		}
	}

	// update migration
	const sqlUpdateMigration = `INSERT INTO migration (value) VALUES (?);`
	_, err = tx.Exec(sqlUpdateMigration, len(migrations)-1)
	if err != nil {
		return err
	}

	// commit transaction
	err = tx.Commit()

	return err
}
