package task

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	repo  *TaskRepository
	mutex sync.Mutex
}

var globalHandler TaskHandler = TaskHandler{}

func NewTaskHandler(repo *TaskRepository) *TaskHandler {
	if (globalHandler == TaskHandler{}) {
		globalHandler = TaskHandler{repo: repo}
	}
	return &globalHandler
}

func (handler *TaskHandler) PostTask(w http.ResponseWriter, r *http.Request) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if ok, err := task.IsValid(); !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := (*handler.repo).Create(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"id":      id,
		"message": "Task succesfully created",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func (handler *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	id := chi.URLParam(r, "id")
	idTask, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := (*handler.repo).Get(idTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func (handler *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	tasks, err := (*handler.repo).GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func (handler *TaskHandler) PutTask(w http.ResponseWriter, r *http.Request) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	id := chi.URLParam(r, "id")
	idTask, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var task Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if ok, err := task.IsValid(); !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.Id = idTask
	err = (*handler.repo).Update(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task, err = (*handler.repo).Get(idTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if ok, err := task.IsValid(); !ok {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}

func (handler *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	id := chi.URLParam(r, "id")
	idTask, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := (*handler.repo).Get(idTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if ok, _ := task.IsValid(); !ok {
		http.Error(w, errors.New("Task not found").Error(), http.StatusNotFound)
		return
	}

	err = (*handler.repo).Delete(idTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"id":      id,
		"message": "Task successfully deleted",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusOK)
}
