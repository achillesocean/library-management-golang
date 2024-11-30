package main

import (
	"fmt"
	"project-01/controllers"
	"project-01/models"
	"project-01/services"
)

func main() {
	library := &services.Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
	fmt.Println("Welcome to the Library Management System!")
	controllers.RunController(library)
}
