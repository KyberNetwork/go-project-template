package main

import (
	"fmt"
	"log"
	"os"

	libapp "github.com/KyberNetwork/go-project-template/pkg/app"
	"github.com/KyberNetwork/go-project-template/pkg/dbutil"
	"github.com/KyberNetwork/go-project-template/pkg/storage"
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

	db, err := libapp.NewDBFromContext(c)
	if err != nil {
		l.Panicw("cannot init DB connection", "err", err)
	}

	_, err = dbutil.RunMigrationUp(db.DB, c.String(libapp.PostgresMigrationPath), c.String(libapp.PostgresDatabaseFlag))
	if err != nil {
		l.Panicw("cannot init DB", "err", err)
	}

	store := storage.New(db)
	_ = store

	return nil
}
