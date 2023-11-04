package web

import (
	"encoding/json"
	"log"
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
		log.Println("Handler - Error when using json decoding on CreateCategoryInputDto: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := c.CreateCategoryUseCase.Execute(input)
	if err != nil {
		log.Println("Handler - Error when trying to list from the database: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (c *CategoryHandlers) ListCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	output, err := c.ListCategoriesUseCase.Execute()
	if err != nil {
		log.Println("Handler - Error when trying to list from the database: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}