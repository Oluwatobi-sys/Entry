package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
	Faculty    string `json:"faculty"`
	Level      string `json:"level"`
}

func main() {
	var user User

	// Request inputs from the user
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&user.FirstName)

	fmt.Print("Enter your last name: ")
	fmt.Scanln(&user.LastName)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&user.Age)

	fmt.Print("Enter your department: ")
	fmt.Scanln(&user.Department)

	fmt.Print("Enter your faculty: ")
	fmt.Scanln(&user.Faculty)

	fmt.Print("Enter your level: ")
	fmt.Scanln(&user.Level)

	// Print user details
	fmt.Println("\n----- User Information -----")
	fmt.Printf("Name: %s %s\n", user.FirstName, user.LastName)
	fmt.Printf("Age: %d\n", user.Age)
	fmt.Printf("Department: %s\n", user.Department)
	fmt.Printf("Faculty: %s\n", user.Faculty)
	fmt.Printf("Level: %s\n", user.Level)

	// Create data directory if it doesn't exist
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Create a JSON file inside data/
	file, err := os.Create("data/user_data.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode user info as JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(user)
	if err != nil {
		fmt.Println("Error writing JSON:", err)
		return
	}

	fmt.Println("\n✅ User data saved successfully to data/user_data.json")

}
