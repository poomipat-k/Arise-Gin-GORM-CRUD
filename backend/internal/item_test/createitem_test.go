package itemtest

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/poomipat-k/crud-arise/internal/item"
	"github.com/poomipat-k/crud-arise/internal/models"
	"github.com/poomipat-k/crud-arise/internal/schemas"
)

func TestCreateItemHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	name := "Reading"
	input := schemas.CreateItemSchemaInput{
		Name: name,
	}

	inputStr, err := json.Marshal(input)
	if err != nil {
		log.Fatalf("error Marshal input err: %v", err.Error())
	}

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodPost, "/items", bytes.NewBuffer(inputStr))
	c.Request.Header.Set("Content-Type", "application/json")

	mockItemStore := &MockItemStore{
		CreateItemFunc: func(input schemas.CreateItemSchemaInput) (*models.Item, error) {
			return &models.Item{
				ID: 1,
			}, errors.New("New Error")
		},
	}
	itemHandler := item.NewItemHandler(mockItemStore)

	r.POST("/items", itemHandler.CreateItem)
	r.ServeHTTP(res, c.Request)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
