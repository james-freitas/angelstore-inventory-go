package web

import (
	"encoding/json"
	"net/http"

	"github.com/james-freitas/angelstore-inventory-go/usecase"
)

type CategoryHandlers struct {
	CreateCategoryUseCase *usecase.CreateCategoryUseCase
	ListCategoriesUseCase *usecase.ListCategoriesUseCase
}

func NewCategoryHandlers(createCategoryUseCase *usecase.CreateCategoryUseCase, listCategoriesUseCase *usecase.ListCategoriesUseCase) *CategoryHandlers {
	return &CategoryHandlers{
		CreateCategoryUseCase: createCategoryUseCase,
		ListCategoriesUseCase: listCategoriesUseCase,
	}
}

func (c *CategoryHandlers) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateCategoryInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := c.CreateCategoryUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}