package cmd

import (
	"github.com/HienVanNguyen0408/go_project/internal/task"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListTask(args)
		},
	}
}

func RunListTask(args []string) error {
	if len(args) > 0 {
		status := task.TaskStatus(args[0])
		return task.GetListTask(status)
	}

	return task.GetListTask("all")
}
