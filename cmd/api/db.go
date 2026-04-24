package main

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) connectToDb(dsn string) (*sql.DB, error) {
	db, err := openDb(dsn)
	if err != nil {
		return nil, err
	}
	app.logger.Info("Connected to Postgres DB")
	return db, nil
}
