package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/HienVanNguyen0408/go_project/internal/task"
	"github.com/spf13/cobra"
)

func NewUpdateDesciptionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a task",
		Long:  `Update a task in the task list.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateDesciption(args)
		},
	}
	return cmd
}

func RunUpdateDesciption(args []string) error {
	if len(args) == 0 {
		return errors.New("task description is required")
	}

	taskId := args[0]
	taskIdInt, err := strconv.ParseInt(taskId, 10, 32)
	if err != nil {
		fmt.Println("Task ID must be a number")
		return err
	}

	description := args[1]
	return task.UpdateTaskDescription(taskIdInt, description)
}

func NewStatusDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Mark a task as done",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, task.TASK_STATUS_DONE)
		},
	}
	return cmd
}

func NewStatusInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark a task as in-progress",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, task.TASK_STATUS_IN_PROGRESS)
		},
	}
	return cmd
}

func NewStatusTodoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-todo",
		Short: "Mark a task as todo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, task.TASK_STATUS_TODO)
		},
	}
	return cmd
}

func RunUpdateStatusCmd(args []string, status task.TaskStatus) error {
	if len(args) == 0 {
		return fmt.Errorf("task ID is required")
	}

	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	return task.UpdateTask(id, status)
}
