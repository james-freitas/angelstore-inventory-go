package repository

import (
	"database/sql"
	"log"

	"github.com/james-freitas/angelstore-inventory-go/entity"
)

type CategoryRepositoryMySql struct {
	DB *sql.DB
}

func NewCategoryRepositoryMySql(db *sql.DB) *CategoryRepositoryMySql {
	return &CategoryRepositoryMySql{
		DB: db,
	}
}

func (r *CategoryRepositoryMySql) Create(category *entity.Category) error {
	_, err := r.DB.Exec("Insert into category (id, name) values (?,?)",
		category.Id, category.Name)
	if err != nil {
		log.Println("CategoryRepositoryMySql - Error when calling repository to save the category ", err)
		return err
	}
	return nil
}

func (r *CategoryRepositoryMySql) FindAll() ([]*entity.Category, error) {
	rows, err := r.DB.Query("select id, name from category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
