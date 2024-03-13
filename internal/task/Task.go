package task

import "fmt"

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
	id          int        `json:id`
	Title       string     `json:title`
	Description string     `description`
	Status      TaskStatus `json:status`
}

func (t *Task) GetId() int {
	return t.id
}

func (t *Task) IsValidTask() bool {
	return (t.id > -1 &&
		t.Title != "" &&
		t.Description != "" &&
		0 < t.Status && t.Status < 4)
}

func (t *Task) String() string {
	return fmt.Sprintf("[id = %v, Title=\"%v\", Description=\"%v\", Status=\"%v\"",
		t.id, t.Title, t.Description, t.Status)
}
