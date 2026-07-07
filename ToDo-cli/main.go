package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type cmd_args struct {
	add    string
	update string
	del    int
	toggle int
	list   bool
}

func getCmdFlags() (cmd *cmd_args) {

	flag.StringVar(&cmd.add, "Add", "", "Add item to the ToDo list.")
	flag.StringVar(&cmd.update, "Update", "", "Update an item by its index in ToDo list. Format 'id:New_string'")
	flag.IntVar(&cmd.del, "Remove", 0, "Specify index of the item to be removed from ToDo list.")
	flag.IntVar(&cmd.toggle, "Toggle", 0, "Specify index of an item to toggle.")
	flag.BoolVar(&cmd.list, "List", false, "Type List for listing all ToDos.")
	flag.Parse()

	return cmd
}

func (c *cmd_args) Execute(todo *Todos) (err error) {

	switch {
	case c.add != "":
		fmt.Printf("Executing Add..")
		todo.AddTask(c.add)
	case c.del != -1:
		fmt.Printf("Executing Delete..")
		todo.RemoveTask(c.del)
	case c.update != "":
		fmt.Printf("Executing Update..")
		s := strings.Split(c.update, ":")
		index, err := strconv.Atoi(s[0])
		if err != nil {
			return err
		}
		todo.UpdateTask(s[1], index)
	case c.toggle != -1:
		fmt.Printf("Executing Toggle..")
		err = todo.ToggleTask(c.toggle)
	case c.list:
		fmt.Printf("Executing printList..")
		todo.PrintList()
	}

	return
}

func main() {
	var err error
	todo := Todos{}
	fmt.Println("This is a simple ToDo CLI application.")

	filename := "todos.jason"
	todo, err = Load(filename)
	if err != nil {
		fmt.Errorf("Unable to load file '%s'.", filename)
		fmt.Println("Continueing with regular todo tasks...")
	}

	cmd := getCmdFlags()
	err = cmd.Execute(&todo)
	if err != nil {
		fmt.Printf("Error received : %v", err)
	}
	err = Save(todo, filename)
	if err != nil {
		fmt.Println("Unable to save todo list to file.")
	}

	return
}
