package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Dsn string `required:"true"`
}

const AppPrefix = "LAF"

func main() {
	var config Config
	err := envconfig.Process(AppPrefix, &config)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := openDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

func openDB(config Config) (*sql.DB, error) {
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
