package api

import (
	"encoding/json"
	"net/http"
)

//Book type with Name, Title, and ISBN
type Book struct {
	//define the book
}

//ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	return nil
}

//FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	return Book{}
}

//BooksHandleFunc to be used http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
}