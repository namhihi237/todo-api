package routers

import (
	models "todo/models"
	v1 "todo/routers/api/v1"

	"github.com/gin-gonic/gin"

	docs "todo/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.Default() // middleware: gin.Recovery() and gin.Logger()
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Todo APP API"
	docs.SwaggerInfo.Description = "This is a sample server Todo APP server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
