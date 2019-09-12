package api

import (
	"encoding/json"
	"net/http"
)

//Book type with Title, Author, and ISBN
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

//ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

//FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

//Books a slice of Book
var Books = []Book{
	Book{Title: "1Q84", Author: "Haruki Murakami", ISBN: "5656565656"},
	Book{Title: "Things Fall Apart", Author: "Chinua Achebe", ISBN: "1212121212"},
}

//BooksHandleFunc to be used http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(b)
}
