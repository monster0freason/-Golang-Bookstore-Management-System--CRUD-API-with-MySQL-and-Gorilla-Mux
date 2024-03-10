package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/monster0freason/golang-bookstore-management-system/pkg/routes"
)

func main() {
    r := mux.NewRouter()

    // Register bookstore routes
    routes.RegisterBookstoreRoutes(r)

    // Handle functions using controllers
    http.Handle("/", r)

    // Create server
    log.Fatal(http.ListenAndServe("localhost:9010", r))
}