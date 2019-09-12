package api

import (
	"testing"
)

const JSONstring = `{"title":"Cloud Native Go","author":"Robert Thompson","isbn":"0123456789"}`
var book = Book{
	Title: "Cloud Native Go",
	Author: "Robert Thompson",
	ISBN: "0123456789",
}
func TestBookToJSON(t *testing.T) {
	got := string(book.ToJSON())
	want := JSONstring
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(JSONstring)
	got := FromJSON(json)
	want := book
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}