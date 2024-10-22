package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

func GetPathFile() string {
	cwd, error := os.Getwd()

	if error != nil {
		log.Fatal(error)
		return ""
	}

	return path.Join(cwd, "tasks.json")
}

// Read task from file
func ReadTasksFromFile() ([]Task, error) {
	filePath := GetPathFile()
	_, error := os.Stat(filePath)
	// Check exist path file
	if os.IsExist(error) {
		fmt.Println("File not found")
		file, error := os.Create(filePath)
		if error != nil {
			fmt.Println("Error creating file:", error)
			return nil, error
		}

		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())
		defer file.Close()

		return []Task{}, nil
	}

	if error != nil {
		fmt.Println("Error opening file:", error)
		return nil, error
	}

	file, error := os.Open(filePath)
	if error != nil {
		fmt.Println("Error opening file:", error)
		return nil, error
	}

	tasks := []Task{}
	defer file.Close()
	error = json.NewDecoder(file).Decode(&tasks)
	if error != nil {
		fmt.Println("Error decoding file:", error)
		return nil, error
	}

	return tasks, error
}

// Write task to file
func WriteTasksToFile(tasks []Task) error {
	filePath := GetPathFile()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}

	defer file.Close()
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return err
	}

	return nil
}
