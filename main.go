package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id        int        `json:"id" gorm:"primary_key;auto_increment;not null;unique;index;"`
	Title     string     `json:"title" gorm:"size:255; not null;"`
	Status    string     `json:"status" gorm:"size:30; not null; default:'doing';"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"` // why using *time.Time? and how to default value for time
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Can't connect to database:", err)
	}

	log.Println("Connected to database", db)
	// auto migration database
	// note: not delete column
	db.AutoMigrate(&TodoItem{})

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/items", getListItems(db))
	}

	router.Run(":8000")
}

func getListItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"data": []TodoItem{}})
	}
}
