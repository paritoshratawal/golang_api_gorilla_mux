package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/paritoshratawal/golang_api_gorilla_mux/book_store/pkg/models"
	"github.com/paritoshratawal/golang_api_gorilla_mux/book_store/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBooks Controller")
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func GetBookById(w http.ResponseWriter, r *http.Request) {
// 	param := mux.Vars(r)
// 	id, err := strconv.ParseInt(param["id"], 0, 0)
// 	if err != nil {
// 		fmt.Println("Error while parsing id")
// 	}
// 	newBook, _ := models.GetBookById(id)
// 	res, _ := json.Marshal(newBook)
// 	w.Header().set("Content-Type", "application/json")
// 	w.WriteHeader(Http.StatusOK)
// 	w.Write(res)
// }

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******Create Book Controller fn called****")
	create_book := &models.Book{}
	utils.ParseBody(r, create_book)
	create_book_res := create_book.CreateBook()
	res, _ := json.Marshal(create_book_res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
