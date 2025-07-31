package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poomipat-k/crud-arise/internal/database"
	"github.com/poomipat-k/crud-arise/internal/item"
)

func RegisterRoutes(router *gin.Engine) {

	// Hello world
	itemStore := item.NewItemStore(database.DB())
	itemHandler := item.NewItemHandler(itemStore)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello world2")
	})

	router.POST("/items", itemHandler.CreateItem)
	router.GET("/items/:id", itemHandler.GetItemById)
	router.PUT("/items/:id", itemHandler.UpdateItemById)
	router.DELETE("/items/:id", itemHandler.DeleteItemById)
}
