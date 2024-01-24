package usecase

import (
	"testing"

	"github.com/james-freitas/angelstore-inventory-go/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/oklog/ulid"
)

// MockCategoryRepository is a mock implementation of the CategoryRepository interface
type MockCategoryRepository struct {
	mock.Mock
}

// Create is a method of the MockCategoryRepository struct that creates a new category.
//
// It takes a pointer to an entity.Category object as a parameter.
// It returns an error.
func (m *MockCategoryRepository) Create(category *entity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

// FindAll retrieves all categories from the mock category repository.
//
// It returns a slice of pointers to entity.Category and an error.
func (m *MockCategoryRepository) FindAll() ([]*entity.Category, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Category), args.Error(1)
}

func TestCreateCategoryUseCase_Execute(t *testing.T) {
	
	// Arrange
	categoryRepo := new(MockCategoryRepository)
	createCategoryUseCase := NewCreateCategoryUseCase(categoryRepo)

	input := CreateCategoryInputDto {
		Name: "TestCategory",
	}

	ulidObject, ulidErr := ulid.New(ulid.Now(), nil)
	assert.NoError(t, ulidErr)

	expectedCategory := &entity.Category{
		Id:   ulidObject.String(),
		Name: input.Name,
	}

	categoryRepo.On("Create", expectedCategory).Return(nil)

	// Act
	result, err := createCategoryUseCase.Execute(input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedCategory.Id, result.Id)
	assert.Equal(t, expectedCategory.Name, result.Name)

	// Verify that the Create method of the mock repository was called
	categoryRepo.AssertExpectations(t)
}

