package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"vaidik/medium/internal/data"
)

const version = "1.0.0"

type config struct {
	env  string
	port int
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConn  int
		mexIdleTime  string
	}
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "addr", 4000, "http network address")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "dsn for database connection")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "max open connections for database")
	flag.IntVar(&cfg.db.maxIdleConn, "db-max-idle-conn", 25, "max idle connections for database")
	flag.StringVar(&cfg.db.mexIdleTime, "db-max-idle-timeout", "15m", "max timeout for idle connections")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.port),
		Handler:     app.routes(),
		IdleTimeout: 10 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on port %d", cfg.env, cfg.port)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
