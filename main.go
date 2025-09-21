package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupDatabase() (*gorm.DB, error) {
	// Create the database connection string (DSN) using environment variables
	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatalf("Error loading .env file: %v", env_err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	config := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}
        
    // Initialize the database connection with GORM
	db, err := gorm.Open(postgres.Open(dsn), config)

    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    // Get the underlying *sql.DB instance to configure the connection pool
	sqlDB, err := db.DB()

    if err != nil {
        return nil, fmt.Errorf("failed to get database instance: %w", err)
    }

    // Configure the connection pool for optimal performance:
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

    return db, nil
}

func main() {
	_, err := setupDatabase()// Attempt to set up the database connection

    if err != nil {
        log.Fatal(err)  // If setup fails, log the error and exit
    }

    log.Println("Successfully connected to database")
}