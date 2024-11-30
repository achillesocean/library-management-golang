// define the book struct
package models

type Book struct {
	ID     int
	Title  string
	Author string
	Status string // available or borrowed
}

// what am I gonna be using these constants for?
const (
	StatusAvailable = "available"
	StatusBorrowed  = "borrowed"
)
