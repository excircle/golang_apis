# "Books CRUD API"

This folder contains source code for a simple Go-based API which demonstrates basic CRUD functionality. This API utilizes "github.com/gorilla/mux" for routing capabilities.

When initialized, this API will launch a simple web server, allows you to visit the below URL list to receive variable API responses.

## API Anatomy

The API responses are as follows:

```
    "http://127.0.0.1:8000/books" - getBooks - GET Request to fetch all books
    "http://127.0.0.1:8000/books/{id}" - getBook - GET Request to fetch single book
    "http://127.0.0.1:8000/books" - addBook - POST Request to add book
    "http://127.0.0.1:8000/books" - updateBook - PUT Request to alter a book
    "http://127.0.0.1:8000/books/{id}" - removeBook - DELETE Request to remove a book
```

### <a href="https://github.com/excircle/golang_apis/blob/master/books_CRUD_API/main.go">main.go</a>

```go
package main

import "log"
import "net/http"
import "encoding/json"
import "strconv"

import "github.com/gorilla/mux"

type Book struct {
    ID      int     `json:"id"`
    Title   string  `json:"title"`
    Author  string  `json:"author"`
    Year    string  `json:"year"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    i, _ := strconv.Atoi(params["id"])

    for _, book := range books {
        if book.ID == i {
            json.NewEncoder(w).Encode(&book)
        }
    }
}

func addBook(w http.ResponseWriter, r *http.Request) {
    log.Println("addBook is called")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    log.Println("updateBook is called")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
    log.Println("removeBook is called")
}

var books []Book

func main() {
    r := mux.NewRouter()

    books = append(books,
        Book{ID: 1, Title: "Pointers", Author: "Mr.Pointers", Year: "2010"},
        Book{ID: 2, Title: "Goroutines", Author: "Mr.Goroutines", Year: "2011"},
        Book{ID: 3, Title: "Routers", Author: "Mr.Routers", Year: "2012"},
        Book{ID: 4, Title: "Concurrency", Author: "Mr.Concurrency", Year: "2013"},
        Book{ID: 5, Title: "Parts", Author: "Mr.Parts", Year: "2014"})

    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/books", addBook).Methods("POST")
    r.HandleFunc("/books", updateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", r))
}
```