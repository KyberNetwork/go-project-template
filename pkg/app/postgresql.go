package app

import (
	"github.com/KyberNetwork/go-project-template/pkg/dbutil"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // nolint sql driver name: "postgres"
	"github.com/urfave/cli"
)

var (
	PostgresHost = cli.StringFlag{ // nolint: gochecknoglobals
		Name:   "postgres-host",
		Usage:  "PostgresSQL host to connect",
		EnvVar: "POSTGRES_HOST",
		Value:  "127.0.0.1",
	}
	PostgresPort = cli.IntFlag{ // nolint: gochecknoglobals
		Name:   "postgres-port",
		Usage:  "PostgresSQL port to connect",
		EnvVar: "POSTGRES_PORT",
		Value:  5432, // nolint: gomnd
	}
	PostgresUser = cli.StringFlag{ // nolint: gochecknoglobals
		Name:   "postgres-user",
		Usage:  "PostgresSQL user to connect",
		EnvVar: "POSTGRES_USER",
		Value:  "go_project_template",
	}
	PostgresPassword = cli.StringFlag{ // nolint: gochecknoglobals
		Name:   "postgres-password",
		Usage:  "PostgresSQL password to connect",
		EnvVar: "POSTGRES_PASSWORD",
		Value:  "go_project_template",
	}
	PostgresDatabase = cli.StringFlag{ // nolint: gochecknoglobals
		Name:   "postgres-database",
		Usage:  "Postgres database to connect",
		EnvVar: "POSTGRES_DATABASE",
		Value:  "go_project_template",
	}
	PostgresMigrationPath = cli.StringFlag{ // nolint: gochecknoglobals
		Name:   "migration-path",
		Value:  "migrations",
		EnvVar: "MIGRATION_PATH",
	}
)

// PostgresSQLFlags creates new cli flags for PostgreSQL client.
func PostgresSQLFlags(defaultDB string) []cli.Flag {
	db := PostgresDatabase
	db.Value = defaultDB

	return []cli.Flag{
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPassword,
		db,
		PostgresMigrationPath,
	}
}

// NewDBFromContext creates a DB instance from cli flags configuration.
func NewDBFromContext(c *cli.Context) (*sqlx.DB, error) {
	const driverName = "postgres"

	connStr := dbutil.FormatDSN(map[string]interface{}{
		"host":     c.String(PostgresHost.Name),
		"port":     c.Int(PostgresPort.Name),
		"user":     c.String(PostgresUser.Name),
		"password": c.String(PostgresPassword.String()),
		"dbname":   c.String(PostgresDatabase.Name),
		"sslmode":  "disable",
	})

	return sqlx.Connect(driverName, connStr)
}
