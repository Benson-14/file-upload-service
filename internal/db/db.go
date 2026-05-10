package db

import (
	"database/sql"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	*sql.DB
}

func NewDatabase(dbURL string) (*Database, error) {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		slog.Error("error opening database: " + err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		slog.Error("could not connect to database", "url", dbURL, "error", err)
		return nil, err
	}

	slog.Info("Connected to DB successfully")
	return &Database{db}, nil
}
