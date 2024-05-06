package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/i101dev/books-mysql/pkg/models"
	"github.com/i101dev/books-mysql/pkg/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {

	newBook := &models.Book{}

	utils.ParseBody(r, newBook)

	b := newBook.CreateBook()

	res, err := json.Marshal(b)

	if err != nil {
		fmt.Println("\n*** >>> error marshalling JSON")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	newBooks := models.GetAllBooks()

	res, err := json.Marshal(newBooks)

	if err != nil {
		fmt.Println("\n*** >>> error marshalling JSON")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
func GetBookByID(w http.ResponseWriter, r *http.Request) {

	bookID, err := getBookID(r)

	if err != nil {
		fmt.Print("\n*** >>> error parsing [bookId]")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bookDat, _ := models.GetBookByID(bookID)

	MarshalAndSend(bookDat, w)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var updatedBook = &models.Book{}

	utils.ParseBody(r, updatedBook)

	bookID, err := getBookID(r)

	if err != nil {
		fmt.Print("\n*** >>> error parsing [bookId]")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bookDat, db := models.GetBookByID(bookID)

	if updatedBook.Name != "" {
		bookDat.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDat.Author = updatedBook.Author
	}
	if updatedBook.Publisher != "" {
		bookDat.Publisher = updatedBook.Publisher
	}

	db.Save(&bookDat)

	MarshalAndSend(bookDat, w)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	bookID, err := getBookID(r)

	if err != nil {
		fmt.Print("\n*** >>> error parsing [bookId]")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bookDat := models.DeleteBook(bookID)

	MarshalAndSend(&bookDat, w)
}

// ------------------------------------------------------------------

func getBookID(r *http.Request) (int64, error) {

	vars := mux.Vars(r)

	bookId := vars["bookID"]

	ID, err := strconv.ParseInt(bookId, 0, 64)

	return ID, err
}
func MarshalAndSend(dat *models.Book, w http.ResponseWriter) {

	res, err := json.Marshal(dat)

	if err != nil {
		fmt.Println("\n*** >>> error marshalling JSON")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
