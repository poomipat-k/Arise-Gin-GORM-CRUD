package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/poomipat-k/crud-arise/internal/controllers"
)

func RegisterRoutes(router *gin.Engine) {

	// Hello world
	router.GET("/", controllers.HelloWorld)
}
