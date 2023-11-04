package usecase

import (
	"log"

	"github.com/james-freitas/angelstore-inventory-go/entity"
)

type CreateCategoryInputDto struct {
	Name string
}

type CreateCategoryOutputDto struct {
	Id string
	Name string
}

type CreateCategoryUseCase struct {
	CategoryRepository entity.CategoryRepository
}
func NewCreateCategoryUseCase(categoryRepository entity.CategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		CategoryRepository: categoryRepository,
	}
}

func (u *CreateCategoryUseCase) Execute(input CreateCategoryInputDto) (*CreateCategoryOutputDto, error) {
	category := entity.NewCategory(input.Name)
	err := u.CategoryRepository.Create(category)
	if err != nil {
		log.Println("CreateCategoryUseCase - Error when calling repository to save the category ", err)
		return nil, err
	}
	return &CreateCategoryOutputDto{
		Id: category.Id,
		Name: category.Name,
	}, nil
}