package main

import (
	"fmt"

	"github.com/HienVanNguyen0408/go_project/cmd"
	"github.com/charmbracelet/lipgloss"
)

func main() {

	rootCmd := cmd.InitCmd()

	if err := rootCmd.Execute(); err != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Padding(1, 2).
			Bold(true).
			Render(fmt.Sprintf("Error: %s", err))
		fmt.Println(errorStyle)
	}
}
