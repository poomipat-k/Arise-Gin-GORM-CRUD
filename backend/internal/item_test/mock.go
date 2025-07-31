package itemtest

import (
	"github.com/poomipat-k/crud-arise/internal/models"
	"github.com/poomipat-k/crud-arise/internal/schemas"
)

type MockItemStore struct {
	CreateItemFunc     func(input schemas.CreateItemSchemaInput) (*models.Item, error)
	GetItemByIdFunc    func(id uint) (models.Item, error)
	DeleteItemByIdFunc func(id uint) error
}

func (m *MockItemStore) CreateItem(input schemas.CreateItemSchemaInput) (*models.Item, error) {
	return m.CreateItemFunc(input)
}

func (m *MockItemStore) GetItemById(id uint) (models.Item, error) {
	return m.GetItemByIdFunc(id)
}

func (m *MockItemStore) DeleteItemById(id uint) error {
	return m.DeleteItemByIdFunc(id)
}
