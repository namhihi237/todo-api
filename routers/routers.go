package routers

import (
	models "todo/models"
	v1 "todo/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.Default() // middleware: gin.Recovery() and gin.Logger()

	db := models.GetDatabase()

	// simple group router
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/items", v1.GetListItems(db))
		apiV1.POST("/items", v1.CreateItem(db))
		apiV1.GET("/items/:id", v1.GetItem(db))
		apiV1.DELETE("/items/:id", v1.DeleteItem(db))
		apiV1.PUT("/items/:id", v1.UpdateItem(db))
	}

	return router
}
