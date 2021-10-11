package fake

import (
	"awesomeProject/internal/domain"
)

type FakeTaskRepository struct {
}

var tasks = []*domain.Task{
	{Id: 0, Description: "Description0", Title: "Title0"},
	{Id: 1, Description: "Description1", Title: "Title1"},
	{Id: 2, Description: "Description2", Title: "Title2"},
	{Id: 3, Description: "Description3", Title: "Title3"},
	{Id: 4, Description: "Description4", Title: "Title4"},
}

var counterId = len(tasks)

func (f FakeTaskRepository) Get(id int) *domain.Task {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			return tasks[i]
		}
	}
	return nil
}

func (f FakeTaskRepository) Add(task *domain.Task) error {
	task.Id = counterId
	counterId++

	tasks = append(tasks, task)
	return nil
}

func (f FakeTaskRepository) Remove(id int) error {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	return nil
}

func (f FakeTaskRepository) GetAll() []*domain.Task {
	return tasks
}

func (f FakeTaskRepository) RemoveAll() error {
	tasks = []*domain.Task{}
	return nil
}
