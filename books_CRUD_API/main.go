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

    var book Book

    json.NewDecoder(r.Body).Decode(&book)
    books = append(books, book)
    json.NewEncoder(w).Encode(books)
    log.Printf("ID: %d, Title: '%s' was added", book.ID, book.Title)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    log.Println("updateBook is called")

    var book Book

    json.NewDecoder(r.Body).Decode(&book)

    for i, item := range books {
        if item.ID == book.ID {
            log.Printf("ID: %d, Title: '%s' was updated", item.ID, item.Title)
            books[i] = book
            log.Printf("ID: %d is now -> '%s'", item.ID, book.Title)
        }
    }
}

func removeBook(w http.ResponseWriter, r *http.Request) {
    log.Println("removeBook is called")

    params := mux.Vars(r)

    id, _ := strconv.Atoi(params["id"])

    for i, item := range books {
        if item.ID == id {
            log.Printf("ID: %d, Title: '%s' was deleted", item.ID, item.Title)
            books = append(books[:i], books[i+1:]...)
        }
    }
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

