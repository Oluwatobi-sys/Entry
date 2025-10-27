package main

import (
	"fmt"
)

func main() {
	var firstName, lastName, department, faculty, level string
	var age int

	// Request inputs from the user
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Print("Enter your department: ")
	fmt.Scanln(&department)

	fmt.Print("Enter your faculty: ")
	fmt.Scanln(&faculty)

	fmt.Print("Enter your level: ")
	fmt.Scanln(&level)

	// Print user details
	fmt.Println("\n----- User Information -----")
	fmt.Printf("Name: %s %s\n", firstName, lastName)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Department: %s\n", department)
	fmt.Printf("Faculty: %s\n", faculty)
	fmt.Printf("Level: %s\n", level)
}
