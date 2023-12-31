package main_test

import (
	"testing"

	"github.com/MGTheTrain/github-action-workflow-samples/libraries/go/common-lib/src/web/dtos"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	// Initialize the struct with valid data
	u := &dtos.UserRequestDto{
		UserName:     StringPointer("username"),
		UserPassword: StringPointer("password123"),
		Email:        StringPointer("test@example.com"),
	}

	// Call the Validate method
	errors := u.Validate()

	// Assert that there are no errors
	assert.Empty(t, errors, "Expected no validation errors")
}

// Helper function to create string pointers
func StringPointer(s string) *string {
	return &s
}
