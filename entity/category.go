package entity

import "github.com/google/uuid"

type CategoryRepository interface {
	Create(category *Category) error
	FindAll() ([]*Category, error)
}

type Category struct {
	Id   string
	Name string
}

func NewCategory(id string, name string) *Category {
	return &Category{
		Id:   uuid.New().String(),
		Name: name,
	}
}
