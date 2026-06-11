package database

import (
	"context"
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

func (db *DB) Run(
	toSql func() (string, []any, error),
	dest any,
	ctx context.Context,
) error {
	query, args, err := toSql()
	if err != nil {
		return err
	}
	return db.DB.GetContext(ctx, dest, query, args...)
}

func (db *DB) Close() error {
	return db.DB.Close()
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

	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)

	return &DB{
		DB:      db,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}
