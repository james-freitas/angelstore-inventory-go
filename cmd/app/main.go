package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/james-freitas/angelstore-inventory-go/infra/repository"
	"github.com/james-freitas/angelstore-inventory-go/usecase"
	"github.com/james-freitas/angelstore-inventory-go/web"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/angelstore")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewCategoryRepositoryMySql(db)
	createCategoryUseCase := usecase.NewCreateCategoryUseCase(repository)
	listCategoriesUseCase := usecase.NewListCategoriesUseCase(repository)

	categoryHandlers := web.NewCategoryHandlers(createCategoryUseCase, listCategoriesUseCase)

	r := chi.NewRouter()

	r.Post("/categories", categoryHandlers.CreateCategoryHandler)
	r.Get("/categories", categoryHandlers.ListCategoriesHandler)

	http.ListenAndServe(":8000", r)
}
