package storages

import (
	"fmt"

	"github.com/svetlana-rezvaya/go-recipes-backend/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenDB ...
func OpenDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to open the DB: %w", err)
	}

	if err := db.AutoMigrate(&entities.Ingredient{}); err != nil {
		return nil, fmt.Errorf("unable to auto-migrate: %w", err)
	}

	return db, nil
}
