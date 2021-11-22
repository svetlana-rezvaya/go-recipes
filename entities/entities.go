package entities

import (
	"gorm.io/gorm"
)

// Ingredient ...
type Ingredient struct {
	gorm.Model

	Name string
}
