package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

// NewDatabaseConnection creates a new database connection.
// It will return an error if the connection fails.
func NewDatabaseConnection(ctx context.Context, cfg *Config) (*sql.DB, error) {
	// define disable as default ssl mode
	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}

	connStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v&binary_parameters=yes", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error().Err(err).Msg("pkg.postgres.NewDatabaseConnection: failed to open database")
		return nil, err
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)

	if err := db.PingContext(ctx); err != nil {
		log.Error().Err(err).Msg("pkg.postgres.NewDatabaseConnection: failed to ping database")
		return nil, err
	}

	return db, nil
}
