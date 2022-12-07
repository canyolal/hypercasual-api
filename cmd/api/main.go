package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/canyolal/hypercasual-inventories/internal/data"
	"github.com/canyolal/hypercasual-inventories/internal/jsonlog"
	"github.com/canyolal/hypercasual-inventories/internal/mailer"
	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	db   struct {
		username string
		password string
		name     string
		sslmode  string
		host     string
		port     int
	}
	cors struct {
		trustedOrigins []string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type application struct {
	config         config
	logger         *jsonlog.Logger
	models         data.Models
	gamesAndGenres map[string]string
	mailer         mailer.Mailer
	wg             sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4001, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.username, "db-username", os.Getenv("POSTGRES_USER"), "Postgres Username")
	flag.StringVar(&cfg.db.password, "db-password", os.Getenv("POSTGRES_PASSWORD"), "Postgres Password")
	flag.StringVar(&cfg.db.name, "db-name", os.Getenv("POSTGRES_DB"), "Postgres DB Name")
	flag.StringVar(&cfg.db.sslmode, "db-sslmode", "disable", "Postgres SSL Mode")
	flag.StringVar(&cfg.db.host, "db-host", "postgres", "Postgres Host")
	flag.IntVar(&cfg.db.port, "db-port", 5432, "Postgres Port")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 587, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", os.Getenv("SMTP_username"), "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", os.Getenv("SMTP_password"), "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Hypercasual Tracker <no-reply@cyy.com>", "SMTP sender")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()
	logger.PrintInfo("database connection is established", nil)

	app := &application{
		config:         cfg,
		logger:         logger,
		models:         data.NewModels(db),
		gamesAndGenres: make(map[string]string),
		mailer:         mailer.New(cfg.smtp.port, cfg.smtp.host, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	app.runCronGameUpdater()

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

// openDB opens a sql connection pool
func openDB(cfg config) (*sql.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.db.host, cfg.db.port, cfg.db.username, cfg.db.password, cfg.db.name, cfg.db.sslmode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
