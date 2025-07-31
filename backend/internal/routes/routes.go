package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/poomipat-k/crud-arise/internal/database"
	"github.com/poomipat-k/crud-arise/internal/item"
)

func RegisterRoutes(router *gin.Engine) {

	// Hello world
	itemStore := item.NewItemStore(database.DB())
	itemHandler := item.NewItemHandler(itemStore)

	router.POST("/items", itemHandler.CreateItem)
	router.GET("/items/:id", itemHandler.GetItemById)
	router.DELETE("/items/:id", itemHandler.DeleteItemById)
}
