package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/paritoshratawal/golang_api_gorilla_mux/book_store/pkg/models"
	"github.com/paritoshratawal/golang_api_gorilla_mux/book_store/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBooks Controller")
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBookById Controller called")
	param := mux.Vars(r)
	id, err := strconv.ParseInt(param["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing id")
	}
	newBook := models.GetBookById(id)
	fmt.Println("newBook", newBook)
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Book Controller fn called****")
	create_book := &models.Book{}
	utils.ParseBody(r, create_book)
	create_book_res := create_book.CreateBook()
	res, _ := json.Marshal(create_book_res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteBookById Controller called")
	param := mux.Vars(r)
	id, err := strconv.ParseInt(param["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing id")
	}
	deletedBook := models.DeleteBookById(id)
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Book Controller called")
	update_book := &models.Book{}
	utils.ParseBody(r, update_book)
	param := mux.Vars(r)
	id, err := strconv.ParseInt(param["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing id")
	}
	bookDetails := models.UpdateBookById(id, update_book)
	// bookDetails := models.GetBookById(id)
	// if update_book.Name != "" {
	// 	bookDetails.Name = update_book.Name
	// }
	// if update_book.Author != "" {
	// 	bookDetails.Author = update_book.Author
	// }
	// if update_book.Publication != "" {
	// 	bookDetails.Publication = update_book.Publication
	// }
	// models.DB.Save(&update_book)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
