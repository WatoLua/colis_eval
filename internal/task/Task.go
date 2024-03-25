package task

import (
	"errors"
	"fmt"
)

type TaskStatus int

const (
	Todo      TaskStatus = 1
	InProgess TaskStatus = 2
	Done      TaskStatus = 3
)

func (ts TaskStatus) String() string {
	if 0 < ts && ts < 4 {
		ts = 0
	}
	return [...]string{
		"Unexpected value",
		"To do",
		"In Progress",
		"Done",
	}[ts]

}

type Task struct {
	Id          int64      `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Status      TaskStatus `json:"status" db:"status"`
}

func (t *Task) GetId() int64 {
	return t.Id
}

func (t *Task) IsValid() (bool, error) {
	if t.Title == "" {
		return false, errors.New("Missing value for Title field")
	}
	if t.Description == "" {
		return false, errors.New("Missing value for Title field")
	}
	if t.Status <= 0 || t.Status >= 4 {
		return false, errors.New("Missing or Wrong value for Status field")
	}

	return true, nil
}

func (t *Task) String() string {
	return fmt.Sprintf("[id = %v, Title=\"%v\", Description=\"%v\", Status=\"%v\"",
		t.Id, t.Title, t.Description, t.Status)
}
