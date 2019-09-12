package api

import (
	"testing"
)

func TestBookToJSON(t *testing.T) {
	book := Book{
		Title: "Cloud Native Go",
		Author: "Robert Thompson",
		ISBN: "0123456789",
	}
	got := string(book.ToJSON())
	want := `{"Title":"Cloud Native Go","Author":"Robert Thompson","ISBN":"0123456789"}`
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}