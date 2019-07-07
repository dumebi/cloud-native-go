package main

import (
	"fmt"
	"jude/cloud-native-go/api"
	"net/http"
	"os"
	// "github.com/dumebi/cloud-native-go/api"
)

func main() {
	http.HandleFunc("/", api.HelloHandleFunc)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello cloud native Go!")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
