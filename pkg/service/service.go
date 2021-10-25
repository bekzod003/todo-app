package service

import "todo-app/pkg/repository"

type Authorithation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorithation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
