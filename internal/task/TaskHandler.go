package task

import (
	"sync"
)

type TaskHandler struct {
	repo        *TaskRepository
	mutexCreate sync.Mutex
	mutexAccess sync.Mutex
}

var globalHandler TaskHandler = TaskHandler{}

func NewTaskHandler(repo *TaskRepository) *TaskHandler {
	if (globalHandler == TaskHandler{}) {
		globalHandler = TaskHandler{repo: repo}
	}

	return &globalHandler
}

func (handler *TaskHandler) Create(t Task) (int64, error) {
	handler.mutexCreate.Lock()
	defer handler.mutexCreate.Unlock()

	return (*handler.repo).Create(t)
}

func (handler *TaskHandler) Get(id int64) (Task, error) {
	handler.mutexAccess.Lock()
	defer handler.mutexAccess.Unlock()

	return (*handler.repo).Get(id)
}

func (handler *TaskHandler) GetAll() (map[int64]Task, error) {
	handler.mutexAccess.Lock()
	defer handler.mutexAccess.Unlock()

	return (*handler.repo).GetAll()
}

func (handler *TaskHandler) Update(t Task) error {
	handler.mutexAccess.Lock()
	defer handler.mutexAccess.Unlock()

	return (*handler.repo).Update(t)
}

func (handler *TaskHandler) Delete(t Task) error {
	handler.mutexAccess.Lock()
	defer handler.mutexAccess.Unlock()

	return (*handler.repo).Delete(t)
}
