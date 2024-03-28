package task

import (
	"app/internal/postgres"
	"testing"
)

func TestEchecOperationOnTaskInPostgresCausedByConnection(t *testing.T) {
	connection := postgres.GetTestConnection()
	repository := NewTaskPsqlRepository(connection)

	idtask, err := repository.Create(Task{Title: "TestSuccess", Description: "A valid task", Status: 1})
	if err == nil {
		t.Errorf("Unable to finish Get task test, creation has failed")
	}

	repository = NewTaskPsqlRepository(nil)

	idTask, err := repository.Create(Task{Title: "TestSuccess", Description: "A valid task", Status: 1})
	if idTask != -1 || err == nil {
		t.Error("Task should not be created")
	}
	_, err = repository.Get(idtask)
	if err == nil {
		t.Error("Error attempted on getting task when database connection isn't ok")
	}

	_, err = repository.GetAll()
	if err == nil {
		t.Error("Error attempted on getting all tasks when database connection isn't ok")
	}

	err = repository.Delete(idtask)
	if err == nil {
		t.Error("Error attempted on deleting a task when database connection isn't ok")
	}

	postgres.ResetTable(postgres.Infos, "task")
}

func TestCreateTaskInPostgres(t *testing.T) {
	connection := postgres.GetTestConnection()
	repository := NewTaskPsqlRepository(connection)

	idTask, err := repository.Create(Task{Title: "TestSuccess", Description: "A valid task", Status: 1})

	if idTask == -1 || err != nil {
		t.Error("Task has not been created")
	}

	postgres.ResetTable(postgres.Infos, "task")
}

func TestEchecCreateTaskInPostgresCausedByInvalidTask(t *testing.T) {
	connection := postgres.GetTestConnection()
	repository := NewTaskPsqlRepository(connection)

	idTask, err := repository.Create(Task{})

	if idTask != -1 || err == nil {
		t.Error("Task should not be created")
	}

	postgres.ResetTable(postgres.Infos, "task")
}

func TestGetExistingTask(t *testing.T) {
	connection := postgres.GetTestConnection()
	repository := NewTaskPsqlRepository(connection)

	idTask, err := repository.Create(Task{Title: "TestSuccess", Description: "A valid task", Status: 1})
	if err == nil {
		t.Errorf("Unable to finish Get task test, creation has failed")
	}

	_, err = repository.Get(idTask)
	if err != nil {
		t.Error("Error occured during Get")
	}

	postgres.ResetTable(postgres.Infos, "task")
}

func TestGetNotFoundTask(t *testing.T) {
	connection := postgres.GetTestConnection()
	repository := NewTaskPsqlRepository(connection)

	_, err := repository.Get(-1)
	if err == nil {
		t.Error("Trying to get a task that doesn't exists should return an error")
	}

	postgres.ResetTable(postgres.Infos, "task")
}

func TestGetAllExistingTask(t *testing.T) {
	connection := postgres.GetTestConnection()
	repository := NewTaskPsqlRepository(connection)

	idTask, err := repository.Create(Task{Title: "TestSuccess", Description: "A valid task", Status: 1})
	if err == nil {
		t.Errorf("Unable to finish Get task test, creation has failed")
	}

	idTask2, err := repository.Create(Task{Title: "TestSuccess", Description: "A valid task", Status: 1})
	if err == nil {
		t.Errorf("Unable to finish Get task test, creation has failed")
	}

	tasks, err := repository.GetAll()
	if err != nil {
		t.Error("Error occured during Get")
	}

	_, ok := tasks[idTask]
	_, ok2 := tasks[idTask2]
	_, ok3 := tasks[idTask2+1]

	if !ok || !ok2 {
		t.Error("Error occured during Get on all tasks")
	}
	if ok3 {
		t.Error("Unattempted another task has been found")
	}

	postgres.ResetTable(postgres.Infos, "task")
}

//func TestDeleteExistingTask
