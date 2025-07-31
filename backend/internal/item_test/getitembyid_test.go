package itemtest

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/poomipat-k/crud-arise/internal/item"
	"github.com/poomipat-k/crud-arise/internal/models"
)

func TestGetItemByIdHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		testName       string
		itemId         any
		store          *MockItemStore
		expectedStatus int
	}{
		{
			testName:       "should error when id is invalid (string)",
			itemId:         "abc",
			expectedStatus: http.StatusBadRequest,
		},
		{
			testName:       "should error when id is invalid (negative)",
			itemId:         -10,
			expectedStatus: http.StatusBadRequest,
		},
		{
			testName:       "should error when id is invalid (float)",
			itemId:         2.35,
			expectedStatus: http.StatusBadRequest,
		},
		{
			testName:       "should error when id is not exist",
			itemId:         1,
			expectedStatus: http.StatusNotFound,
			store: &MockItemStore{
				GetItemByIdFunc: func(id uint) (models.Item, error) {
					return models.Item{}, errors.New("item not exist")
				},
			},
		},
		{
			testName:       "should return exist item",
			itemId:         1,
			expectedStatus: http.StatusOK,
			store: &MockItemStore{
				GetItemByIdFunc: func(id uint) (models.Item, error) {
					return models.Item{
						ID:   1,
						Name: "Testing",
					}, nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {

			res := httptest.NewRecorder()
			c, r := gin.CreateTestContext(res)
			path := fmt.Sprintf("/items/%v", tt.itemId)

			c.Request = httptest.NewRequest(http.MethodGet, path, nil)
			c.Request.Header.Set("Content-Type", "application/json")

			itemHandler := item.NewItemHandler(tt.store)

			r.GET("/items/:id", itemHandler.GetItemById)
			r.ServeHTTP(res, c.Request)

			if status := res.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

		})
	}
}
