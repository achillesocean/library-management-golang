// place the interface in here. the implementation will also reside here?
package services

import (
	"fmt"
	"project-01/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

// implement the interface in a Library struct
// that is different from implementing the methods in the inteface?

// create the Library struct here or in the main.go?

type Library struct {
	// implement the methods here?
	// where do I keep the books? inside a field in this struct here- a map wit book id and book. also a field to store the members.
	Members map[int]models.Member
	Books   map[int]models.Book // how do I create a map?
	// I shouldn't initialize the Books field map directly here?

}

func (l *Library) AddBook(book models.Book) error {
	// use the pointer receiver so that the changes persist
	// add to the books field.
	if _, exists := l.Books[book.ID]; exists {
		return fmt.Errorf("a book with ID %d already exists", book.ID)
	}

	l.Books[book.ID] = book // how do I refer to an instance of the current struct instance? like self. in python?
	return nil
}

func (l *Library) RemoveBook(bookID int) error {
	// how to check if the bookID even is inside the books

	if _, exists := l.Books[bookID]; !exists {
		return fmt.Errorf("Book with ID %d does not exist", bookID)
	}
	delete(l.Books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	// check to see if bookID exists in books, and the same for memberID, and if the books is available or not. then just update their status.
	book, bookExists := l.Books[bookID]
	member, memberExists := l.Members[memberID]

	// where do I include the error?
	if !bookExists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}

	if !memberExists {
		return fmt.Errorf("member with ID %d not found", memberID)
	}

	// Update book status and member's borrowed books
	book.Status = "borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, bookID)
	l.Members[memberID] = member

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	member, memberExists := l.Members[memberID]

	if !bookExists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}

	if !memberExists {
		return fmt.Errorf("member with ID %d not found", memberID)
	}

	for i, borrowedBookID := range member.BorrowedBooks {
		if borrowedBookID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	// Update book status
	book.Status = "available"
	l.Books[bookID] = book
	l.Members[memberID] = member

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	ret := make([]models.Book, 0)
	// should I recreate this list everytime though?
	for _, book := range l.Books {
		if book.Status == "available" {
			ret = append(ret, book)
		}
	}
	return ret
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return []models.Book{} // Return an empty slice if the member doesn't exist
	}
	ret := make([]models.Book, 0)
	for bookID := range member.BorrowedBooks {
		ret = append(ret, l.Books[bookID])
	}
	return ret
}
