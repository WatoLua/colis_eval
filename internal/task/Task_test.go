package task

import "testing"

func TestValidTask(t *testing.T) {
	task := Task{Title: "Title", Description: "Description", Status: 1}
	if ok, _ := task.IsValid(); !ok {
		t.Error("Task should be valid")
	}
}

func TestInvalidTaskOnTitle(t *testing.T) {
	task := Task{Title: "", Description: "Description", Status: 2}
	if ok, _ := task.IsValid(); ok {
		t.Error("Task should not be valid, title is empty")
	}
}

func TestInvalidTaskOnDescription(t *testing.T) {
	task := Task{Title: "Title", Description: "", Status: 2}
	if ok, _ := task.IsValid(); ok {
		t.Error("Task should not be valid, description is empty")
	}
}

func TestInvalidTaskOnStatus(t *testing.T) {
	task := Task{Title: "Title", Description: "Description", Status: 0}
	if ok, _ := task.IsValid(); ok {
		t.Error("Task should not be valid, status must be between values 1 and 3")
	}
}
