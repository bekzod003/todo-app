package repository

import "github.com/jmoiron/sqlx"

type Authorithation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorithation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
