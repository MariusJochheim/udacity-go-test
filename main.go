package main

import (
	"context"
	"log"
	"math"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatal("DATABASE_URL environment variable is not set")

	}
	var conn *pgx.Conn 
	var err error
	maxRetries := 5

	for i:=0; i<maxRetries; i++ {
		conn, err = pgx.Connect(context.Background(), connString)
		if err == nil {
			break
		}
		log.Printf("Connection failed:%v. Retrying...", err)
		time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
	}
	if (err != nil) {
		log.Fatalf("Unable to connect to database after %d attempts: %v\n", maxRetries, err)
	}
	defer conn.Close(context.Background())

	log.Println("Successfully connected to the database")
}