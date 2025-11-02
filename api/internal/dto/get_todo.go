package dto

import "github.com/alexeevayaan/portfolio-react-golang-todo/api/internal/domain"



type GetTodoOutput struct{
	domain.Todo
}

type GetTodoInput struct{
	Id string
}