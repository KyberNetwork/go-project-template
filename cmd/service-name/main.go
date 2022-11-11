package main

import (
	"fmt"
	"log"
	"os"

	libapp "github.com/KyberNetwork/go-project-template/internal/app"
	"github.com/KyberNetwork/go-project-template/internal/dbutil"
	"github.com/KyberNetwork/go-project-template/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load("sample_file.env")
	app := libapp.NewApp()
	app.Name = "alert service go script runner"
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}

func run(c *cli.Context) error {
	logger, _, flush, err := libapp.NewLogger(c)
	if err != nil {
		return fmt.Errorf("new logger: %w", err)
	}

	defer flush()

	zap.ReplaceGlobals(logger)
	l := logger.Sugar()
	l.Infow("app starting ..")

	db, err := initDB(c)
	if err != nil {
		l.Panicw("cannot init DB", "err", err)
	}

	store := storage.New(db)
	_ = store

	return nil
}

func initDB(c *cli.Context) (*sqlx.DB, error) {
	db, err := libapp.NewDB(map[string]interface{}{
		"host":     c.String(libapp.PostgresHost.Name),
		"port":     c.Int(libapp.PostgresPort.Name),
		"user":     c.String(libapp.PostgresUser.Name),
		"password": c.String(libapp.PostgresPassword.String()),
		"dbname":   c.String(libapp.PostgresDatabase.Name),
		"sslmode":  "disable",
	})
	if err != nil {
		return nil, err
	}

	_, err = dbutil.RunMigrationUp(db.DB, c.String(libapp.PostgresMigrationPath.Name),
		c.String(libapp.PostgresDatabase.Name))
	if err != nil {
		return nil, err
	}
	return db, nil
}
