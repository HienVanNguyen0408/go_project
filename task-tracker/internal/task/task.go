package task

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

type ActionStatus string

const (
	DELETE ActionStatus = "delete"
	UPDATE ActionStatus = "update"
	ADD    ActionStatus = "add"
)

type Task struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func GetTaskStatus(status string) TaskStatus {
	switch status {
	case "todo":
		return TASK_STATUS_TODO
	case "in-progress":
		return TASK_STATUS_IN_PROGRESS
	case "done":
		return TASK_STATUS_DONE
	}
	return TASK_STATUS_TODO
}

func GetColorTextTaskStatus(status TaskStatus) string {
	switch status {
	case TASK_STATUS_TODO:
		return "#3C3C3C"
	case TASK_STATUS_IN_PROGRESS:
		return "202"
	case TASK_STATUS_DONE:
		return "#04B575"
	default:
		return "#3C3C3C"
	}
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		fmt.Println("Error reading tasks from file:", err)
		return err
	}

	// Get ID increment
	id := GetIDInrement(tasks)
	task := NewTask(id, description)
	tasks = append(tasks, *task)
	PrintCmd(ADD, id)
	return WriteTasksToFile(tasks)
}

func GetIDInrement(tasks []Task) int64 {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}

func UpdateTask(id int64, status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	resultTasks := []Task{}
	existTask := false
	for _, task := range tasks {
		if task.ID == id {
			existTask = true
			task.Status = status
			task.UpdatedAt = time.Now()
		}
	}

	if !existTask {
		fmt.Println("Task update not found")
	}
	PrintCmd(UPDATE, id)

	return WriteTasksToFile(resultTasks)
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	resultTasks := []Task{}
	existTask := false
	for _, task := range tasks {
		if task.ID == id {
			existTask = true
			task.Description = description
			task.UpdatedAt = time.Now()
		}
	}

	if !existTask {
		fmt.Println("Task update not found")
	}
	PrintCmd(UPDATE, id)

	return WriteTasksToFile(resultTasks)
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	resultTasks := []Task{}
	for _, task := range tasks {
		if task.ID != id {
			resultTasks = append(resultTasks, task)
		}
	}

	PrintCmd(DELETE, id)

	return WriteTasksToFile(resultTasks)
}

func PrintCmd(actionStatus ActionStatus, id int64) {
	idPrint := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66")).
		Render(fmt.Sprintf("(ID: %d)", id))

	switch actionStatus {
	case DELETE:
		fmt.Printf("\nTask deleted successfully: %s\n\n", idPrint)
		break
	case ADD:
		fmt.Printf("\nTask Added successfully: %s\n\n", idPrint)
		break
	case UPDATE:
		fmt.Printf("\nTask Updated successfully: %s\n\n", idPrint)
		break
	}
}

func GetListTask(status TaskStatus) error {
	tasks, err := ReadTasksFromFile()

	if err != nil {
		fmt.Println("Error reading tasks from file:", err)
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	filteredTasks := GetTaskByStatus(tasks, status)
	if len(filteredTasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	fmt.Println()
	fmt.Println(
		lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC66")).
			MarginBottom(1).
			Render(fmt.Sprintf("Tasks (%s)", status)))
	for _, task := range filteredTasks {
		formattedId := lipgloss.NewStyle().
			Bold(true).
			Width(5).
			Render(fmt.Sprintf("ID:%d", task.ID))
		formattedStatus := lipgloss.NewStyle().
			Bold(true).
			Width(12).
			Foreground(lipgloss.Color(GetColorTextTaskStatus(task.Status))).
			Render(string(task.Status))

		relativeUpdatedTime := task.UpdatedAt.Format("2006-01-02 15:04:05")

		taskStyle := lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("#3C3C3C")).
			Render(fmt.Sprintf("%s %s %s (%s)", formattedId, formattedStatus, task.Description, relativeUpdatedTime))
		fmt.Println(taskStyle)
	}
	fmt.Println()

	return nil
}

func GetTaskByStatus(tasks []Task, status TaskStatus) []Task {
	filterTasks := []Task{}
	switch status {
	case "all":
		filterTasks = tasks
		break
	case TASK_STATUS_TODO:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_TODO {
				filterTasks = append(filterTasks, task)
			}
		}
		break
	case TASK_STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_IN_PROGRESS {
				filterTasks = append(filterTasks, task)
			}
		}
		break
	case TASK_STATUS_DONE:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_IN_PROGRESS {
				filterTasks = append(filterTasks, task)
			}
		}
		break
	}

	return filterTasks
}
