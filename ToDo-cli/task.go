// This file will contain the logic for managing tasks in the ToDo CLI application.
// It will include functions for adding, removing, and listing tasks, as well as marking tasks as completed.
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type ToDoTask struct {
	Task        string
	CreatedAt   time.Time
	CompletedAt time.Time
	Completed   bool
}

type Todos []ToDoTask

func (t *Todos) AddTask(task string) {
	local := ToDoTask{
		Task:      task,
		CreatedAt: time.Now(),
		Completed: false,
	}
	*t = append(*t, local)
}

func (t *Todos) RemoveTask(index int) error {
	if index < 0 || index >= len(*t) {
		return fmt.Errorf("Invalid Index value...")
	}

	local := *t
	*t = append(local[:index], local[index+1:]...)

	return nil
}

func (t *Todos) ToggleTask(index int) error {
	if index < 0 || index >= len(*t) {
		return fmt.Errorf("Invalid index value.")
	}

	todo := *t
	task := &todo[index]
	if !task.Completed {
		task.CompletedAt = time.Now()
	}
	task.Completed = !task.Completed

	return nil
}

func (t *Todos) UpdateTask(task string, index int) error {
	if index < 0 || index >= len(*t) {
		return fmt.Errorf("Invalid index.")
	}

	todo := *t
	local := &todo[index]
	local.Task = task

	return nil
}

func (t *Todos) PrintList() {

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)
	fmt.Fprintln(w, "Index\tTask\tCompleted\tCreatedAt\tCompletedAt")
	fmt.Fprintln(w, "-----\t----\t---------\t---------\t-----------")
	for i, val := range *t {
		fmt.Fprintf(w, "%d\t%s\t%t\t%v\t%v\n", i+1, val.Task, val.Completed, val.CreatedAt, val.CompletedAt)
	}
	fmt.Printf("Successfully printed all list of items.\n")

}
