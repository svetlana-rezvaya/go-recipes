package storages

import (
	"gorm.io/gorm"
)

// IngredientStorage ...
type IngredientStorage struct {
	db *gorm.DB
}

// NewIngredientStorage ...
func NewIngredientStorage(db *gorm.DB) IngredientStorage {
	return IngredientStorage{db: db}
}
