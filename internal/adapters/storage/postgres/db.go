package postgres

import (
	"context"
	"embed"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var migrationsFS embed.FS

type DB struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

// New creates a new instance of DB struct. It initializes a new PostgreSQL connection pool,//+
// pings the database to ensure the connection is working, and returns a pointer to the DB struct.//+
// //+
// ctx: The context for the database connection.//+
// config: The configuration for the database connection. The type of this parameter is "any" to allow for flexibility.//+
// //+
// Returns://+
// - A pointer to the DB struct if the connection is successful and the ping passes.//+
// - An error if the connection fails or the ping fails.//
func New(ctx context.Context, config any) (*DB, error) {
	url := fmt.Sprint("")

	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DB{
		db,
		&psql,
		url,
	}, nil
}

// Migrate applies database migrations using the embedded migrationsFS.
// It initializes a new migration source using the iofs package and the "migrations" directory.
// Then, it creates a new migration instance using the initialized source and the database URL.
// Finally, it applies the migrations using the Up method.
// If any error occurs during the migration process, it returns the error.
// If no changes are made during the migration process, it returns nil.
func (db *DB) Migrate() error {
	driver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", driver, db.url)
	if err != nil {
		return err
	}

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

// ErrorCode extracts and returns the PostgreSQL error code from the given error.
//
// This function is intended to be used in conjunction with the DB struct's methods that interact with the PostgreSQL database.
// It assumes that the provided error is of type *pgconn.PgError, which is the type of error returned by the pgx library when a PostgreSQL error occurs.
//
// Parameters:
// - err: The error returned by a DB struct method. This error should be of type *pgconn.PgError.
//
// Returns:
// - A string representing the PostgreSQL error code. If the provided error is not of type *pgconn.PgError, this function will panic.
func (db *DB) ErrorCode(err error) string {
	pgErr := err.(*pgconn.PgError)
	return pgErr.Code
}

// Close closes the database connection pool.
//
// This function should be called when the DB instance is no longer needed to free up resources.
// After calling this method, the DB instance should not be used for any further database operations.
//
// It is safe to call this method multiple times.
//
// No return value is expected.
func (db *DB) Close() {
	db.Pool.Close()
}
