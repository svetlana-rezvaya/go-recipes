package storages

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenDB ...
func OpenDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to open the DB: %w", err)
	}

	return db, nil
}
