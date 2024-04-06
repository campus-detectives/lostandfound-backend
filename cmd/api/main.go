package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/campus-detectives/lostandfound-backend/internal/data"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"github.com/kelseyhightower/envconfig"
)

const apiVersion = "1.0.0"

type config struct {
	Dsn      string `required:"true"`
	HttpAddr string `envconfig:"http_addr"`
}

type application struct {
	logger *log.Logger
	models data.Models
}

const AppPrefix = "LAF"

func main() {
	var cfg config
	err := envconfig.Process(AppPrefix, &cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Printf("database connection established")

	app := application{
		logger: logger,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr: cfg.HttpAddr, Handler: app.routes(),
	}

	log.Printf("starting server on %s", cfg.HttpAddr)

	err = srv.ListenAndServe()
	log.Fatal(err)

}

func openDB(config config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
