package task

type TaskRepository interface {
	FindById(id TaskId) (*Task, error)
	Save(task *Task) error
}
