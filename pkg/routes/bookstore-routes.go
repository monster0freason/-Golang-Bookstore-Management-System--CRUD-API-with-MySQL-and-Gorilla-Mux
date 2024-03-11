package routes

import (
    "github.com/gorilla/mux"
    "github.com/monster0freason/golang-bookstore-management-system/pkg/controllers"
    
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
    // Handle POST request to create a new book
    router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")

    // Handle GET request to retrieve all books
    router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")

    // Handle GET request to retrieve a specific book by ID
    router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")

    // Handle PUT request to update an existing book
    router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")

    // Handle DELETE request to delete a book by ID
    router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}



