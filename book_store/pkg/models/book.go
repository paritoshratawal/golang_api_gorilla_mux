package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/paritoshratawal/golang_api_gorilla_mux/book_store/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	Id          uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name        string `gorm: "not null" json: "name"`
	Author      string `gorm: "not null" json: "author"`
	Publication string `gorm: "not null" json: "publication"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db_config := &config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
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

func GetBookById(id int64) *Book {
	var book Book
	db.Where("id=?", id).Find(&book)
	return &book
}

func DeleteBookById(id int64) *Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return &book
}

func UpdateBookById(id int64, update_book *Book) *Book {
	var book_details Book
	db.Where("id=?", id).Find(&book_details)
	if update_book.Name != "" {
		book_details.Name = update_book.Name
	}
	if update_book.Author != "" {
		book_details.Author = update_book.Author
	}
	if update_book.Publication != "" {
		book_details.Publication = update_book.Publication
	}
	db.Save(&book_details)
	return &book_details
}
