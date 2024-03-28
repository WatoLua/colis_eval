package main

import (
	"net/http"

	"app/internal/postgres"
	"app/internal/task"

	"github.com/go-chi/chi/v5"
)

func main() {

	infos := postgres.DBInfos{
		Host:     "172.19.0.2",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Dbname:   "eval",
	}

	connection := postgres.GetConnection(infos)
	defer postgres.CloseConnection()
	repository := task.NewTaskPsqlRepository(connection)
	taskHandler := task.NewTaskHandler(&repository)

	r := chi.NewRouter()

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
