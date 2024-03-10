package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/monster0freason/golang-bookstore-management-system/pkg/models"
    "github.com/monster0freason/golang-bookstore-management-system/pkg/utils"
    "github.com/gorilla/mux"
)

// GetBooks handles the retrieval of all books from the database.
func GetBooks(w http.ResponseWriter, r *http.Request) {
    newBooks := models.GetAllBooks() // Retrieve all books from the database
    res, _ := json.Marshal(newBooks)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// GetBookByID retrieves a book from the database by its ID.
func GetBookByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID := vars["bookId"]
    ID, err := strconv.ParseInt(bookID, 10, 64) // Assuming bookId is of type int64
    if err != nil {
        fmt.Println("Error while parsing:", err)
        return
    }

    bookDetails, _ := models.GetBookByID(ID)
    res, _ := json.Marshal(bookDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// CreateBook creates a new book in the database.
func CreateBook(w http.ResponseWriter, r *http.Request) {
    newBook := &models.Book{}
    utils.ParseBody(r, newBook)

    createdBook := newBook.CreateBook()
    res, _ := json.Marshal(createdBook)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(res)
}

// UpdateBook updates an existing book in the database.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    var updateBook models.Book
    utils.ParseBody(r, &updateBook)

    vars := mux.Vars(r)
    bookID := vars["bookId"]
    ID, err := strconv.ParseInt(bookID, 10, 64) // Assuming bookId is of type int64
    if err != nil {
        fmt.Println("Error while parsing:", err)
        return
    }

    bookDetails, db := models.GetBookByID(ID)
    if updateBook.Name != "" {
        bookDetails.Name = updateBook.Name
    }
    if updateBook.Author != "" {
        bookDetails.Author = updateBook.Author
    }
    if updateBook.Publication != "" {
        bookDetails.Publication = updateBook.Publication
    }
    db.Save(&bookDetails)
    res, _ := json.Marshal(bookDetails)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

// DeleteBook deletes a book from the database by its ID.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID := vars["bookId"]
    ID, err := strconv.ParseInt(bookID, 10, 64) // Assuming bookId is of type int64
    if err != nil {
        fmt.Println("Error while parsing:", err)
        return
    }

    deletedBook := models.DeleteBook(ID)
    res, _ := json.Marshal(deletedBook)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}


