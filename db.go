package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(connString string) {
	var err error

	for i := 0; i < 10; i++ {
		DB, err = pgxpool.New(context.Background(), connString)
		if err == nil {
			err = DB.Ping(context.Background())
		}
		if err == nil {
			break
		}
		log.Printf("Database not ready yet (%v). Retrying in 2s...\n", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Unable to connect to database : %v\n", err)
	}

	log.Println("Connected to database successfully !")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS logs (
		id SERIAL PRIMARY KEY,
		user_id VARCHAR(255) NOT NULL,
		action TEXT NOT NULL,
		latency_ms INT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = DB.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create logs table: %v\n", err)
	}

	log.Println("Logs table is ready!")
}
