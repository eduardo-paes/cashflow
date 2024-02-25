package data

import (
	"fmt"
	"os"

	// Database schema migration library
	core "github.com/eduardo-paes/cashflow/core/entities"
	"github.com/spf13/viper"  // Configuration management library
	"gorm.io/driver/postgres" // Postgres Driver
	"gorm.io/gorm"            // Database ORM
)

// GetConnection returns a Gorm DB instance for PostgreSQL
func GetConnection() (*gorm.DB, error) {
	databaseURL := viper.GetString("database.url")

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return db, nil
}

// RunMigrations run scripts on path database/migrations
func RunMigrations(db *gorm.DB) error {
	// Automatically create tables for all registered models
	if err := db.AutoMigrate(&core.Expense{}); err != nil {
		return err
	}

	// Automatically create tables for all registered models
	if err := db.AutoMigrate(&core.User{}); err != nil {
		return err
	}

	return nil
}
