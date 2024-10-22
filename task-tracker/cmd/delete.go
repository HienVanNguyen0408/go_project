package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/HienVanNguyen0408/go_project/internal/task"
	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Long:  `Delete a task from the task list.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(args)
		},
	}
	return cmd
}

func RunDeleteTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("task description is required")
	}

	taskId := args[0]
	//Check ID number
	taskIDInt, err := strconv.ParseInt(taskId, 10, 32)
	if err != nil {
		fmt.Println("Task ID must be a number")
		return err
	}

	return task.DeleteTask(taskIDInt)
}
