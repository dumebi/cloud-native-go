package api

import (
	"encoding/json"
	"net/http"
)

// Book type with Name, Author and ISBN
type Book struct {
	// define the book
	Title  string
	Author string
	ISBN   string
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

// BooksHandleFunc to be used as HandleFunc for the books api
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

}
