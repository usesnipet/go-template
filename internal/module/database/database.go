package database

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/mayron1806/api-template/internal/module/config"
)

type DB struct {
	*sqlx.DB
	Builder squirrel.StatementBuilderType
}

func newDatabase(cfg *config.Config) (*DB, error) {
	dbConfig := cfg.Database
	dsn := dbConfig.URL
	if dsn == "" {
		dsn = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.SSLMode,
		)
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return &DB{
		DB:      db,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}
