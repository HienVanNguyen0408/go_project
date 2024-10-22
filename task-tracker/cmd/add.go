package cmd

import (
	"errors"

	"github.com/HienVanNguyen0408/go_project/internal/task"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new task",
		Long:  `Add a new task to the task list.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTaskCmd(args)
		},
	}
	return cmd
}

func RunAddTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("task description is required")
	}

	description := args[0]
	return task.AddTask(description)
}
