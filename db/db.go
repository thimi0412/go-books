package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/thimi0412/go-books/model"
)

func Init() {
	env := os.Getenv("ENV")

	if "production" == env {
		env = "production"
	} else {
		env = "development"
	}

	godotenv.Load(".env." + env)
	godotenv.Load()

	db, err := gorm.Open("mysql", os.Getenv("CONNECT"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Book{})
	db.Close()
}

func Connection() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("CONNECT"))
	if err != nil {
		panic(err)
	}
	return db
}
