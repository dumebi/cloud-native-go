package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJson(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "Jude Dike", ISBN: "1234-5678"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go","author":"Jude Dike","isbn":"1234-5678"}`, string(json), "Book Json marshalling wrong")
}

func TestBookFromJson(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"Jude Dike","isbn":"1234-5678"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "Jude Dike", ISBN: "1234-5678"}, book, "Book Json unmarshalling wrong")
}
