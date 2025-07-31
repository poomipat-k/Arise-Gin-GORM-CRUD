package item

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/poomipat-k/crud-arise/internal/models"
	"github.com/poomipat-k/crud-arise/internal/schemas"
)

type ItemStore interface {
	CreateItem(input schemas.CreateItemSchemaInput) (*models.Item, error)
	GetItemById(id uint) (models.Item, error)
	DeleteItemById(id uint) error
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

	c.JSON(http.StatusOK, gin.H{"message": "Item created", "id": item.ID})
}

func (h *ItemHandler) GetItemById(c *gin.Context) {
	idStr := c.Param("id")
	itemId, err := strconv.Atoi(idStr)
	if err != nil || itemId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item ID"})
		return
	}

	item, err := h.store.GetItemById(uint(itemId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) DeleteItemById(c *gin.Context) {
	idStr := c.Param("id")
	itemId, err := strconv.Atoi(idStr)
	if err != nil || itemId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item ID"})
		return
	}

	err = h.store.DeleteItemById(uint(itemId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}
