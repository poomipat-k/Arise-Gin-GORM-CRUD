package itemtest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/poomipat-k/crud-arise/internal/item"
)

func TestDeleteItemById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		testName       string
		itemId         any
		store          *MockItemStore
		expectedStatus int
	}{
		{
			testName:       "should delete item successfully",
			itemId:         5,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {

			res := httptest.NewRecorder()
			c, r := gin.CreateTestContext(res)
			path := fmt.Sprintf("/items/%v", tt.itemId)

			c.Request = httptest.NewRequest(http.MethodDelete, path, nil)
			c.Request.Header.Set("Content-Type", "application/json")

			itemHandler := item.NewItemHandler(tt.store)

			r.DELETE("/items/:id", itemHandler.DeleteItemById)
			r.ServeHTTP(res, c.Request)
			t.Log(res)

			if status := res.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}
