package itemtest

import (
	"github.com/poomipat-k/crud-arise/internal/models"
	"github.com/poomipat-k/crud-arise/internal/schemas"
)

type MockItemStore struct {
	CreateItemFunc func(input schemas.CreateItemSchemaInput) (*models.Item, error)
}

func (m *MockItemStore) CreateItem(input schemas.CreateItemSchemaInput) (*models.Item, error) {
	return m.CreateItemFunc(input)
}
