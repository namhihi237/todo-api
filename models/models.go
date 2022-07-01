package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error

	e := godotenv.Load()
	if e != nil {
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

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerConfig,
	})

	if err != nil {
		log.Fatalln("Can't connect to database:", err)
	}

	log.Println("Connected to database", db)

	// auto migration database
	db.AutoMigrate(&TodoItem{})
	db.AutoMigrate(&User{})
}

func GetDatabase() *gorm.DB {
	return db
}
