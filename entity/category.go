package entity

import (
	"log"

	"github.com/oklog/ulid"
)

type CategoryRepository interface {
	Create(category *Category) error
	FindAll() ([]*Category, error)
}

type Category struct {
	Id   string
	Name string
}

func NewCategory(name string) *Category {

	ulidObject, ulidErr := ulid.New(ulid.Now(), nil)
	if ulidErr != nil {
		log.Println("Category - Error when generating Ulid object")
		panic("Category - Error when generating Ulid object")
	}

	return &Category{
		Id:   ulidObject.String(),
		Name: name,
	}
}
