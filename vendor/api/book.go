package api

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

//Book type with Title, Author, and ISBN
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

//Books a slice of Book
var Books = []Book{
	Book{Title: "1Q84", Author: "Haruki Murakami", ISBN: "5656565656"},
	Book{Title: "Things Fall Apart", Author: "Chinua Achebe", ISBN: "1212121212"},
}

//BooksHandleFunc to be used http.HandleFunc for Book API
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
			w.Header().Add("Location", "/api/books/" +isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

//BookHandleFunc to be used http.HandleFunc for Book API
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/book/"):]
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
	}

}

//JSON FUNCTIONS

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

func writeJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(b)
}

//BOOK FUNCTIONS

//GetBook gets a book
func GetBook(isbn string) (Book, bool) {
	for _, v := range Books {
		if v.ISBN == isbn {
			return v, true
		}
	}
	return Book{}, false
}

//CreateBook makes a book
func CreateBook(b Book) (string, bool) {
	Books = append(Books, b)
	return b.ISBN, true
}

//AllBooks returns the books
func AllBooks() []Book {
	return Books
}

//UpdateBook updates the book
func UpdateBook(isbn string, book Book) bool {
	for _, v := range Books {
		if v.ISBN == isbn {
			v = book
			return true
		}
	}
	return false
}

//DeleteBook deletes a book
func DeleteBook(isbn string) {
	var index int
	for i, v := range Books {
		if v.ISBN == isbn {
			index = i
			break
		}
	}
	Books = append(Books[0:index], Books[index:]...)
}