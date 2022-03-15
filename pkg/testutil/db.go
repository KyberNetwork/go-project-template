package testutil

import (
	"fmt"
	"strconv"

	libapp "github.com/KyberNetwork/go-project-template/pkg/app"
	"github.com/KyberNetwork/go-project-template/pkg/dbutil"
	_ "github.com/golang-migrate/migrate/v4/source/file" //nolint go migrate
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //nolint sql driver name: "postgres"
)

// MustNewDevelopmentDB creates a new development DB.
// It also returns a function to teardown it after the test.
func MustNewDevelopmentDB(migrationPath string) (*sqlx.DB, func() error) {
	const dbNameLen = 8

	dbName := RandomString(dbNameLen)
	dsn := dbutil.FormatDSN(map[string]string{
		"host":     libapp.DefaultPostgresHost,
		"port":     strconv.Itoa(libapp.DefaultPostgresPort),
		"user":     libapp.DefaultPostgresUser,
		"password": libapp.DefaultPostgresPassword,
		"sslmode":  "disable",
	})

	ddlDB, err := dbutil.NewDB(dsn)
	if err != nil {
		panic(err)
	}

	ddlDB.MustExec(fmt.Sprintf(`CREATE DATABASE "%s"`, dbName))

	if err := ddlDB.Close(); err != nil {
		panic(err)
	}

	dsnWithDB := dbutil.FormatDSN(map[string]string{
		"host":     libapp.DefaultPostgresHost,
		"port":     strconv.Itoa(libapp.DefaultPostgresPort),
		"user":     libapp.DefaultPostgresUser,
		"password": libapp.DefaultPostgresPassword,
		"sslmode":  "disable",
		"dbname":   dbName,
	})

	db, err := dbutil.NewDB(dsnWithDB)
	if err != nil {
		panic(err)
	}

	m, err := dbutil.RunMigrationUp(db.DB, migrationPath, dbName)
	if err != nil {
		panic(err)
	}

	return db, func() error {
		if _, err := m.Close(); err != nil {
			return err
		}

		ddlDB, err := dbutil.NewDB(dsn)
		if err != nil {
			return err
		}

		if _, err = ddlDB.Exec(fmt.Sprintf(`DROP DATABASE "%s"`, dbName)); err != nil {
			return err
		}

		return ddlDB.Close()
	}
}
