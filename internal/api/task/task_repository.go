package task

import (
	"awesomeProject/internal/api/task/fake"
	"awesomeProject/internal/domain"
)

func NewRepository() domain.Repository {
	return fake.FakeTaskRepository{}
}
