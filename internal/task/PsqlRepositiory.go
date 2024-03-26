package task

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TaskPsqlRepository struct {
	connection *sqlx.DB
}

func NewTaskPsqlRepository(connection *sqlx.DB) TaskRepository {
	psqlRepo := TaskPsqlRepository{connection: connection}
	var repo TaskRepository = &psqlRepo
	return repo
}

func (repo *TaskPsqlRepository) Create(t Task) (int64, error) {

	query := "insert into task (id, title, description, status)" +
		"values (default, :title, :description, :status) returning id"

	params := map[string]interface{}{
		"title":       t.Title,
		"description": t.Description,
		"status":      t.Status,
	}

	rows, err := repo.connection.NamedQuery(query, params)

	if err != nil {
		return -1, err
	}

	defer rows.Close()
	var id int64
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return -1, errors.New("Unable to scan id of the created task")
		}
	} else {
		return -1, errors.New("no id returned in insert")
	}
	return id, nil
}

func (repo *TaskPsqlRepository) Get(id int64) (Task, error) {

	t := Task{}

	query := "select * from task where id = :id"
	params := map[string]interface{}{
		"id": id,
	}

	rows, err := repo.connection.NamedQuery(query, params)

	if err != nil {
		return Task{}, err
	}

	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&t); err != nil {
			return Task{}, errors.New("Unable to scan sql result as a Task object")
		}
		return t, nil
	}

	return Task{}, nil
}

func (repo *TaskPsqlRepository) GetAll() (map[int64]Task, error) {
	var tasks map[int64]Task = make(map[int64]Task)

	query := "select * from task"

	rows, err := repo.connection.Queryx(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		t := Task{}
		if err := rows.StructScan(&t); err != nil {
			fmt.Println(fmt.Errorf("Unable to scan a task. Detailed error : %w", err))
		} else {
			tasks[t.Id] = t
		}
	}

	return tasks, nil
}

func (repo *TaskPsqlRepository) Update(t Task) error {

	query := "update task set title = :title, description = :description, status = :status where id = :id"

	params := map[string]interface{}{
		"id":          t.Id,
		"title":       t.Title,
		"description": t.Description,
		"status":      t.Status,
	}

	res, err := repo.connection.NamedExec(query, params)

	if err != nil {
		return errors.New("Unable to get the task in database")
	}

	i, err := res.RowsAffected()

	if err != nil {
		return errors.New("Error occured for affected rows for Task object")
	}

	if i != 1 {
		return fmt.Errorf("Unexpected number of updated rows : %v", i)
	}

	return nil
}

func (repo *TaskPsqlRepository) Delete(id int64) error {
	query := "delete from task where id = :id"

	params := map[string]interface{}{
		"id": id,
	}

	res, err := repo.connection.NamedExec(query, params)

	if err != nil {
		return errors.New("Unable to get the task in database")
	}

	i, err := res.RowsAffected()

	if err != nil {
		return errors.New("Error occured for affected rows for Task object")
	}

	if i != 1 {
		return fmt.Errorf("Unexpected number of deleted rows : %v", i)
	}

	return nil
}
