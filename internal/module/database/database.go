package database

import (
	"context"
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

// Get loads a single row into dest. Use for SELECT or INSERT/UPDATE ... RETURNING.
// Pass a squirrel builder query or raw SQL via squirrel.Expr("SELECT ...", args...).
func (db *DB) Get(ctx context.Context, dest any, q squirrel.Sqlizer) error {
	query, args, err := toSQL(q)
	if err != nil {
		return err
	}

	if err := db.GetContext(ctx, dest, query, args...); err != nil {
		return fmt.Errorf("get: %w", err)
	}

	return nil
}

// Select loads all matching rows into dest. dest must be a pointer to a slice.
func (db *DB) Select(ctx context.Context, dest any, q squirrel.Sqlizer) error {
	query, args, err := toSQL(q)
	if err != nil {
		return err
	}

	if err := db.SelectContext(ctx, dest, query, args...); err != nil {
		return fmt.Errorf("select: %w", err)
	}

	return nil
}

// Exec runs a statement that does not return rows (INSERT/UPDATE/DELETE without RETURNING).
func (db *DB) Exec(ctx context.Context, q squirrel.Sqlizer) error {
	query, args, err := toSQL(q)
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func toSQL(q squirrel.Sqlizer) (string, []any, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("build query: %w", err)
	}

	return query, args, nil
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
