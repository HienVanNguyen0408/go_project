package cmd

import (
	"github.com/spf13/cobra"
)

func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task Tracker is a CLI tool for managing tasks",
		Long:  `Task Tracker tool is a CLI tool for managing tasks.`,
	}

	cmd.AddCommand(NewAddCmd())
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewUpdateDesciptionCmd())
	cmd.AddCommand(NewStatusDoneCmd())
	cmd.AddCommand(NewStatusInProgressCmd())
	cmd.AddCommand(NewStatusTodoCmd())
	return cmd
}
