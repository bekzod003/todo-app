package service

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

func NewService() *Service {
	return &Service{}
}
