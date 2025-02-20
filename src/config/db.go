package config

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Pool *pgxpool.Pool
	once sync.Once
)

// InitDB initializes the connection pool once
func InitDB() {
	once.Do(func() {
		var err error
		Pool, err = pgxpool.New(context.Background(), os.Getenv("PGX_URL"))

		if err != nil {
			log.Fatalf("Failed to initialize database connection pool: %v", err)
		}
		log.Println("Database connection pool initialized")
	})
}
