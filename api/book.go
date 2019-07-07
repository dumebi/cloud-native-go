package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Book type with Name, Author and ISBN
type Book struct {
	// define the book
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

// ToJSON for Marshalling of the book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON for Unmarshalling of the book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var books = map[string]Book{
	"12345678": Book{Title: "Cloud Native Go", Author: "Jude Dike", ISBN: "12345678"},
	"12638277": Book{Title: "Another book", Author: "Jude Dike", ISBN: "12638277"},
}

// BooksHandleFunc to be used as HandleFunc for the books api
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Book Created"))
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported Request Method."))
	}
}

// BookHandleFunc to be used as http.HandleFunc for Book API
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

// AllBooks returns a slice of all books
func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

// CreateBook creates a new Book if it does not exist
func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}

// GetBook returns the book for a given ISBN
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

// UpdateBook updates an existing book
func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if exists {
		books[isbn] = book
	}
	return exists
}

// DeleteBook removes a book from the map by ISBN key
func DeleteBook(isbn string) {
	delete(books, isbn)
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
