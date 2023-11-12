package entity

import "testing"

// TestNewCategory is a unit test function for the NewCategory function.
func TestNewCategory(t *testing.T) {
	// Set up test data
	name := "CategoryName"

	// Call the function being tested
	category := NewCategory(name)

	// Assert the result
	if category.Name != name {
		t.Errorf("Expected %s, got %s", name, category.Name)
	}
	// Assert the category Id was assigned
	if category.Id == "" {
		t.Errorf("Expected non-nil Id, got nil")
	}
}
