package item

import (
	"fmt"

	"github.com/poomipat-k/crud-arise/internal/models"
	"github.com/poomipat-k/crud-arise/internal/schemas"
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

func NewItemStore(db *gorm.DB) *store {
	return &store{
		db: db,
	}
}

func (s *store) CreateItem(input schemas.CreateItemSchemaInput) (*models.Item, error) {
	newItem := models.Item{
		Name: input.Name,
	}

	if err := s.db.Create(&newItem).Error; err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	return &newItem, nil
}
