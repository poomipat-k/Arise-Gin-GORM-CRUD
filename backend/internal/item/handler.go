package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poomipat-k/crud-arise/internal/models"
	"github.com/poomipat-k/crud-arise/internal/schemas"
)

type ItemStore interface {
	CreateItem(input schemas.CreateItemSchemaInput) (*models.Item, error)
}

type ItemHandler struct {
	store ItemStore
}

func NewItemHandler(is ItemStore) *ItemHandler {
	return &ItemHandler{
		store: is,
	}
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var input schemas.CreateItemSchemaInput

	// Validation
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "errorDetails": err.Error()})
		return
	}

	item, err := h.store.CreateItem(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Item created", "id": item.ID})
}
