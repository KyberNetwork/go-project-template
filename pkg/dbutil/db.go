package dbutil

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // go migrate
	"github.com/jmoiron/sqlx"
)

// NewDB creates a DB instance from dsn
func NewDB(dsn string) (*sqlx.DB, error) {
	const driverName = "postgres"
	return sqlx.Connect(driverName, dsn)
}

// FormatDSN ..
func FormatDSN(props map[string]string) string {
	var s strings.Builder
	for k, v := range props {
		s.WriteString(k)
		s.WriteString("=")
		s.WriteString(v)
		s.WriteString(" ")
	}
	return s.String()
}

// RunMigrationUp ...
func RunMigrationUp(db *sql.DB, migrationFolderPath, databaseName string) (*migrate.Migrate, error) {
	driver, err := migratepostgres.WithInstance(db, &migratepostgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationFolderPath),
		databaseName, driver,
	)
	if err != nil {
		return nil, err
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return m, nil
}
