package github_actions_for_go_packages

import (
	"dtos"
	"fmt"
	"models"
)

func StringPointer(s string) *string {
	return &s
}

func Test() {
	// Initialize the struct with valid data
	ur := &dtos.UserRequestDto{
		UserName:     StringPointer("username"),
		UserPassword: StringPointer("password123"),
		Email:        StringPointer("test@example.com"),
	}

	u := &models.User{
		UserID:          StringPointer("123"),
		UserName:        StringPointer("username"),
		UserPassword:    StringPointer("password123"),
		Email:           StringPointer("test@example.com"),
		DateTimeCreated: StringPointer("2023-01-01"),
		DateTimeUpdated: StringPointer("2023-01-01"),
	}

	// Print values of u (models.User) and ur (dtos.UserRequestDto)
	fmt.Println("Values of u (models.User):")
	fmt.Printf("UserID: %s\n", *u.UserID)
	fmt.Printf("UserName: %s\n", *u.UserName)
	fmt.Printf("UserPassword: %s\n", *u.UserPassword)
	fmt.Printf("Email: %s\n", *u.Email)
	fmt.Printf("DateTimeCreated: %s\n", *u.DateTimeCreated)
	fmt.Printf("DateTimeUpdated: %s\n", *u.DateTimeUpdated)

	fmt.Println("\nValues of ur (dtos.UserRequestDto):")
	fmt.Printf("UserName: %s\n", *ur.UserName)
	fmt.Printf("UserPassword: %s\n", *ur.UserPassword)
	fmt.Printf("Email: %s\n", *ur.Email)
}
