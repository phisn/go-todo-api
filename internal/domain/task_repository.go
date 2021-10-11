package domain

type Repository interface {
	Get(id int) *Task
	Add(task *Task) error
	Remove(id int) error
	GetAll() []*Task
	RemoveAll() error
}
