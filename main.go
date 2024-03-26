package main

import (
	"net/http"

	postgres "app/internal/connections"
	"app/internal/task"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	connection := postgres.GetPostgresConnection()
	repository := task.NewTaskPsqlRepository(connection)
	taskHandler := task.NewTaskHandler(&repository)

	r.Get("/isAlive", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is alive\n"))
	})

	r.Post("/task", taskHandler.PostTask)
	r.Get("/task/{id}", taskHandler.GetTask)
	r.Get("/tasks", taskHandler.GetAllTasks)
	r.Put("/task/{id}", taskHandler.PutTask)
	r.Delete("/task/{id}", taskHandler.DeleteTask)

	http.ListenAndServe(":7800", r)
}
