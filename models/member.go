// define the member struct
package models

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []int // slice to hold borrowed books ids
}
