package usecase

import "github.com/james-freitas/angelstore-inventory-go/entity"

type ListCategoriesOutputDto struct {
	Id 		string
	Name 	string
}

type ListCategoriesUseCase struct {
	CategoryRepository entity.CategoryRepository
}

func NewListCategoriesUseCase(categoryRepository entity.CategoryRepository) *ListCategoriesUseCase {
	return &ListCategoriesUseCase{
		CategoryRepository: categoryRepository,
	}
}

func (u *ListCategoriesUseCase) Execute() ([]*ListCategoriesOutputDto, error) {
	categories, err := u.CategoryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesOutput []*ListCategoriesOutputDto
	for _, category := range categories {
		categoriesOutput = append(categoriesOutput, &ListCategoriesOutputDto{
			Id: category.Id,
			Name: category.Name,
		})
	}
	return categoriesOutput, nil
}