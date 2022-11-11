package testutil

import (
	"fmt"

	"github.com/KyberNetwork/go-project-template/internal/dbutil"
	_ "github.com/golang-migrate/migrate/v4/source/file" // nolint go migrate
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // nolint sql driver name: "postgres"
)

// MustNewDevelopmentDB creates a new development DB.
// It also returns a function to teardown it after the test.
func MustNewDevelopmentDB(migrationPath string) (*sqlx.DB, func() error) {
	const dbNameLen = 8

	dbName := RandomString(dbNameLen)
	defaultHost := "127.0.0.1"
	defaultPort := 5432
	defaultUser := "test"
	defaultPassword := "test"

	dsn := dbutil.FormatDSN(map[string]interface{}{
		"host":     defaultHost,
		"port":     defaultPort,
		"user":     defaultUser,
		"password": defaultPassword,
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

	dsnWithDB := dbutil.FormatDSN(map[string]interface{}{
		"host":     defaultHost,
		"port":     defaultPort,
		"user":     defaultUser,
		"password": defaultPassword,
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
