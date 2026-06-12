package database

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/mayron1806/api-template/config"
)

type DB struct {
	*sqlx.DB
	Builder squirrel.StatementBuilderType
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func newDatabase(cfg *config.Config) (*DB, error) {
	dbConfig := cfg.Database
	dsn := dbConfig.URL

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)

	return &DB{
		DB:      db,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}
