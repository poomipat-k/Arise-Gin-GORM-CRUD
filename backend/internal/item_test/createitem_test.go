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

type ResponseBody struct {
	Id     int
	Status string
}

func TestCreateItemHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		testName       string
		name           string
		store          *MockItemStore
		expectedStatus int
		expectedId     int
	}{
		{
			testName:       "should error when name is empty",
			name:           "",
			expectedStatus: http.StatusBadRequest},
		{
			testName:       "should error when name length is greater than 100",
			name:           "abcdefghijklmnopqrstuvwxyabcdefghijklmnopqrstuvwxyabcdefghijklmnopqrstuvwxyabcdefghijklmnopqrstuvwxya", // 101 chars
			expectedStatus: http.StatusBadRequest,
		},
		{
			testName:       "should error when store failed to create item",
			name:           "Testing",
			expectedStatus: http.StatusInternalServerError,
			store: &MockItemStore{
				CreateItemFunc: func(input schemas.CreateItemSchemaInput) (*models.Item, error) {
					return nil, errors.New("something wrong")
				},
			},
		},
		{
			testName:       "should create item successfully",
			name:           "Testing",
			expectedStatus: http.StatusOK,
			store: &MockItemStore{
				CreateItemFunc: func(input schemas.CreateItemSchemaInput) (*models.Item, error) {
					return &models.Item{
						Name: "Testing",
						ID:   1,
					}, nil
				},
			},
			expectedId: 1,
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
			c.Request = httptest.NewRequest(http.MethodPost, "/items", bytes.NewBuffer(inputStr))
			c.Request.Header.Set("Content-Type", "application/json")

			itemHandler := item.NewItemHandler(tt.store)

			r.POST("/items", itemHandler.CreateItem)
			r.ServeHTTP(res, c.Request)

			if status := res.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
			if tt.expectedId > 0 {
				var body ResponseBody
				err := json.Unmarshal(res.Body.Bytes(), &body)
				if err != nil {
					t.Errorf("Error unmarshal ResponseBody err:%v", err)
				}
				if body.Id != tt.expectedId {
					t.Errorf("handler returned wrong Id: got %v want %v", body.Id, tt.expectedId)
				}
			}
		})
	}
}
