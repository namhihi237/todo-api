package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TodoItem struct {
	Id        int        `json:"id" gorm:"primary_key;auto_increment;not null;unique;index;"`
	Title     string     `json:"title" gorm:"size:255; not null;"`
	Status    string     `json:"status" gorm:"size:30; not null; default:'doing';"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"` // why using *time.Time?
	UpdatedAt *time.Time `json:"updated_at"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"default:null;"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.ExpandEnv("${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Local")
	// parseTime=True => auto scan DATE DATETIME to time.Time
	// loc=Local => use local timezone

	go_env := os.Getenv("GO_ENV")
	var loggerConfig logger.Interface
	if go_env != "production" {
		loggerConfig = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				Colorful: false,
				LogLevel: logger.Info,
			},
		)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerConfig,
	})

	if err != nil {
		log.Fatalln("Can't connect to database:", err)
	}

	log.Println("Connected to database", db)
	// auto migration database
	// note: not delete column
	db.AutoMigrate(&TodoItem{})

	router := gin.Default() // middleware: gin.Recovery() and gin.Logger()

	// simple group router
	v1 := router.Group("/api/v1")
	{
		v1.GET("/items", getListItems(db))
		v1.POST("/items", createItem(db))
		v1.GET("/items/:id", getItem(db))
		v1.DELETE("/items/:id", deleteItem(db))
		v1.PUT("/items/:id", updateItem(db))
	}

	router.Run(":8000")
}

func getListItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var items []TodoItem
		if err := db.Find(&items).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": items})
	}
}

func createItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item TodoItem

		//To bind a request body into a type, use model binding.
		//When using the Bind-method, Gin tries to infer the binder depending on the Content-Type header.
		if err := c.ShouldBind(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		item.Title = strings.TrimSpace(item.Title)
		if item.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
			return
		}

		if err := db.Create(&item).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": item.Id})
	}
}

func getItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id")) // convert string to int
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var item TodoItem
		if err := db.First(&item, id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": item})
	}
}

func deleteItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var item TodoItem
		if err := db.First(&item, id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Delete(&item).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": item})
	}
}

func updateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var item TodoItem

		if err := c.ShouldBind(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.First(&item, id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		item.Title = strings.TrimSpace(item.Title)
		if item.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
			return
		}

		if err := db.Save(&item).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": item})
	}
}
