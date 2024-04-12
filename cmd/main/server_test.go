package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/monster0freason/golang-bookstore-management-system/pkg/controllers"
	"github.com/monster0freason/golang-bookstore-management-system/pkg/models"
    
    "encoding/json"
    "io"
	"strings"
)

func TestGetBooks1(t *testing.T) {
    testServer := httptest.NewServer(http.HandlerFunc(controllers.GetBooks))
    defer testServer.Close()

    testClient := testServer.Client()
    fmt.Println((testServer.URL))
    testClient.Get(testServer.URL)

    response, err := testClient.Get(testServer.URL)
    if err != nil {
        t.Error(err)
    }

    // Read the response body
    body, err := io.ReadAll(response.Body)
    if err != nil {
        t.Error(err)
    }

    // Assert that the HTTP status code is equal to http.StatusOK (200)
    assert.Equal(t, http.StatusOK, response.StatusCode)

    // Add assertions for the response body
    // For example, you can assert that the body contains certain data or has a specific format
    // For now, let's just assert that the body is not empty
    assert.NotEmpty(t, body)
}


func TestCreateBook(t *testing.T) {
    testServer := httptest.NewServer(http.HandlerFunc(controllers.CreateBook))
    defer testServer.Close()

    testClient := testServer.Client()
    fmt.Println(testServer.URL)

    // Create a new book object to be sent in the request body
    newBook := models.Book{
        Name:        "Test Book",
        Author:      "Test Author",
        Publication: "Test Publisher",
    }

    // Convert the book object to JSON format
    reqBody, err := json.Marshal(newBook)
    if err != nil {
        t.Error(err)
    }

    // Create a POST request with the JSON request body
    req, err := http.NewRequest("POST", testServer.URL, strings.NewReader(string(reqBody)))
    if err != nil {
        t.Error(err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Send the POST request
    response, err := testClient.Do(req)
    if err != nil {
        t.Error(err)
    }
    defer response.Body.Close()

    // Assert that the HTTP status code is equal to http.StatusCreated (201)
    assert.Equal(t, http.StatusCreated, response.StatusCode)

    // Read the response body
    body, err := io.ReadAll(response.Body)
    if err != nil {
        t.Error(err)
    }

    // Parse the response body into a map
    var responseBody map[string]interface{}
    err = json.Unmarshal(body, &responseBody)
    if err != nil {
        t.Error(err)
    }

    // Assert that the response body contains the created book data
    assert.NotNil(t, responseBody["ID"])
    assert.Equal(t, newBook.Name, responseBody["name"])
    assert.Equal(t, newBook.Author, responseBody["author"])
    assert.Equal(t, newBook.Publication, responseBody["publication"])
	print("TEST")
}


