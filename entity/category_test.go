package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// MockCategoryRepository is a mock implementation of CategoryRepository for testing purposes.
type MockCategoryRepository struct {
	Categories []*Category
}

func (m *MockCategoryRepository) Create(category *Category) error {
	m.Categories = append(m.Categories, category)
	return nil
}

func (m *MockCategoryRepository) FindAll() ([]*Category, error) {
	return m.Categories, nil
}

func TestNewCategory(t *testing.T) {
	name := "TestCategory"
	category := NewCategory(name)

	// Ensure that id is generated
	assert.NotEmpty(t, category.Id)

	// Ensure that Name is set correctly
	assert.Equal(t, name, category.Name)
}

func TestCategoryRepositoryCreate(t *testing.T) {
	
	// Setup
	repo := &MockCategoryRepository{}
	name := "TestCategory"
	category := NewCategory(name)

	// Test
	err := repo.Create(category)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, repo.Categories, 1)
	assert.Equal(t, category, repo.Categories[0])	
}

func TestCategoryRepositoryFindAll(t *testing.T) {
	
	// Setup
	repo := &MockCategoryRepository{}
	name := "TestCategory"
	category := NewCategory(name)
	_ = repo.Create(category)

	// Test
	categories, err := repo.FindAll()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, categories, 1)
	assert.Equal(t, category, categories[0])
}