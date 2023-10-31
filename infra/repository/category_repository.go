package repository

import (
	"database/sql"

	"github.com/james-freitas/angelstore-inventory-go/entity"
)

type CategoryRepositoryPostgres struct {
	DB *sql.DB
}

func NewCategoryRepositoryPostgres(db *sql.DB) *CategoryRepositoryPostgres {
	return &CategoryRepositoryPostgres{
		DB: db,
	}
}

func (r *CategoryRepositoryPostgres) Create(category *entity.Category) error {
	_, err := r.DB.Exec("Insert into category (id, name) values (?,?,?)",
	category.Id, category.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepositoryPostgres) FindAll() ([]*entity.Category, error) {
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