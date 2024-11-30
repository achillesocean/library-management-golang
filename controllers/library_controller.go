// Map user commands (like "Add a book", "Borrow a book") to service methods (AddBook, BorrowBook, etc.).
//Handle user prompts (e.g., "Enter book title:") and pass the inputs to the service layer.
//Format and display outputs (e.g., "Book borrowed successfully!" or "Book not found.").
// Bridge between UI and Logic...even if the UI is console

// will the controller run forever, waiting for commands/prompts?
package controllers

import (
	"fmt"
	"project-01/models"
	"project-01/services"
	"strings"
)

func RunController(library *services.Library) {
	for {
		fmt.Println("Enter a command (add, remove, borrow, return, list, exit):")

		var command string
		fmt.Scan(&command)
		command = strings.ToLower(command)

		switch command {
		case "add":
			// need: the book to add. how do I take these as inputs?
			var bookID int
			var title, author string

			fmt.Println("Enter Book ID (integer):")
			_, err := fmt.Scan(&bookID)
			if err != nil || bookID <= 0 {
				fmt.Println("Invalid Book ID. Please enter a positive integer.")
				break
			}

			fmt.Println("Enter Title:")
			fmt.Scan(&title)
			if len(title) == 0 {
				fmt.Println("Title cannot be empty.")
				break
			}

			fmt.Println("Enter Author:")
			fmt.Scan(&author)
			if len(author) == 0 {
				fmt.Println("Author cannot be empty.")
				break
			}

			// create a book
			// import the Book struct, and make an instance of it
			// type newBook models.Book {
			// 	// how do I initialize the fields here?
			// }

			newBook := models.Book{
				ID:     bookID,
				Title:  title,
				Author: author,
				Status: "available",
			}

			if err := library.AddBook(newBook); err != nil {
				fmt.Printf("Failed to add book: %s\n", err)
			} else {
				fmt.Println("New book added successfully!")
			}

		case "remove":
			// remove book: provided book id
			fmt.Println("Please provide the ID of the book you want to remove:")
			var removeID int
			_, err := fmt.Scan(&removeID)
			if err != nil || removeID <= 0 {
				fmt.Println("Invalid Book ID. Please enter a positive integer.")
				break
			}

			if err := library.RemoveBook(removeID); err != nil {
				fmt.Printf("Failed to remove book: %s\n", err)
			} else {
				fmt.Println("Book removed successfully!")
			}

		case "return":
			fmt.Println("Enter the Book ID to return:")
			var returnBookID, returnMemberID int
			_, err := fmt.Scan(&returnBookID)
			if err != nil || returnBookID <= 0 {
				fmt.Println("Invalid Book ID. Please enter a positive integer.")
				break
			}
			_, err = fmt.Scan(&returnMemberID)
			if err != nil || returnMemberID <= 0 {
				fmt.Println("Invalid Member ID. Please enter a positive integer.")
				break
			}

			library.ReturnBook(returnBookID, returnMemberID)

		case "list":
			// list available books
			availableBooks := library.ListAvailableBooks()
			if len(availableBooks) == 0 {
				fmt.Println("No books available.")
			} else {
				fmt.Println("Available books:")
				for _, book := range availableBooks {
					fmt.Printf("ID: %d, Title: %s, Author: %s \n", book.ID, book.Title, book.Author)
				}
			}
		case "exit":
			fmt.Println("Exiting the LMS. Goodbye!")
			return
		default:
			fmt.Println("Invalid command. Please try again")
		}
	}
}
