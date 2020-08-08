# "Books CRUD API"

This folder contains source code for a simple Go-based API which demonstrates basic CRUD functionality. This API utilizes "github.com/gorilla/mux" for routing capabilities.

When initialized, this API will launch a simple web server which allows you to visit the below URLs and see their respective API response.

## API Anatomy

The API responses are as follows:

```
    getBooks -  "http://127.0.0.1:8000/books" - GET Request to fetch all books
    getBook -  "http://127.0.0.1:8000/books/{id}" - GET Request to fetch single book
    addBook -  "http://127.0.0.1:8000/books" - POST Request to add book
    updateBook -  "http://127.0.0.1:8000/books" - PUT Request to alter a book
    removeBook -  "http://127.0.0.1:8000/books/{id}" - DELETE Request to remove a book
```

#### getBooks curl Example

```bash
curl --request GET http://localhost:8000/books
```

#### getBook curl Example

```bash
curl --request GET http://localhost:8000/books/3
```

#### addBook curl Example

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"id": 7,"title": "C is really old","author": "Mr.C","year": "2016"}' \
  http://localhost:8000/books
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