package itemtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/poomipat-k/crud-arise/internal/item"
	"github.com/poomipat-k/crud-arise/internal/schemas"
	"gorm.io/gorm"
)

func TestUpdateByIdHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		testName       string
		itemId         any
		name           string
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
			testName:       "should error when name is empty",
			name:           "",
			itemId:         1,
			expectedStatus: http.StatusBadRequest,
		},
		{
			testName:       "should error when name length is greater than 100",
			name:           "abcdefghijklmnopqrstuvwxyabcdefghijklmnopqrstuvwxyabcdefghijklmnopqrstuvwxyabcdefghijklmnopqrstuvwxya", // 101 chars
			itemId:         1,
			expectedStatus: http.StatusBadRequest,
		},
		{
			testName:       "should error if itemId is not exist",
			itemId:         1000,
			name:           "Swimming",
			expectedStatus: http.StatusNotFound,
			store: &MockItemStore{
				UpdateItemByIdFunc: func(id uint, input schemas.UpdateItemSchemaInput) error {
					return gorm.ErrRecordNotFound
				},
			},
		},
		{
			testName:       "should error if itemId is deleted",
			itemId:         8,
			name:           "Swimming",
			expectedStatus: http.StatusNotFound,
			store: &MockItemStore{
				UpdateItemByIdFunc: func(id uint, input schemas.UpdateItemSchemaInput) error {
					return gorm.ErrRecordNotFound
				},
			},
		},
		{
			testName:       "should update item name successfully",
			itemId:         1,
			name:           "Swimming",
			expectedStatus: http.StatusOK,
			store: &MockItemStore{
				UpdateItemByIdFunc: func(id uint, input schemas.UpdateItemSchemaInput) error {
					return nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			input := schemas.CreateItemSchemaInput{
				Name: tt.name,
			}

			inputStr, err := json.Marshal(input)
			if err != nil {
				log.Fatalf("error Marshal input err: %v", err)
			}

			res := httptest.NewRecorder()
			c, r := gin.CreateTestContext(res)
			path := fmt.Sprintf("/items/%v", tt.itemId)
			c.Request = httptest.NewRequest(http.MethodPut, path, bytes.NewBuffer(inputStr))
			c.Request.Header.Set("Content-Type", "application/json")

			itemHandler := item.NewItemHandler(tt.store)

			r.PUT("/items/:id", itemHandler.UpdateItemById)
			r.ServeHTTP(res, c.Request)

			if status := res.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

		})
	}
}
