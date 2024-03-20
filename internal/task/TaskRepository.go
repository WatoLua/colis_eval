package task

type TaskRepository interface {
	Create(t Task) (bool, error)
	Get(id int64) (Task, error)
	GetAll() (map[int64]Task, error)
	Update(t Task) bool
	Delete(t Task) bool
}