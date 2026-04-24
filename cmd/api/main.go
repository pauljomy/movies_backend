package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
	DSN    string
	DB     *sql.DB
}

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	var addr string
	var dsn string

	flag.StringVar(&addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&dsn, "dsn", "host=localhost port=5433 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	app := &application{
		logger: logger,
		DSN:    dsn,
	}

	conn, err := app.connectToDb()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app.DB = conn

	defer app.DB.Close()

	logger.Info("starting server", "addr", addr)

	err = http.ListenAndServe(addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
