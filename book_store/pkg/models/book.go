package models

import (
	"fmt"

	"github.com/paritoshratawal/golang_api_gorilla_mux/book_store/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	Name        string `gorm: "not null" json: "name"`
	Author      string `gorm: "not null" json: "author"`
	Publication string `gorm: "not null" json: "publication"`
}

func init() {
	db_config := &config.Config{
		Host:     "localhost",
		Port:     "5050",
		Password: "paritosh",
		User:     "postgres",
		DBName:   "golang_fiber_demo",
		SSLMode:  "disable",
	}
	config.Connect(db_config)
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db_err := db.Create(&b).Error
	if db_err != nil {
		panic(db_err)
	}
	return b
}

func GetAllBooks() []Book {
	fmt.Println("GetAllBooks")
	book_list := []Book{}
	db_err := db.Find(&book_list).Error
	if db_err != nil {
		panic(db_err)
	}
	return book_list
}

func (b *Book) GetBookById(id int64) (*Book, *gorm.DB) {
	book := Book{}
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}
