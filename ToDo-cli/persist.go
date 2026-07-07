// This file will contain the logic for saving and loading tasks from a persistent storage, such as a file or a database.
// It will include functions for reading and writing tasks to the storage medium.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(t Todos, filename string) error {
	var err error
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	return err
}

func Load(filename string) (t Todos, err error) {
	if filename == "" {
		err = fmt.Errorf("Invalid filename.")
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, t)
	return
}
