package app

import (
	"github.com/KyberNetwork/go-project-template/pkg/dbutil"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //nolint sql driver name: "postgres"
	"github.com/urfave/cli"
)

const (
	PostgresHostFlag    = "postgres-host"
	DefaultPostgresHost = "127.0.0.1"

	PostgresPortFlag    = "postgres-port"
	DefaultPostgresPort = 5432

	PostgresUserFlag    = "postgres-user"
	DefaultPostgresUser = "go_project_template"

	PostgresPasswordFlag    = "postgres-password"
	DefaultPostgresPassword = "go_project_template"

	PostgresDatabaseFlag = "postgres-database"

	PostgresMigrationPath = "migration-path"
)

// PostgresSQLFlags creates new cli flags for PostgreSQL client.
func PostgresSQLFlags(defaultDB string) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   PostgresHostFlag,
			Usage:  "PostgresSQL host to connect",
			EnvVar: "POSTGRES_HOST",
			Value:  DefaultPostgresHost,
		},
		cli.IntFlag{
			Name:   PostgresPortFlag,
			Usage:  "PostgresSQL port to connect",
			EnvVar: "POSTGRES_PORT",
			Value:  DefaultPostgresPort,
		},
		cli.StringFlag{
			Name:   PostgresUserFlag,
			Usage:  "PostgresSQL user to connect",
			EnvVar: "POSTGRES_USER",
			Value:  DefaultPostgresUser,
		},
		cli.StringFlag{
			Name:   PostgresPasswordFlag,
			Usage:  "PostgresSQL password to connect",
			EnvVar: "POSTGRES_PASSWORD",
			Value:  DefaultPostgresPassword,
		},
		cli.StringFlag{
			Name:   PostgresDatabaseFlag,
			Usage:  "Postgres database to connect",
			EnvVar: "POSTGRES_DATABASE",
			Value:  defaultDB,
		},
		cli.StringFlag{
			Name:   PostgresMigrationPath,
			Value:  "migrations",
			EnvVar: "MIGRATION_PATH",
		},
	}
}

// NewDBFromContext creates a DB instance from cli flags configuration.
func NewDBFromContext(c *cli.Context) (*sqlx.DB, error) {
	const driverName = "postgres"

	connStr := dbutil.FormatDSN(map[string]string{
		"host":     c.String(PostgresHostFlag),
		"port":     c.String(PostgresPortFlag),
		"user":     c.String(PostgresUserFlag),
		"password": c.String(PostgresPasswordFlag),
		"dbname":   c.String(PostgresDatabaseFlag),
		"sslmode":  "disable",
	})

	return sqlx.Connect(driverName, connStr)
}
