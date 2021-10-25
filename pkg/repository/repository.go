package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
