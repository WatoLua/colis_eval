package task

type TaskRepository interface {
	Create(t Task) (int64, error)
	Get(id int64) (Task, error)
	GetAll() (map[int64]Task, error)
	Update(t Task) error
	Delete(id int64) error
}
